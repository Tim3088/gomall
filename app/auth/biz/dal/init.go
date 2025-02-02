package dal

import (
	"Go-Mall/app/auth/biz/dal/mysql"
	"Go-Mall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
