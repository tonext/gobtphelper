package gobtphelper

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GlobalGrpcPort string

func SendBeat(port string) {
	//credentials.NewClientTLSFromFile: 从输入的证眉眼文件中为客户端构造TLS凭证
	//grpc.WithTransportCredentials: 配置连接级别的安全凭证（例如 tls/ssl 返回一个dialoption
	grpcClient, err := grpc.DialContext(context.Background(), GetConfig("center_address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接 frame-center 失败！", err)
		return
	}
	client := NewCenterServiceClient(grpcClient)
	ips, err := GetLocalIPs()
	if err != nil {
		log.Println("获取本地IP地址失败!", err)
	}
	add := ips[0]

	timeout := GetConfig("timeout")
	if timeout == "" {
		timeout = "1000"
	}
	times, _ := strconv.Atoi(timeout)

	serviceFullName := GetServiceFullName()

	address := add + ":" + port

	for {
		res, err := client.SendBeat(context.Background(), &BeatReq{
			ServiceName: serviceFullName,
			Address:     address,
		})

		GlobalServices = res.Services

		if err != nil {
			log.Fatalln("网络连接错误!")
			break
		}
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

func GetServiceFullName() string {
	appName := GetConfig("app_name")
	code := generateRandomString(8)
	serviceFullName := appName + "@" + code
	zoneCode := GetConfig("zone")
	zoneCodeEnv := GetArgValue("-zone=")
	if zoneCodeEnv != "" {
		zoneCode = zoneCodeEnv
	}
	if zoneCode != "" {
		serviceFullName = appName + "-" + zoneCode + "@" + code
	}
	return serviceFullName
}

func GetGrpcPort() string {
	portString := GetConfig("port")
	protEnv := GetArgValue("-port=")
	if protEnv != "" {
		portString = protEnv
	}
	// port, err := strconv.Atoi(portString)
	// if err != nil {
	// 	log.Print("端口号获取失败！")
	// }
	// var result string
	// for {
	// 	result = checkPort(port)
	// 	if result != "" {
	// 		break
	// 	}
	// 	port++
	// }
	return portString
}

// func checkPort(port int) string {
// 	portString := strconv.Itoa(port)
// 	// 定义要检查的地址和端口
// 	address := "127.0.0.1:" + portString
// 	// 尝试创建一个TCP监听器
// 	listener, err := net.Listen("tcp", address)
// 	if err != nil {
// 		// 如果出错，打印错误信息
// 		fmt.Printf("无法监听 %s: %v\n", address, err)
// 		// 检查错误类型，判断是否为端口被占用
// 		return ""
// 		// if opErr, ok := err.(*net.OpError); ok {
// 		// 	if se, ok := opErr.Err.(*net.AddrError); ok && se.Err == "address already in use" {
// 		// 		fmt.Println("端口 " + portString + " 已被占用。")
// 		// 		return ""
// 		// 	}
// 		// }
// 	}
// 	// 如果监听成功，记得关闭监听器以释放资源
// 	listener.Close()
// 	fmt.Println("获取到端口 " + portString)
// 	return portString
// }

func GetWsPort() string {
	portString := GetConfig("ws_port")
	protEnv := GetArgValue("-ws_port=")
	if protEnv != "" {
		portString = protEnv
	}
	// port, err := strconv.Atoi(portString)
	// if err != nil {
	// 	log.Print("WebSocket端口号获取失败！")
	// }
	// var result string
	// for {
	// 	result = checkPort(port)
	// 	if result != "" {
	// 		break
	// 	}
	// 	port++
	// }
	return portString
}
