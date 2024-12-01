package gobtphelper

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// var GlobalLogicServiceList []LogicServiceChan

// var GlobalGatewayServiceList []GatewayServiceChan

// type LogicServiceChan struct {
// 	ServiceFullName string //格式：logic-hall-z0001
// 	LogicChan       chan *gobtphelper.ProtoMessage
// }

// type GatewayServiceChan struct {
// 	ServiceFullName string //格式：logic-hall-z0001
// 	GatewayChan     chan *gobtphelper.ProtoMessageResult
// }

var logicClientManager = NewLogicClientManager()
var gatewayClientManager = NewGatewayClientManager()

func StartGrpcClients() {
	for {
		//log.Printf("%#v", gobtphelper.GlobalServices)
		for _, item := range GlobalServices {
			if strings.HasPrefix(item.ServiceName, "frame-gateway") {
				_, exists := gatewayClientManager.GetClient(item.ServiceName)
				if !exists {
					gatewayClientManager.AddClient(GatewayClient{
						Conn:            startGatewayGrpcClient(item.ServiceName, item.Address),
						ServiceFullName: item.ServiceName,
					})
				}
			} else {
				_, exists := logicClientManager.GetClient(item.ServiceName)
				if !exists {
					logicClientManager.AddClient(LogicClient{
						Conn:            startLogicGrpcClient(item.ServiceName, item.Address),
						ServiceFullName: item.ServiceName,
					})
				}
			}
		}
		time.Sleep(time.Second * 2)
	}
}

func startLogicGrpcClient(serviceFullName string, serviceAddress string) *LogicServiceClient {
	log.Println("startGrpcClient serviceFullName = ", serviceFullName)
	//defer removeGrpcClient(serviceFullName)
	//credentials.NewClientTLSFromFile: 从输入的证眉眼文件中为客户端构造TLS凭证
	//grpc.WithTransportCredentials: 配置连接级别的安全凭证（例如 tls/ssl 返回一个dialoption
	grpcClient, err := grpc.DialContext(context.Background(), serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("连接 %v 失败！%v\n", serviceFullName, err)
		return nil
	}
	client := NewLogicServiceClient(grpcClient)
	return &client
}

func startGatewayGrpcClient(serviceFullName string, serviceAddress string) *GatewayServiceClient {
	log.Println("startGrpcClient serviceFullName = ", serviceFullName)
	grpcClient, err := grpc.DialContext(context.Background(), serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("连接 %v 失败！%v\n", serviceFullName, err)
		return nil
	}
	client := NewGatewayServiceClient(grpcClient)
	return &client
}

func SendToGateway(nodeCode string, gwCode string, fromServiceName string, accountId int64, actionName string, data []byte) *ProtoInt {
	zoneCode := GetZoneCode()
	message := &ProtoMessageResult{
		MsgId:       strconv.Itoa(int(time.Now().UnixMilli())),
		AccountId:   accountId,
		NodeCode:    &nodeCode,
		ZoneCode:    &zoneCode,
		GwCode:      &gwCode,
		ServiceName: &fromServiceName,
		ActionName:  &actionName,
		Data:        data,
		IsAck:       0,
	}
	serviceFullName := "frame-gateway-" + zoneCode + "@" + gwCode
	client, exists := gatewayClientManager.GetClient(serviceFullName)
	if exists {
		//log.Printf("找到client")
		resultMsg, err := (*client).SendToGateway(context.Background(), message)
		if err != nil {
			log.Printf("serviceFullName=%v, error=%v\n", serviceFullName, err)
		}
		return resultMsg
	}
	log.Printf("没有找到client, serviceFullName=%v", serviceFullName)
	return &ProtoInt{
		IsOk: 0,
	}
}

func SendToLogic(nodeCode string, serviceName string, req *ProtoMessage) *ProtoMessageResult {
	serviceFullName := GetServiceFullName(serviceName, nodeCode)
	client, exists := logicClientManager.GetClient(serviceName)
	//client := NewLogicServiceClient(grpcClient)
	if exists {
		resultMsg, err := (*client).SendToLogic(context.Background(), req)
		if err != nil {
			log.Printf("请求失败！serviceName=%v, error=%v\n", serviceFullName, err)
		}
		return resultMsg
	}
	log.Printf("没有找到client, serviceFullName=%v", serviceFullName)
	return nil
}