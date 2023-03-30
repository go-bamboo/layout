package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/layout/internal/dao/model"
	"github.com/go-bamboo/layout/internal/dao/query"
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
	c           *conf.Data
	db          *gormx.DB
	q           *query.Query
	sf          *sonyflake.Sonyflake
	klineLock   sync.RWMutex
	klines      map[string]map[string]dataframe.DataFrame
	depthLock   sync.RWMutex
	depths      map[string]map[float64]float64
	SymbolPrice map[string]float64
}

func NewQuery(c *conf.Data) (db *gormx.DB) {
	db, err := gormx.New(c.Database)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// New new a dao and return.
func New(c *conf.Data, db *gormx.DB) (d Dao, cb func(), err error) {
	return newDao(c, db)
}

func newDao(c *conf.Data, db *gormx.DB) (d *dao, cb func(), err error) {
	d = &dao{
		c:  c,
		db: db,
		q:  query.Use(db),
		sf: sonyflake.NewSonyflake(sonyflake.Settings{MachineID: func() (uint16, error) {
			return 2, nil
		}}),
		klines:      map[string]map[string]dataframe.DataFrame{},
		depths:      map[string]map[float64]float64{},
		SymbolPrice: map[string]float64{},
	}
	cb = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	for symbol, m := range d.klines {
		for interval, df := range m {
			if df.Nrow() <= 0 {
				continue
			}
			startTimeCol, err := df.Subset(df.Nrow() - 1).Col("StartTime").Int()
			if err != nil {
				log.Error(err)
				continue
			}
			endTimeCol, err := df.Subset(0).Col("EndTime").Int()
			if err != nil {
				log.Error(err)
				continue
			}
			log.Debugf("event: %v, symbol: %v, interval: %v, st: %v, et: %v", "kline", symbol, interval, startTimeCol[0], endTimeCol[0])
			filename := fmt.Sprintf("%v_%v_%v_%v.csv", "kline", symbol, interval, startTimeCol[0])
			csvPath := path.Join(d.c.CsvPath, filename)
			fn, err := os.OpenFile(csvPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
			if err != nil {
				log.Error(err)
				continue
			}
			df.WriteCSV(fn, dataframe.WriteHeader(true))
			id, _ := d.sf.NextID()
			d.updateKlineCSV(context.TODO(), &model.CexBinanceKlineCsv{
				ID:        int64(id),
				Event:     "kline",
				Symbol:    symbol,
				Interval:  interval,
				StartTime: int64(startTimeCol[0]),
				EndTime:   int64(endTimeCol[0]),
				File:      filename,
			})
		}
	}
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

func (d *dao) Price(symbol string) float64 {
	return d.SymbolPrice[symbol]
}
