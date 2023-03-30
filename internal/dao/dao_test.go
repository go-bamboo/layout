package dao

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/go-bamboo/layout/internal/conf"
	"github.com/go-bamboo/pkg/config"
)

var d *dao
var ctx = context.Background()

func TestMain(m *testing.M) {
	flag.Set("conf", "file:///../../configs/test.yaml")
	var err error
	var bc conf.Bootstrap
	config.Load(&bc)
	var cf func()
	if d, cf, err = newDao(bc.Data, nil); err != nil {
		panic(err)
	}
	ret := m.Run()
	cf()
	os.Exit(ret)
}

func TestDao_KLine(t *testing.T) {
	d.KLine(&binance.WsKlineEvent{
		Event:  "a",
		Time:   1,
		Symbol: "BTCUSDT",
	})
}
