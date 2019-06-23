### Server.go
# Start():启动websocket服务
# Send():向连接websocket的管道chan写入数据
# Read():读取在websocket管道中的数据
# Write():通过websocket协议向连接到ws的客户端发送数据

### wshandle.go
# WsHandle() 为中间件,  升级协议, 判断用户信息
