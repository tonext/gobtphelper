package gobtphelper

import (
	"context"
	"fmt"
	"log"
	"sync"
)

var gatewayClientCount = 0
var logicClientCount = 0

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
	logicClientCount++
	log.Printf("add serviceName = %v, logicClient总连接数: %v\n", client.ServiceFullName, logicClientCount)
}

func (sm *GatewayClientManager) AddClient(client GatewayClient) {
	sm.mutex.Lock()
	sm.clients[client.ServiceFullName] = client.Conn
	sm.mutex.Unlock()
	gatewayClientCount++
	log.Printf("add serviceName = %v, gatewayClient总连接数: %v\n", client.ServiceFullName, gatewayClientCount)
}

// 删除会话
func (sm *LogicClientManager) RemoveClient(serviceFullName string) {
	sm.mutex.Lock()
	delete(sm.clients, serviceFullName)
	sm.mutex.Unlock()
	logicClientCount--
	log.Printf("delete serviceName = %v, logicClient总连接数: %v\n", serviceFullName, logicClientCount)
}

func (sm *GatewayClientManager) RemoveClient(serviceFullName string) {
	sm.mutex.Lock()
	delete(sm.clients, serviceFullName)
	sm.mutex.Unlock()
	gatewayClientCount--
	log.Printf("delete serviceName = %v, gatewayClient总连接数: %v\n", serviceFullName, gatewayClientCount)
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
