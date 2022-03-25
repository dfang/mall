package svc

import (
	"github.com/dfang/mall/service/order/model"
	"github.com/dfang/mall/service/order/rpc/internal/config"
	"github.com/dfang/mall/service/product/rpc/product"
	"github.com/dfang/mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel

	UserRpc    user.User
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.Cache),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
