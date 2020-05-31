module laracom/user-cli

go 1.13

require (
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/yujunnan/laracom/user-service/proto/user v0.0.0-20200530131054-29205d389e0d
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b
)

replace github.com/yujunnan/laracom/user-service/proto/user => ../user-service/proto/user
