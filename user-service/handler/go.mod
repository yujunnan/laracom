module laracom/user-service/handler

go 1.13

require (
	github.com/yujunnan/laracom/user-service/proto/user v0.0.0-20200530131054-29205d389e0d
	github.com/yujunnan/laracom/user-service/repo v0.0.0-00010101000000-000000000000
	github.com/yujunnan/laracom/user-service/service v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b
)

replace github.com/yujunnan/laracom/user-service/proto/user => ../proto/user

replace github.com/yujunnan/laracom/user-service/service => ../service

replace github.com/yujunnan/laracom/user-service/repo => ../repo
