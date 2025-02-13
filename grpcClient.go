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
		//清理不存在的服务
		getCleanList()

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
		time.Sleep(time.Second)
	}
}

func getCleanList() ([]string, []string) {
	var logicList []string
	var gatewayList []string
	for key := range logicClientManager.clients {
		if !isExistInCenter(key) {
			logicList = append(logicList, key)
		}
	}
	for key := range gatewayClientManager.clients {
		if !isExistInCenter(key) {
			gatewayList = append(gatewayList, key)
		}
	}
	// log.Printf("logicList = %#v \n", logicList)
	// log.Printf("gatewayList = %#v \n", gatewayList)
	for _, item := range logicList {
		logicClientManager.RemoveClient(item)
	}
	for _, item := range gatewayList {
		gatewayClientManager.RemoveClient(item)
	}
	return logicList, gatewayList
}

func isExistInCenter(serviceName string) bool {
	for _, item := range GlobalServices {
		if item.ServiceName == serviceName {
			return true
		}
	}
	return false
}

func startLogicGrpcClient(serviceFullName string, serviceAddress string) *LogicServiceClient {
	//log.Println("startGrpcClient serviceFullName = ", serviceFullName)
	//defer removeGrpcClient(serviceFullName)
	//credentials.NewClientTLSFromFile: 从输入的证眉眼文件中为客户端构造TLS凭证
	//grpc.WithTransportCredentials: 配置连接级别的安全凭证（例如 tls/ssl 返回一个dialoption
	debug := GetSectionConfig("helper", "debug")
	grpcClient, err := grpc.DialContext(context.Background(), serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		if debug == "true" {
			log.Printf("连接 %v 失败！%v\n", serviceFullName, err)
		}
		return nil
	}
	client := NewLogicServiceClient(grpcClient)
	return &client
}

func startGatewayGrpcClient(serviceFullName string, serviceAddress string) *GatewayServiceClient {
	//log.Println("startGrpcClient serviceFullName = ", serviceFullName)
	debug := GetSectionConfig("helper", "debug")
	grpcClient, err := grpc.DialContext(context.Background(), serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		if debug == "true" {
			log.Printf("连接 %v 失败！%v\n", serviceFullName, err)
		}
		return nil
	}
	client := NewGatewayServiceClient(grpcClient)
	return &client
}

func SendToGateway(fromServiceName string, zoneCode string, accountId int64, actionName string, data []byte) *ProtoInt {
	debug := GetSectionConfig("helper", "debug")
	message := &ProtoMessageResult{
		MsgId:       strconv.Itoa(int(time.Now().UnixMilli())),
		AccountId:   accountId,
		NodeCode:    &GlobalNodeCode,
		ZoneCode:    &zoneCode,
		ServiceName: &fromServiceName,
		ActionName:  &actionName,
		Data:        data,
		IsAck:       0,
	}
	serviceFullName := "frame-gateway-" + zoneCode
	//log.Printf("serviceFullName=%v", serviceFullName)
	count := 1
	for _, item := range GlobalServices {
		//log.Printf("正在处理第%v条注册信息.", count)
		if strings.HasPrefix(item.ServiceName, serviceFullName) {
			//log.Printf("找到对应的注册信息, item=%v", item)
			client, exists := gatewayClientManager.GetClient(item.ServiceName)
			if exists {
				//log.Printf("找到client, 发送消息到网关, gateway = %v", item.ServiceName)
				_, err := (*client).SendToGateway(context.Background(), message)
				if err != nil {
					if debug == "true" {
						log.Printf("error serviceFullName=%v, error=%v\n", item.ServiceName, err)
					}
				}
			} else {
				if debug == "true" {
					log.Printf("没有找到client, serviceFullName=%v", item.ServiceName)
				}
			}
		}
		count++
	}
	// client, exists := gatewayClientManager.GetClient(serviceFullName)
	// if exists {
	// 	//log.Printf("找到client")
	// 	resultMsg, err := (*client).SendToGateway(context.Background(), message)
	// 	if err != nil {
	// 		log.Printf("serviceFullName=%v, error=%v\n", serviceFullName, err)
	// 	}
	// 	return resultMsg
	// }
	// log.Printf("没有找到client, serviceFullName=%v", serviceFullName)
	return &ProtoInt{
		IsOk: 1,
	}
}

func SendToLogic(req *ProtoMessage) *ProtoMessageResult {
	debug := GetSectionConfig("helper", "debug")
	serviceFullName := GetServiceFullName(*req.ServiceName, *req.NodeCode, *req.ZoneCode)
	client, exists := logicClientManager.GetClient(serviceFullName)
	//client := NewLogicServiceClient(grpcClient)
	if exists {
		resultMsg, err := (*client).SendToLogic(context.Background(), req)
		if err != nil {
			if debug == "true" {
				log.Printf("请求失败！serviceName=%v, error=%v\n", serviceFullName, err)
			}
			return &ProtoMessageResult{
				MsgId:     req.MsgId,
				IsAck:     0,
				AccountId: 0,
			}
		}
		return resultMsg
	}
	if debug == "true" {
		log.Printf("没有找到client, serviceFullName=%v", serviceFullName)
	}
	return &ProtoMessageResult{
		MsgId:     req.MsgId,
		IsAck:     0,
		AccountId: 0,
	}
}
