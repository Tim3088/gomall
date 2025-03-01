package rpc

import (
	"Go-Mall/app/client/conf"
	clientutils "Go-Mall/app/client/utils"
	"Go-Mall/common/clientsuite"
	"Go-Mall/rpc_gen/kitex_gen/auth/authservice"
	"Go-Mall/rpc_gen/kitex_gen/cart/cartservice"
	"Go-Mall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"Go-Mall/rpc_gen/kitex_gen/order/orderservice"
	"Go-Mall/rpc_gen/kitex_gen/product/productcatalogservice"
	"Go-Mall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	AuthClient     authservice.Client
	OrderClient    orderservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client

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
		initAuthClient()
		initOrderClient()
		initCartClient()
		initCheckoutClient()
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

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth", commonSuite)
	clientutils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	clientutils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	clientutils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	clientutils.MustHandleError(err)
}
