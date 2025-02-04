package rpc

import (
	"Go-Mall/app/client/conf"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/common/clientsuite"
	"Go-Mall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	UserClient userservice.Client

	once         sync.Once
	err          error
	registryAddr string
	commonSuite  client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: clientutils.ServiceName,
		})
		initUserClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	clientutils.MustHandleError(err)
}
