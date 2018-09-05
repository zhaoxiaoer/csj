package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var SM *SessionManager = NewSessionManager()

// session操作接口
type Session interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Remove(key interface{}) error
	UpdateLastAccessed()
	IsTimeout() bool // 空闲时间是否超时
	GetId() string
}

// session实现
type SessionFromMemory struct {
	sid string
	// 防止同一用户快速地连续提交相同的请求而发生的竞态
	lock             sync.RWMutex
	lastAccessedTime time.Time
	maxAge           int64
	// 用户哈希表
	data map[interface{}]interface{}
}

func newSessionFromMemory(id string, ma int64) *SessionFromMemory {
	return &SessionFromMemory{
		sid:              id,
		lastAccessedTime: time.Now(),
		maxAge:           ma,
		data:             make(map[interface{}]interface{}),
	}
}

func (sfm *SessionFromMemory) Set(key, value interface{}) {
	sfm.lock.Lock()
	defer sfm.lock.Unlock()

	sfm.data[key] = value
}

func (sfm *SessionFromMemory) Get(key interface{}) interface{} {
	sfm.lock.RLock()
	defer sfm.lock.RUnlock()

	if value := sfm.data[key]; value != nil {
		return value
	}
	return nil
}

func (sfm *SessionFromMemory) Remove(key interface{}) error {
	sfm.lock.Lock()
	defer sfm.lock.Unlock()

	if value := sfm.data[key]; value != nil {
		delete(sfm.data, key)
	}
	return nil
}

func (sfm *SessionFromMemory) UpdateLastAccessed() {
	sfm.lock.Lock()
	defer sfm.lock.Unlock()

	sfm.lastAccessedTime = time.Now()
}

func (sfm *SessionFromMemory) IsTimeout() bool {
	return (sfm.lastAccessedTime.Unix()+sfm.maxAge < time.Now().Unix())
}

func (sfm *SessionFromMemory) GetId() string {
	return sfm.sid
}

// 全局存储器
type Storage interface {
	InitSession(sid string, maxAge int64) (Session, error)
	GetSession(sid string) (Session, error)
	//	SetSession(session Session) error
	DestroySession(sid string) error
	GCSession()
}

// 内存全局存储器
type FromMemory struct {
	// 所有请求都会访问该结构，此处的锁保证数据的一致性
	lock sync.Mutex
	// 全局哈希表
	sessions map[string]Session
}

func newFromMemory() *FromMemory {
	return &FromMemory{
		sessions: make(map[string]Session),
	}
}

func (fm *FromMemory) InitSession(sid string, maxAge int64) (Session, error) {
	fm.lock.Lock()
	defer fm.lock.Unlock()

	sfm := newSessionFromMemory(sid, maxAge)
	fm.sessions[sid] = sfm
	return sfm, nil
}

func (fm *FromMemory) GetSession(sid string) (Session, error) {
	fm.lock.Lock()
	defer fm.lock.Unlock()

	if session, ok := fm.sessions[sid]; ok {
		session.UpdateLastAccessed()
		return session, nil
	}
	return nil, fmt.Errorf("%v", "No session")
}

//func (fm *FromMemory) SetSession(session Session) error {
//	fm.lock.Lock()
//	defer fm.lock.Unlock()

//	fm.sessions[session.GetId()] = session
//	return nil
//}

func (fm *FromMemory) DestroySession(sid string) error {
	fm.lock.Lock()
	defer fm.lock.Unlock()

	if _, ok := fm.sessions[sid]; ok {
		delete(fm.sessions, sid)
	}
	return nil
}

func (fm *FromMemory) GCSession() {
	fm.lock.Lock()
	defer fm.lock.Unlock()

	if len(fm.sessions) < 1 {
		return
	}
	for k, v := range fm.sessions {
		if v.IsTimeout() {
			delete(fm.sessions, k)
		}
	}
}

type SessionManager struct {
	cookieName string
	storage    Storage
	maxAge     int64
	lock       sync.Mutex
}

func NewSessionManager() *SessionManager {
	sm := &SessionManager{
		cookieName: "JSESSIONID",
		storage:    newFromMemory(),
		maxAge:     60 * 30,
	}
	go sm.gc()

	return sm
}

func (sm *SessionManager) GetSession(w http.ResponseWriter,
	r *http.Request) Session {
	//	sm.lock.Lock()
	//	defer sm.lock.Unlock()
	cookie, err := r.Cookie(sm.cookieName)
	if err != nil || cookie.Value == "" {
		sid := sm.randomSID()
		session, _ := sm.storage.InitSession(sid, sm.maxAge)

		cookieS := &http.Cookie{
			Name:     sm.cookieName,
			Value:    url.QueryEscape(sid), // 转义特殊符号@#$%+-*等
			Path:     "/csj",
			HttpOnly: true,
			MaxAge:   int(sm.maxAge),
			Expires:  time.Now().Add(time.Duration(sm.maxAge)),
		}
		http.SetCookie(w, cookieS)
		return session
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, err := sm.storage.GetSession(sid)
		if err != nil {
			sid := sm.randomSID()
			sessionNew, _ := sm.storage.InitSession(sid, sm.maxAge)

			cookieS := &http.Cookie{
				Name:     sm.cookieName,
				Value:    url.QueryEscape(sid), // 转义特殊符号@#$%+-*等
				Path:     "/csj",
				HttpOnly: true,
				MaxAge:   int(sm.maxAge),
				Expires:  time.Now().Add(time.Duration(sm.maxAge)),
			}
			http.SetCookie(w, cookieS)
			return sessionNew
		}
		return session
	}
}

func (sm *SessionManager) gc() {
	//	sm.lock.Lock()
	//	defer sm.lock.Unlock()

	sm.storage.GCSession()
	time.AfterFunc(time.Duration(sm.maxAge*10), func() {
		sm.gc()
	})
}

func (sm *SessionManager) randomSID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
