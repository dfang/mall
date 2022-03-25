package svc

import (
	"github.com/dfang/mall/service/order/rpc/order"
	"github.com/dfang/mall/service/pay/model"
	"github.com/dfang/mall/service/pay/rpc/internal/config"
	"github.com/dfang/mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel
	UserRpc  user.User
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.Cache),
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
