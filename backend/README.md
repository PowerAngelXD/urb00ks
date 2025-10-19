# Backend for the urB00ks

# Project Structure
```
/
|-api      对应接口实现
|-dao      数据库相关操作
|-logger   自定义日志format
|-model    对bookinfo和userinfo的定义
|-service  介于dao和api两层之间的服务层，有利于解耦
|-test     测试代码
|-main.go  入口点
```

# Install
想要启动urB00ks的后端服务，您只需要：
```
go run main.go
```
即可