package logic

import (
	"context"

	"github.com/dfang/mall/service/order/api/internal/svc"
	"github.com/dfang/mall/service/order/api/internal/types"
	"github.com/dfang/mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (*types.CreateResponse, error) {
	res, err := l.svcCtx.OrderRpc.Create(l.ctx, &order.CreateRequest{
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
