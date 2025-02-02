package dal

import (
	"Go-Mall/app/product/biz/dal/mysql"
	"Go-Mall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
