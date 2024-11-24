package gobtphelper

import (
	"os"

	"gopkg.in/ini.v1"
)

func GetConfig(key string) string {
	config, _ := ini.Load(getFilePath())
	return config.Section("").Key(key).String()
}

func GetSectionConfig(section string, key string) string {
	config, _ := ini.Load(getFilePath())
	return config.Section(section).Key(key).String()
}

func getFilePath() (filePath string) {
	args := os.Args
	if len(args) > 1 {
		filePath = "./conf/app-" + args[1] + ".ini"
	} else {
		filePath = "./conf/app.ini"
	}
	return
}

// func GetRemoteConfig(section string, key string) string {
// 	addresses := GetServiceAddress("frame-config")
// 	log.Printf("address=%v", addresses[0])
// 	grpcClient, err := grpc.DialContext(context.Background(), addresses[0], grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Printf("连接 frame-config 失败, 未能正确读取配置! key = %v", key)
// 		return ""
// 	}
// 	client := ConfigService.NewConfigServiceClient(grpcClient)
// 	args := os.Args
// 	var env string
// 	if len(args) > 1 {
// 		env = args[1]
// 	} else {
// 		env = "local"
// 	}
// 	res, err := client.LoadConfig(context.Background(), &ConfigService.ConfigReq{
// 		Env:     env,
// 		Section: section,
// 		Key:     key,
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		log.Println("连接配置中心出错!")
// 		return ""
// 	}
// 	return res.Value
// }
