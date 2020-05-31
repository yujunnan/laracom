module laracom/user-service/repo

go 1.13

replace github.com/yujunnan/laracom/user-service/proto/user => ../proto/user

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/yujunnan/laracom/user-service/proto/user v0.0.0-00010101000000-000000000000
)
