package dal

import (
	"Go-Mall/app/checkout/biz/dal/mysql"
	"Go-Mall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
