package logic

import (
	"context"

	"github.com/dfang/mall/service/order/api/internal/svc"
	"github.com/dfang/mall/service/order/api/internal/types"
	"github.com/dfang/mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (*types.UpdateResponse, error) {
	_, err := l.svcCtx.OrderRpc.Update(l.ctx, &order.UpdateRequest{
		Id:     req.Id,
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}
