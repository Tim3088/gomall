package dal

import (
	"Go-Mall/app/client/biz/dal/mysql"
	"Go-Mall/app/client/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
