package rpc

import (
	"Go-Mall/app/client/conf"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/common/clientsuite"
	"Go-Mall/rpc_gen/kitex_gen/product/productcatalogservice"
	"Go-Mall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client

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
		initProductClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	clientutils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	clientutils.MustHandleError(err)
}
