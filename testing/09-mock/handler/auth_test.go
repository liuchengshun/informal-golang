package v1

import (
	"encoding/gob"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
	"github.com/gophercloud/gophercloud/openstack/identity/v2/tokens"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"

	"razor/pkg/domain"
	mock "razor/pkg/domain/mocks"
)

func gomockInit(t *testing.T) *mock.MockAuthService {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mock.NewMockAuthService(ctrl)
}

func TestLoginUser(t *testing.T) {
	// mock methods
	var redisStore *redisstore.RedisStore
	mGet := gomonkey.ApplyMethod(reflect.TypeOf(redisStore), "Get", func(_ *redisstore.RedisStore, r *http.Request, name string) (*sessions.Session, error) {
		store := sessions.NewCookieStore([]byte(os.Getenv("RAZOR")))
		gob.Register(time.Time{})
		return store.Get(r, "razor")
	})
	defer mGet.Reset()

	var sess *sessions.Session
	mSave := gomonkey.ApplyMethod(reflect.TypeOf(sess), "Save", func(_ *sessions.Session, r *http.Request, w http.ResponseWriter) error {
		return nil
	})
	defer mSave.Reset()

	// mock interface
	m := gomockInit(t)
	auth := domain.Auth{
		User:               new(tokens.User),
		UserUnscopedToken:  new(tokens.Token),
		ProjectScopedToken: new(tokens.Token),
	}
	m.EXPECT().AuthenticateUserByPassword("Tongfang", "youyun").Return(nil)

	// start testing
	reader := strings.NewReader(`{"username":"Tongfang", "password":"youyun"}`)
	req, _ := http.NewRequest(http.MethodPost, "/login", reader)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	NewAuthHandler(m, router)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("login failed: want %d code but got %d code", http.StatusCreated, rr.Code)
	}
}

func TestLogoutUser(t *testing.T) {
	// mock methods
	var redisStore *redisstore.RedisStore
	mGet := gomonkey.ApplyMethod(reflect.TypeOf(redisStore), "Get", func(_ *redisstore.RedisStore, r *http.Request, name string) (*sessions.Session, error) {
		store := sessions.NewCookieStore([]byte(os.Getenv("RAZOR")))
		gob.Register(time.Time{})
		return store.Get(r, "razor")
	})
	defer mGet.Reset()

	m := gomockInit(t)
	req, _ := http.NewRequest(http.MethodGet, "/logout", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	NewAuthHandler(m, router)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("logout failed: want %d code but got %d code", http.StatusNoContent, rr.Code)
	}
}

func TestCurrentUser(t *testing.T) {
	// gomock init
	m := gomockInit(t)
	gob.Register(time.Time{})

	t.Run("good request testing", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/current", nil)
		rr := httptest.NewRecorder()
		var p *redisstore.RedisStore
		var method = gomonkey.ApplyMethod(reflect.TypeOf(p), "Get", func(_ *redisstore.RedisStore, r *http.Request, name string) (*sessions.Session, error) {
			store := sessions.NewCookieStore([]byte(os.Getenv("RAZOR")))
			sess, err := store.Get(r, "razor")

			expireTime := time.Date(2022, time.December, 0, 0, 0, 0, 0, time.Local)
			sess.Values["p_token_expires_at"] = expireTime
			sess.Values["p_id"] = "project_ID"
			sess.Values["p_name"] = "project_name"
			sess.Values["u_id"] = "user_ID"
			sess.Values["u_name"] = "user_name"
			sess.Values["is_admin"] = false
			return sess, err
		})
		defer method.Reset()

		// start testing
		router := mux.NewRouter()
		NewAuthHandler(m, router)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("currentHandler: want %d code but got %d code", http.StatusOK, rr.Code)
		}
	})

	t.Run("bad request testing", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/current", nil)
		rr := httptest.NewRecorder()
		var p *redisstore.RedisStore
		var m1 = gomonkey.ApplyMethod(reflect.TypeOf(p), "Get", func(_ *redisstore.RedisStore, r *http.Request, name string) (*sessions.Session, error) {
			store := sessions.NewCookieStore([]byte(os.Getenv("RAZOR")))
			sess, err := store.Get(r, "razor")

			expireTime := time.Date(2000, time.December, 0, 0, 0, 0, 0, time.Local)
			sess.Values["p_token_expires_at"] = expireTime
			return sess, err
		})
		defer m1.Reset()

		// start testing
		router := mux.NewRouter()
		NewAuthHandler(m, router)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Fatalf("bad request testing: want %d code but got %d code", http.StatusUnauthorized, rr.Code)
		}
	})
}
