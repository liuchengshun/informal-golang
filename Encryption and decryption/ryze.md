Rezy 项目使用了redis 的哪些功能
Set
func (c Client)  Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
1. 第一个输入参数为：context.Context类型，在使用的时候可以写入context的根上下文：context.Background(),也可以写入context.TODO()
2. 第二个参数是key-value的键值key,类型为string
3. 第三个参数为key-value的值value,类型为interface{}
4. 第四个参数为时间段，可以理解为数据的“保质期”，当过了这个设置的期限就会丢失。当设置为0时表示永不过期。
5. 返回值redis.StatusCmd类型，这个类型主要有很多方法来对Set()方法的结果处理，所以可以使用链式调用来返回需要的值，例如：Err()，返回Set()调用的错误值，Result()返回查询值和错误值等等。
6. 具体功能：将一对key-value键值对存储到redis中

Get
func (c Client)  Get(ctx context.Context, key string) *StringCmd
1. 第一个输入参数为：context.Context类型，在使用的时候可以写入context的根上下文：context.Background(),也可以写入context.TODO()
2. 第二个参数是key-value的键值key,类型为string
3. 返回值redis.StatusCmd类型，这个类型主要有很多方法来对Set()方法的结果处理，所以可以使用链式调用来返回需要的值，针对于Get()方法可以使用Resutl()方法来返回一个查询结果和错误信息。

Expire
func (c Client) Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
1. 第一个输入参数为：context.Context类型，在使用的时候可以写入context的根上下文：context.Background(),也可以写入context.TODO()
2. 第二个参数是key-value的键值key,类型为string
3. 第三个参数为过期时间
4. 返回值为redis.BoolCmd，其功能和redis.StatusCmd大致相同。也有很多方法来设置Expire的结果。
5. 具体功能：设置key的过期时间

使用具体场景：
场景1
1. 位置： ryze/pkg/handler/v1/auth_register.go  94行
2. 代码内容： rdb.Set(context.Background(), key, value, 24*time.Hour).Err()
3. 实现的功能：将使用uuid生成的string类型的值作为key，将结构体对象进行序列化后作为value，以24小时的过期时间存储到缓存中，如果出现错误则返回错误。
场景2
1. 位置： ryze/pkg/handler/v1/auth_register.go  143行
2. 代码内容：rdb.Get(context.Background(), code).Result()
3. 实现的功能：在redis中查询键值为code的value值并返回它。
场景3
1. 位置：ryze/pkg/handler/v1/auth_register.go  251行
2. 具体代码内容：rdb.Expire(context.Background(), code, time.Second)
3. 实现的功能：设置键值为code的键值对的过期时间为1秒
场景4
1. 位置：ryze/pkg/model/keystone/token.go   136行
2. 代码内容：rdb.Get(ctx, "ryze_token").Result()
3. 实现的功能：查询key为"ryze_token"的value值并返回这个
