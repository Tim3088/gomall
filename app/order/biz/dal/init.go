package dal

import (
	"Go-Mall/app/order/biz/dal/mysql"
	"Go-Mall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
