package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

// CommonGrpcClientSuite 定义自定义Suite
type CommonGrpcClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

// Options 实现Suite接口的Options方法
func (s CommonGrpcClientSuite) Options() []client.Option {
	// 创建consul解析器
	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		panic(err)
	}

	// 基础客户端配置
	opts := []client.Option{
		// 指定一个 Resolver 进行服务发现
		client.WithResolver(r),
		// 基础客户端配置
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		// 使用 GRPC 作为传输协议
		client.WithTransportProtocol(transport.GRPC),
	}

	opts = append(opts,
		// 设置 Client 侧的 Service 信息，用于服务发现
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// 设置链路追踪
		client.WithSuite(tracing.NewClientSuite()),
	)

	return opts
}
