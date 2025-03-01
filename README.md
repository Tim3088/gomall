# gomall
## 配置环境（脚手架）
### 安装 go 1.23.2
https://go.dev/
### 安装 make 并配置环境变量
https://gnuwin32.sourceforge.net/packages/make.htm
### 安装 docker
https://www.docker.com/
### 安装 cwgo
```bash
go install github.com/cloudwego/cwgo@latest
```
### 安装 protobuf v29.3 并配置环境变量
https://github.com/protocolbuffers/protobuf/releases
### 安装 kitex
```bash
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```
### 安装 Kitex 提供的服务注册与发现 consul 拓展
```bash
go get github.com/kitex-contrib/registry-consul
```
## 创建代码结构
在项目根目录下
```bash
mkdir rpc_gen
```
在app目录下
```bash
mkdir -p auth user product cart order checkout payment
```
## 根据接口文档（位于idl目录下）生成代码
### 生成 RPC客户端 代码
#### 认证客户端
在./rpc_gen目录下执行
```bash
cwgo client --type RPC --service auth  --module Go-Mall/rpc_gen --I ../idl --idl ..\idl\auth.proto
```
#### 用户客户端
在./rpc_gen目录下执行
```bash
cwgo client --type RPC --service user  --module Go-Mall/rpc_gen --I ../idl --idl ..\idl\user.proto
```
#### 商品客户端
在./rpc_gen目录下执行
```bash
cwgo client --type RPC --service product  --module Go-Mall/rpc_gen --I ../idl --idl ..\idl\product.proto
```
#### 购物车客户端
在./rpc_gen目录下执行
```bash
cwgo client --type RPC --service cart  --module Go-Mall/rpc_gen --I ../idl --idl ..\idl\cart.proto
```
#### 订单客户端
在./rpc_gen目录下执行
```bash
cwgo client --type RPC --service order  --module Go-Mall/rpc_gen --I ../idl --idl ..\idl\order.proto
```
#### 结算客户端
在./rpc_gen目录下执行
```bash
cwgo client --type RPC --service checkout  --module Go-Mall/rpc_gen --I ../idl --idl ..\idl\checkout.proto
```
#### 支付客户端
在./rpc_gen目录下执行
```bash
cwgo client --type RPC --service payment  --module Go-Mall/rpc_gen --I ../idl --idl ..\idl\payment.proto
```
### 生成 RPC服务端 代码
#### 认证服务
在./app/auth目录下执行
```bash
cwgo server --type RPC --service auth --module Go-Mall/app/auth --pass "-use Go-Mall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/auth.proto
```
#### 用户服务
在./app/user目录下执行
```bash
cwgo server --type RPC --service user --module Go-Mall/app/user --pass "-use Go-Mall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto
```
#### 商品服务
在./app/product目录下执行
```bash
cwgo server --type RPC --service product --module Go-Mall/app/product --pass "-use Go-Mall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto
```
#### 购物车服务
在./app/cart目录下执行
```bash
cwgo server --type RPC --service cart --module Go-Mall/app/cart --pass "-use Go-Mall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto
```
#### 订单服务
在./app/order目录下执行
```bash
cwgo server --type RPC --service order --module Go-Mall/app/order --pass "-use Go-Mall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto
```
#### 结算服务
在./app/checkout目录下执行
```bash
cwgo server --type RPC --service checkout --module Go-Mall/app/checkout --pass "-use Go-Mall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.proto
```
#### 支付服务
在./app/payment目录下执行
```bash
cwgo server --type RPC --service payment --module Go-Mall/app/payment --pass "-use Go-Mall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.proto
```
### 生成 HTTP服务端 代码
#### user_page
在./app/client目录下执行
```bash
cwgo server --type HTTP --idl ..\..\idl\client\user_page.proto --service client -module Go-Mall/app/client -I ../../idl
```
#### product_page
在./app/client目录下执行
```bash
cwgo server --type HTTP --idl ..\..\idl\client\product_page.proto --service client -module Go-Mall/app/client -I ../../idl
```
#### cart_page
在./app/client目录下执行
```bash
cwgo server --type HTTP --idl ..\..\idl\client\cart_page.proto --service client -module Go-Mall/app/client -I ../../idl
```
#### order_page
在./app/client目录下执行
```bash
cwgo server --type HTTP --idl ..\..\idl\client\order_page.proto --service client -module Go-Mall/app/client -I ../../idl
```
#### checkout_page
在./app/client目录下执行
```bash
cwgo server --type HTTP --idl ..\..\idl\client\checkout_page.proto --service client -module Go-Mall/app/client -I ../../idl
```
## 启动方式
### 从 dockerhub 中拉取镜像
1. golang:1.23
2. busybox:latest
### 构建docker镜像
${v}为版本号
#### 构建 HTTP服务端 镜像
```bash
make build-client v=latest
```
#### 构建 RPC服务端 镜像
```bash
make build-svc svc=auth v=latest
```
```bash
make build-svc svc=cart v=latest
```
```bash
make build-svc svc=checkout v=latest
```
```bash
make build-svc svc=order v=latest
```
```bash
make build-svc svc=payment v=latest
```
```bash
make build-svc svc=product v=latest
```
```bash
make build-svc svc=user v=latest
```
#### 启动项目根目录下的 docker-compose.yml
```bash
docker compose up
```

## 信用卡信息的验证
```bash
cd app/payment
go work use .
go mod tidy
go get github.com/durango/go-credit-card
```

项目目录结构
```
gomall/
├── app/
│   ├── auth/
│   │   ├── biz/
│   │   │   ├── service/
│   │   ├── conf/
│   │   │   ├── dev/
│   │   │   ├── online/
│   │   │   ├── test/
│   │   ├── log/
│   │   ├── script/
│   │   ├── .env
│   │   ├── build.sh
│   │   ├── go.mod
│   │   ├── handler.go
│   │   ├── main.go
│   ├── cart/
│   │   ├── biz/
│   │   │   ├── dal/
│   │   │   │   ├── mysql/
│   │   │   │   ├── redis/
│   │   │   ├── model/
│   │   │   ├── service/
│   │   ├── conf/
│   │   │   ├── dev/
│   │   │   ├── online/
│   │   │   ├── test/
│   │   ├── log/
│   │   ├── rpc/
│   │   ├── script/
│   │   ├── .env
│   │   ├── build.sh
│   │   ├── go.mod
│   │   ├── handler.go
│   │   ├── main.go
│   ├── checkout/
│   │   ├── biz/
│   │   │   ├── dal/
│   │   │   │   ├── mysql/
│   │   │   │   ├── redis/
│   │   │   ├── service/
│   │   ├── conf/
│   │   │   ├── dev/
│   │   │   ├── online/
│   │   │   ├── test/
│   │   ├── log/
│   │   ├── rpc/
│   │   ├── script/
│   │   ├── .env
│   │   ├── build.sh
│   │   ├── go.mod
│   │   ├── handler.go
│   │   ├── main.go
│   ├── client/
│   │   ├── biz/
│   │   │   ├── handler/
│   │   │   │   ├── cart/
│   │   │   │   ├── checkout/
│   │   │   │   ├── order/
│   │   │   │   ├── product/
│   │   │   │   ├── user/
│   │   │   ├── router/
│   │   │   │   ├── cart/
│   │   │   │   ├── checkout/
│   │   │   │   ├── order/
│   │   │   │   ├── product/
│   │   │   │   ├── user/
│   │   │   ├── service/
│   │   │   ├── utils/
│   │   ├── conf/
│   │   │   ├── dev/
│   │   │   ├── online/
│   │   │   ├── test/
│   │   ├── hertz_gen/
│   │   │   ├── api/
│   │   │   ├── client/
│   │   │   │   ├── cart/
│   │   │   │   ├── checkout/
│   │   │   │   ├── common/
│   │   │   │   ├── order/
│   │   │   │   ├── product/
│   │   │   │   ├── user/
│   │   ├── infra/
│   │   │   ├── rpc/
│   │   ├── log/
│   │   ├── middleware/
│   │   ├── script/
│   │   ├── utils/
│   │   ├── .env
│   │   ├── .hz
│   │   ├── build.sh
│   │   ├── go.mod
│   │   ├── main.go
│   ├── order/
│   │   ├── biz/
│   │   │   ├── dal/
│   │   │   │   ├── mysql/
│   │   │   │   ├── redis/
│   │   │   ├── model/
│   │   │   ├── service/
│   │   ├── conf/
│   │   │   ├── dev/
│   │   │   ├── online/
│   │   │   ├── test/
│   │   ├── log/
│   │   ├── rpc/
│   │   ├── script/
│   │   ├── .env
│   │   ├── build.sh
│   │   ├── go.mod
│   │   ├── handler.go
│   │   ├── main.go
│   ├── payment/
│   │   ├── biz/
│   │   │   ├── dal/
│   │   │   │   ├── mysql/
│   │   │   │   ├── redis/
│   │   │   ├── model/
│   │   │   ├── service/
│   │   ├── conf/
│   │   │   ├── dev/
│   │   │   ├── online/
│   │   │   ├── test/
│   │   ├── log/
│   │   ├── script/
│   │   ├── .env
│   │   ├── build.sh
│   │   ├── go.mod
│   │   ├── handler.go
│   │   ├── main.go
│   ├── product/
│   │   ├── biz/
│   │   │   ├── dal/
│   │   │   │   ├── mysql/
│   │   │   │   ├── redis/
│   │   │   ├── model/
│   │   │   ├── service/
│   │   ├── conf/
│   │   │   ├── dev/
│   │   │   ├── online/
│   │   │   ├── test/
│   │   ├── log/
│   │   ├── script/
│   │   ├── .env
│   │   ├── build.sh
│   │   ├── go.mod
│   │   ├── handler.go
│   │   ├── main.go
│   ├── user/
│   │   ├── biz/
│   │   │   ├── dal/
│   │   │   │   ├── mysql/
│   │   │   │   ├── redis/
│   │   │   ├── model/
│   │   │   ├── service/
│   │   ├── conf/
│   │   │   ├── dev/
│   │   │   ├── online/
│   │   │   ├── test/
│   │   ├── log/
│   │   ├── script/
│   │   ├── .env
│   │   ├── build.sh
│   │   ├── go.mod
│   │   ├── handler.go
│   │   ├── main.go
├── common/
│   ├── clientsuite/
│   ├── mtl/
│   ├── serversuite/
│   ├── utils/
│   ├── go.mod
├── conf/
├── db/
│   ├── sql/
│   │   ├── ini/
├── deploy/
│   ├── config/
├── idl/
│   ├── client/
├── rpc_gen/
│   ├── kitex_gen/
│   ├── rpc/
│   ├── go.mod
├── docker-compose.yaml
├── go.work
├── Makefile
├── README.md
```