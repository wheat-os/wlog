# 实例

> 直接通过函数调用

```go
// 直接调用 wlog 库中的 Debug Info Warn ..., Debugf Infof Warnf ...
// 直接将日志打印在控制台中
wlog.Info("123", 123, map[string]int{"qqq": 111})
log := wlog.New()
log.Info("123", 123, map[string]int{"qqq": 111})
```

> 通过 New() 方法来获取 logger 实例, 对实例进行修改，改变日志输出

```go
// 获取实例
log := wlog.New()

// 选择打开文件
file, err := os.OpenFile("log.text", os.O_WRONLY|os.O_CREATE, 0775)
if err != nil {
return
}

// 修改日志输出
log.SetOptions(wlog.WithOutput(file))

// 日志保存
log.Info("123", 123, map[string]int{"qqq": 111})
log.Debug("123", 123, map[string]int{"qqq": 111})
```