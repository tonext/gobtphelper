package gobtphelper

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GlobalGrpcPort string
var GlobalNodeCode string = generateRandomString(8)
var GloablZoneCode string = GetZoneCode()
var GlobalServiceFullName string = GetRandomServiceFullName()

func SendBeat(port string) {
	//credentials.NewClientTLSFromFile: 从输入的证眉眼文件中为客户端构造TLS凭证
	//grpc.WithTransportCredentials: 配置连接级别的安全凭证（例如 tls/ssl 返回一个dialoption
	grpcClient, err := grpc.DialContext(context.Background(), GetConfig("center_address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接 frame-center 失败！", err)
		return
	}
	client := NewCenterServiceClient(grpcClient)

	ip := GetRegisterIp()

	timeout := GetConfig("timeout")
	if timeout == "" {
		timeout = "1000"
	}
	times, _ := strconv.Atoi(timeout)

	GlobalServiceFullName := GetRandomServiceFullName()

	address := ip + ":" + port

	for {
		res, err := client.SendBeat(context.Background(), &BeatReq{
			ServiceName: GlobalServiceFullName,
			Address:     address,
		})
		if err != nil {
			log.Fatalln("网络连接错误!")
			break
		}
		GlobalServices = res.Services
		//fmt.Printf("%#v", res)
		//log.Println("发送心跳信息.")
		time.Sleep(time.Duration(times) * time.Millisecond)
	}
}

// generateRandomString 生成一个 n 位的随机字符串
func generateRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机数生成器的种子

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetLocalIPs() ([]string, error) {
	var ips []string
	//local_ip := GetConfig("local_ip")
	//ips = strings.Split(local_ip, ",")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	if len(ips) == 0 {
		return nil, fmt.Errorf("no non-loopback IPv4 addresses found")
	}
	return ips, nil
}

func GetRegisterIp() string {
	return GetConfig("register_ip")
}

func GetRandomServiceFullName() string {
	appName := GetConfig("app_name")
	serviceFullName := appName + "@" + GlobalNodeCode
	if GloablZoneCode != "" {
		serviceFullName = appName + "-" + GloablZoneCode + "@" + GlobalNodeCode
	}
	return serviceFullName
}

func GetServiceFullName(serviceName string, nodeCode string, zoneCode string) string {
	for _, v := range GlobalServices {
		// log.Println("v=" + v.ServiceName)
		// log.Println("zoneCode =" + GloablZoneCode)
		tmp := strings.Split(v.ServiceName, "-")
		if len(tmp) == 2 {
			if strings.Contains(v.ServiceName, serviceName) && strings.Contains(v.ServiceName, nodeCode) {
				//tmp := strings.Split(v.ServiceName, "@")
				return v.ServiceName
			}
		} else if len(tmp) == 3 {
			if strings.Contains(v.ServiceName, serviceName) && strings.Contains(v.ServiceName, nodeCode) && strings.Contains(v.ServiceName, zoneCode) {
				//tmp := strings.Split(v.ServiceName, "@")
				return v.ServiceName
			}
		}
	}
	return ""
}

func GetGrpcPort() string {
	portString := GetConfig("port")
	protEnv := GetArgValue("-port=")
	if protEnv != "" {
		portString = protEnv
	}
	return portString
}

func GetWsPort() string {
	portString := GetConfig("ws_port")
	protEnv := GetArgValue("-ws_port=")
	if protEnv != "" {
		portString = protEnv
	}
	return portString
}
