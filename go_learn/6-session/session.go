package main

import (
	"fmt"
	"sync"
	"encoding/base64"
	"crypto/rand"
	"net/http"
	"net/url"
)

var provides = make(map[string]Provider)

// 全局的session管理器
type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxLifeTime int64
}

// 我们知道session是保存在服务器端的数据，它可以以任何的方式存储，
// 比如存储在内存、数据库或者文件中。因此我们抽象出一个Provider接口，
// 用以表征session管理器底层存储结构。
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

// 对Session的处理基本就 设置值、读取值、删除值以及获取当前sessionID这四个操作，
// 所以我们的Session接口也就实现这四个操作。
type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // back current sessionID
}


// 确保唯一的provider
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

// 设置唯一的sessionID
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// 创建session
// 我们需要为每个来访用户分配或获取与他相关连的Session，
// 以便后面根据Session信息来验证操作。SessionStart这个函数就是用来检测
// 是否已经有某个Session与当前来访用户发生了关联，如果没有则创建之。
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	// 加锁，
	manager.lock.Lock()
	defer manager.lock.Unlock()
	// 将request中的cookie参数表示出来
	cookie, err := r.Cookie(manager.cookieName)
	// 当这个cookieName不存在时，表示没有被赋值过cookieId,这里进行赋值
	if err != nil || cookie.Value == "" {
		// 获取sessionID
		sid := manager.sessionId()
		// 初始化session
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}


func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}
