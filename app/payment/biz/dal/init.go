package dal

import (
	"Go-Mall/app/payment/biz/dal/mysql"
	"Go-Mall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
