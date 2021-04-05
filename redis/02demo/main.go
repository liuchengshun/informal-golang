package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/uuid"
	redisV8 "github.com/go-redis/redis/v8"

	"gitlab.ustack.com/ued/ryze/pkg/config"
	"gitlab.ustack.com/ued/ryze/pkg/model/cryptic"
)

type Client = redisV8.UniversalClient

type Options = redisV8.UniversalOptions

type redisHook struct {
	redisV8.Hook
}

var _ redisV8.Hook = redisHook{}

var rdb Client

var encodeData = make(map[string]string)


func (redisHook) BeforeProcess(ctx context.Context, cmd redisV8.Cmder) (context.Context, error) {
	if cmd.Name() == "get" && cmd.FullName() == "get" {
		args := cmd.Args()
		if args[0] == "get" && args[1] == "ryze_token" {
			err := rdb.GetSet(ctx, "ryze_token", cryptic.AesDecryptCFB(encodeData["ryze_token"])).Err()
			if err != nil {
				return ctx, err
			}
		}
	}
	return ctx, nil
}

func (redisHook) AfterProcess(ctx context.Context, cmd redisV8.Cmder) error {
	args := cmd.Args()
	if cmd.Name() == "set" && cmd.FullName() == "set" {
		if args[0] == "set" && args[1] == "ryze_token" {
			if value, ok := args[2].(string); ok {
				valueEn := cryptic.AesEncryptCFB(value)
				err := rdb.GetSet(ctx, "ryze_token", valueEn).Err()
				if err != nil {
					return err
				}
				encodeData["ryze_token"] = valueEn
			}
		}

		if uuidKey, ok := args[1].(string); ok {
			id, err := uuid.Parse(uuidKey)
			idStr := fmt.Sprintf("%v", id)
			if err == nil && args[0] == "set" && args[1] == idStr {
				user := &struct{
					Username    string `json:"username"`
					Password    string `json:"password"`
					CompanyName string `json:"company_name"`
					Email       string `json:"email"`
					Phone       string `json:"phone"`
					Code        string `json:"code"`
				}{}
				if v, ok := args[2].(string); ok {
					if err := json.Unmarshal([]byte(v), &user); err != nil {
						return err
					}
				}
				user.Password = cryptic.AesEncryptCFB(user.Password)
				json.Marshal(user)
			}
		}
	}

	if cmd.Name() == "get" && cmd.FullName() == "get" {
		if args[0] == "get" && args[1] == "ryze_token" {
			err := rdb.GetSet(ctx, "ryze_token", encodeData["ryze_token"]).Err()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 这是提交前的版本  4/5 22:00