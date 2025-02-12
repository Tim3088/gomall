package main

import (
	"Go-Mall/app/payment/biz/dal"
	"Go-Mall/app/payment/conf"
	"Go-Mall/common/mtl"
	"Go-Mall/common/serversuite"
	"Go-Mall/common/utils"
	"Go-Mall/rpc_gen/kitex_gen/payment/paymentservice"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"strings"
)

var (
	serviceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	_ = godotenv.Load()

	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	// 初始化监控
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)

	// 初始化链路追踪
	p := mtl.InitTracing(serviceName)
	defer p.Shutdown(context.Background()) // 在服务关闭前上传剩余链路追踪数据

	dal.Init()
	opts := kitexInit()

	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}

	// 初始化服务
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	opts = append(opts,
		server.WithServiceAddr(addr),
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       RegistryAddr}))
	return
}
