# 本目录主要学习context 

## 01-context_explain

这个目录的主要参考内容链接是：https://blog.csdn.net/u011957758/article/details/82948750

**常用的使用姿势**

1. web编程中，一个请求对应多个goroutine之间的数据交互

2. 超时控制

3. 上下文控制

**context的底层结构**

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

## 02-context_juejin

这个目录主要参考了 `https://juejin.cn/post/6844904070667321357 `


