module github.com/yujunnan/laracom/user-cli

go 1.13

replace github.com/yujunnan/laracom/user-service => ../user-service

require (
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/yujunnan/laracom/user-service v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
)
