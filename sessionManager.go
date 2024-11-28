package gobtphelper

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var sessionCount = 0

// 定义一个会话结构体
type Session struct {
	Conn      *websocket.Conn
	AccountId string
}

// 会话管理器
type SessionManager struct {
	sessions map[string]*Session
	mutex    sync.RWMutex
}

// 新建会话管理器
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}

// 添加会话
func (sm *SessionManager) AddSession(session *Session) {
	sm.mutex.Lock()
	sm.sessions[session.AccountId] = session
	sm.mutex.Unlock()
	sessionCount++
	log.Printf("新的客户端连接: accountId = %v, 总连接数: %v\n", session.AccountId, sessionCount)
}

// 删除会话
func (sm *SessionManager) RemoveSession(accountId string) {
	sm.mutex.Lock()
	delete(sm.sessions, accountId)
	sm.mutex.Unlock()
	sessionCount--
	log.Printf("新的客户端连接: accountId = %v, 总连接数: %v\n", accountId, sessionCount)
}

// 获取会话
func (sm *SessionManager) GetSession(accountId string) (*Session, bool) {
	sm.mutex.RLock()
	session, exists := sm.sessions[accountId]
	sm.mutex.RUnlock()
	return session, exists
}

// 向指定用户发送消息
func (sm *SessionManager) SendMessage(messageResult *ProtoMessageResult) error {
	data, _ := proto.Marshal(messageResult)
	log.Printf("messageResult = %v", messageResult)
	session, exists := sm.GetSession(strconv.Itoa(int(messageResult.AccountId)))
	if !exists {
		return fmt.Errorf("session not found for user %v", messageResult.AccountId)
	}
	err := session.Conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Println("Failed to send message:", err)
		return err
	}
	return nil
}
