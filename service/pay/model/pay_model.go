package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	payFieldNames          = builder.RawFieldNames(&Pay{})
	payRows                = strings.Join(payFieldNames, ",")
	payRowsExpectAutoSet   = strings.Join(stringx.Remove(payFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	payRowsWithPlaceHolder = strings.Join(stringx.Remove(payFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cachePayIdPrefix  = "cache:pay:id:"
	cachePayOidPrefix = "cache:pay:oid:"
)

type (
	PayModel interface {
		Insert(ctx context.Context, data *Pay) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Pay, error)
		FindOneByOid(ctx context.Context, oid int64) (*Pay, error)
		Update(ctx context.Context, data *Pay) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPayModel struct {
		sqlc.CachedConn
		table string
	}

	Pay struct {
		Id         int64     `db:"id"`
		Uid        int64     `db:"uid"`    // 用户ID
		Oid        int64     `db:"oid"`    // 订单ID
		Amount     int64     `db:"amount"` // 产品金额
		Source     int64     `db:"source"` // 支付方式
		Status     int64     `db:"status"` // 支付状态
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewPayModel(conn sqlx.SqlConn, c cache.CacheConf) PayModel {
	return &defaultPayModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`pay`",
	}
}

func (m *defaultPayModel) Insert(ctx context.Context, data *Pay) (sql.Result, error) {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, payRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.Oid, data.Amount, data.Source, data.Status)
	}, payIdKey)
	return ret, err
}

func (m *defaultPayModel) FindOne(ctx context.Context, id int64) (*Pay, error) {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, id)
	var resp Pay
	err := m.QueryRowCtx(ctx, &resp, payIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", payRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPayModel) FindOneByOid(ctx context.Context, oid int64) (*Pay, error) {
	payOidKey := fmt.Sprintf("%s%v", cachePayOidPrefix, oid)
	var resp Pay
	err := m.QueryRow(&resp, payOidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `oid` = ? limit 1", payRows, m.table)
		return conn.QueryRow(v, query, oid)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPayModel) Update(ctx context.Context, data *Pay) error {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, payRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Oid, data.Amount, data.Source, data.Status, data.Id)
	}, payIdKey)
	return err
}

func (m *defaultPayModel) Delete(ctx context.Context, id int64) error {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, payIdKey)
	return err
}

func (m *defaultPayModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePayIdPrefix, primary)
}

func (m *defaultPayModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", payRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
