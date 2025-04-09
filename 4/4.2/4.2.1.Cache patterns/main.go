package main

import (
	"sync"
	"time"
)

type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
}
type SessionManager struct {
	sessions map[string]Session
	mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]Session),
	}
}

func (sm *SessionManager) StartSession(userID string) string {
	expireAt := time.Now().Add(ExpiresAt).UnixNano()
	sm.locker.Lock()
	defer sm.locker.Unlock()
	sm.sessions[userID] = NewSessionManager(value, expireAt)
}

func (sm *SessionManager) GetSession(sessionID string) (Session, bool) {
	sm.locker.RLock()
	defer sm.locker.RUnlock()
	data, found := sm.sessions[sessionID]
	if !found || data.IsExpired() {
		return nil, false
	}

	return data.value, true
}

func main() {

}
