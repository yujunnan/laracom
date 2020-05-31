module laracom/user-cli

go 1.13

replace github.com/yujunnan/laracom/user-service/proto/user => ../user-service/proto/user

require (
	github.com/micro/cli v0.2.0
	github.com/yujunnan/laracom/user-service/proto/user v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b
)
