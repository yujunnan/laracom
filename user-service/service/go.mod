module laracom/user-service/service

go 1.13

replace github.com/yujunnan/laracom/user-service/proto/user => ../proto/user

replace github.com/yujunnan/laracom/user-service/repo => ../repo

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/yujunnan/laracom/user-service/proto/user v0.0.0-00010101000000-000000000000
	github.com/yujunnan/laracom/user-service/repo v0.0.0-00010101000000-000000000000
)
