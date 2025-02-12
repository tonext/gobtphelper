package gobtphelper

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// 定义一个会话结构体
type LogicClient struct {
	Conn            *LogicServiceClient
	ServiceFullName string
}

type GatewayClient struct {
	Conn            *GatewayServiceClient
	ServiceFullName string
}

// 会话管理器
type LogicClientManager struct {
	clients map[string]*LogicServiceClient
	mutex   sync.RWMutex
}

type GatewayClientManager struct {
	clients map[string]*GatewayServiceClient
	mutex   sync.RWMutex
}

// 新建会话管理器
func NewLogicClientManager() *LogicClientManager {
	return &LogicClientManager{
		clients: make(map[string]*LogicServiceClient),
	}
}
func NewGatewayClientManager() *GatewayClientManager {
	return &GatewayClientManager{
		clients: make(map[string]*GatewayServiceClient),
	}
}

// 添加会话
func (sm *LogicClientManager) AddClient(client LogicClient) {
	sm.mutex.Lock()
	sm.clients[client.ServiceFullName] = client.Conn
	sm.mutex.Unlock()
	//logicClientCount++
	debug := GetSectionConfig("helper", "debug")
	if debug == "true" {
		log.Printf("新增logicClient服务: %v, 连接数: %v\n", client.ServiceFullName, len(sm.clients))
	}
}

func (sm *GatewayClientManager) AddClient(client GatewayClient) {
	sm.mutex.Lock()
	sm.clients[client.ServiceFullName] = client.Conn
	sm.mutex.Unlock()
	//gatewayClientCount++
	debug := GetSectionConfig("helper", "debug")
	if debug == "true" {
		log.Printf("新增gatewayClient服务 %v, 连接数: %v\n", client.ServiceFullName, len(sm.clients))
	}
}

// 删除会话
func (sm *LogicClientManager) RemoveClient(serviceFullName string) {
	sm.mutex.Lock()
	delete(sm.clients, serviceFullName)
	sm.mutex.Unlock()
	//logicClientCount--
	debug := GetSectionConfig("helper", "debug")
	if debug == "true" {
		log.Printf("删除logicClient服务 %v, 连接数: %v\n", serviceFullName, len(sm.clients))
	}
}

func (sm *GatewayClientManager) RemoveClient(serviceFullName string) {
	sm.mutex.Lock()
	delete(sm.clients, serviceFullName)
	sm.mutex.Unlock()
	//gatewayClientCount--
	log.Printf("删除gatewayClient服务 %v, 连接数: %v\n", serviceFullName, len(sm.clients))
}

// 获取会话
func (sm *LogicClientManager) GetClient(serviceFullName string) (*LogicServiceClient, bool) {
	sm.mutex.RLock()
	client, exists := sm.clients[serviceFullName]
	sm.mutex.RUnlock()
	return client, exists
}

func (sm *GatewayClientManager) GetClient(serviceFullName string) (*GatewayServiceClient, bool) {
	sm.mutex.RLock()
	client, exists := sm.clients[serviceFullName]
	sm.mutex.RUnlock()
	return client, exists
}

// 向指定用户发送消息
func (sm *LogicClientManager) SendMessage(serviceFullName string, message *ProtoMessage) (*ProtoMessageResult, error) {
	client, exists := sm.GetClient(serviceFullName)
	if !exists {
		return &ProtoMessageResult{}, fmt.Errorf("logicClient not found for serviceFullName= %v", serviceFullName)
	}
	return (*client).SendToLogic(context.Background(), message)
}

func (sm *GatewayClientManager) SendMessage(serviceFullName string, message *ProtoMessageResult) (*ProtoInt, error) {
	client, exists := sm.GetClient(serviceFullName)
	if !exists {
		return &ProtoInt{
			IsOk: 0,
		}, fmt.Errorf("logicClient not found for serviceFullName= %v", serviceFullName)
	}
	return (*client).SendToGateway(context.Background(), message)
}
