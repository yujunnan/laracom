module laracom/user-service/repo

go 1.13

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/yujunnan/laracom/user-service/proto/user v0.0.0-20200530131054-29205d389e0d
)

replace github.com/yujunnan/laracom/user-service/proto/user => ../proto/user
