package gobtphelper

import (
	"os"
	"strings"

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

func getFilePath() string {
	envValue := GetArgValue("-env=")
	var filePath string
	if envValue != "" {
		filePath = "./conf/app-" + envValue + ".ini"
	} else {
		filePath = "./conf/app.ini"
	}
	return filePath
}

func GetArgValue(argName string) string {
	args := os.Args
	argValue := ""
	for _, arg := range args {
		if strings.HasPrefix(arg, argName) {
			// 提取 -env= 后面的值
			argValue = strings.TrimPrefix(arg, argName)
			//fmt.Println("环境变量值:", envValue)
		}
	}
	return argValue
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
