module laracom/demo-cli

go 1.13

require (
	github.com/micro/go-micro v1.18.0
	github.com/yujunnan/laracom/demo-service/proto/demo v0.0.0-20200530131054-29205d389e0d
)

replace github.com/yujunnan/laracom/demo-service/proto/demo => ../demo-service/proto/demo
