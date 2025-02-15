package rpc

import (
	"Go-Mall/app/cart/conf"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/common/clientsuite"
	"Go-Mall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	ProductClient productcatalogservice.Client

	once         sync.Once
	err          error
	registryAddr string
	commonSuite  client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: clientutils.ServiceName,
		})
		initProductClient()
	})
}
func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	clientutils.MustHandleError(err)
}
