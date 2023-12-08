package rpc

import (
	"context"
	//alertmanagerpb "github.com/go-bamboo/api/alertmanager"
	//uuidpb "github.com/go-bamboo/api/uuid"
	"github.com/go-bamboo/pkg/registry"
)

type Cache struct {
	//uuidc         uuidpb.UUIDClient
	//alertmanagerc alertmanagerpb.AlertManagerClient
}

func NewCache(dis registry.Discovery) (uc *Cache, err error) {
	//uuidc := uuidpb.MustNew(context.TODO(), time.Minute, dis)
	//alertmanagerc := alertmanagerpb.MustNew(context.TODO(), time.Minute, dis)
	uc = &Cache{
		//uuidc:         uuidc,
		//alertmanagerc: alertmanagerc,
	}
	return
}

func (c *Cache) ID(ctx context.Context) (uint64, error) {
	//id, err := c.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	//if err != nil {
	//	return 0, err
	//}
	//return id.Id, nil
	return 0, nil
}

//func (c *Cache) Send(ctx context.Context, in *alertmanagerpb.NotifyRequest) error {
//_, err := c.alertmanagerc.Notify(ctx, in)
//if err != nil {
//	return err
//}
//	return nil
//}
