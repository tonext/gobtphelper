package gobtphelper

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendBeat() {
	//credentials.NewClientTLSFromFile: 从输入的证眉眼文件中为客户端构造TLS凭证
	//grpc.WithTransportCredentials: 配置连接级别的安全凭证（例如 tls/ssl 返回一个dialoption
	grpcClient, err := grpc.DialContext(context.Background(), GetConfig("center_address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("连接 frame-center 失败！", err)
		return
	}
	client := NewCenterServiceClient(grpcClient)
	code := generateRandomString(8)
	ips, err := GetLocalIPs()
	if err != nil {
		log.Println("获取本地IP地址失败!", err)
	}
	add := ips[0]
	port := GetConfig("port")
	appName := GetConfig("app_name")
	for {
		res, err := client.SendBeat(context.Background(), &BeatReq{
			ServiceName: appName + "@" + code,
			Address:     add + ":" + port,
		})

		GlobalServices = res.Services

		if err != nil {
			log.Fatalln("网络连接错误!")
			break
		}
		//fmt.Printf("%#v", res)
		log.Println("发送心跳信息.")
		time.Sleep(1000 * time.Millisecond)
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
