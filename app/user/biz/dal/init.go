package dal

import (
	"Go-Mall/app/user/biz/dal/mysql"
	"Go-Mall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
