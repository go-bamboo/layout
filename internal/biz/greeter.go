package biz

import (
	"github.com/go-bamboo/layout/internal/dao"
	"github.com/go-bamboo/layout/internal/dao/query"
)

type GreeterUsecase struct {
	q *query.Query
	d dao.Dao
}

func NewGreeterUsecase(q *query.Query, d dao.Dao) (uc *GreeterUsecase, err error) {
	uc = &GreeterUsecase{
		q: q,
		d: d,
	}

	// 启动通知
	return
}

func (uc *GreeterUsecase) Close() {
}
