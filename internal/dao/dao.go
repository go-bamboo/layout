package dao

import (
	"context"
	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/dao/query"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/store/gormx"
	"github.com/google/wire"
	"github.com/sony/sonyflake"
)

var ProviderSet = wire.NewSet(NewQuery, New)

// Dao dao interface
type Dao interface {
	Ping(ctx context.Context) (err error)

	Id() int64
}

// dao dao.
type dao struct {
	c  *conf.Data
	q  *query.Query
	sf *sonyflake.Sonyflake
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
		sf: sonyflake.NewSonyflake(sonyflake.Settings{MachineID: func() (uint16, error) {
			return 2, nil
		}}),
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

func (d *dao) Id() int64 {
	id, err := d.sf.NextID()
	if err != nil {
		return 0
	}
	return int64(id)
}
