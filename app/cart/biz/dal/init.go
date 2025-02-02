package dal

import (
	"Go-Mall/app/cart/biz/dal/mysql"
	"Go-Mall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
