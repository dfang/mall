package logic

import (
	"context"

	"github.com/dfang/mall/service/product/api/internal/svc"
	"github.com/dfang/mall/service/product/api/internal/types"
	"github.com/dfang/mall/service/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (*types.DetailResponse, error) {
	res, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     res.Id,
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  res.Stock,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
