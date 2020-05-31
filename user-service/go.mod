module laracom/user-service

go 1.13

require (
	github.com/micro/go-micro v1.18.0
	github.com/yujunnan/laracom/user-service/db v0.0.0-00010101000000-000000000000
	github.com/yujunnan/laracom/user-service/handler v0.0.0-20200530131054-29205d389e0d
	github.com/yujunnan/laracom/user-service/proto/user v0.0.0-20200530131054-29205d389e0d
	github.com/yujunnan/laracom/user-service/repo v0.0.0-20200530131054-29205d389e0d
)

replace github.com/yujunnan/laracom/user-service => ./

replace github.com/yujunnan/laracom/user-service/proto/user => ./proto/user

replace github.com/yujunnan/laracom/user-service/db => ./db

replace github.com/yujunnan/laracom/user-service/handler => ./handler

replace github.com/yujunnan/laracom/user-service/service => ./service
