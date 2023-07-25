package dao

import (
	"context"
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/dao/query"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/store/gormx"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewQuery, New)

// Dao dao interface
type Dao interface {
	Ping(ctx context.Context) (err error)
}

// dao dao.
type dao struct {
	c *conf.Data
	q *query.Query
}

func NewQuery(c *conf.Data) *query.Query {
	db, err := gormx.New(c.Database)
	if err != nil {
		log.Fatal(err)
	}
	return query.Use(db)
}

// New new a dao and return.
func New(c *conf.Data, q *query.Query) (d Dao, cb func(), err error) {
	return newDao(c, q)
}

func newDao(c *conf.Data, q *query.Query) (d *dao, cb func(), err error) {
	d = &dao{
		c: c,
		q: q,
	}
	cb = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
