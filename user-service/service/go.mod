module laracom/user-service/service

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/yujunnan/laracom/user-service/proto/user v0.0.0-20200530131054-29205d389e0d
	github.com/yujunnan/laracom/user-service/repo v0.0.0-00010101000000-000000000000
)

replace github.com/yujunnan/laracom/user-service/proto/user => ../proto/user

replace github.com/yujunnan/laracom/user-service/repo => ../repo
