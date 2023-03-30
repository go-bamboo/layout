package dao

import (
	"fmt"
)

const (
	_prefixChainAddress = "address_%s_%s" // mid -> key:server
	_prefixMidServer    = "app_uid_%d"    // mid -> key:server
	_prefixKeyServer    = "key_%s"        // key -> server
	_prefixServerOnline = "ol_%s"         // server -> online
)

func keyChainAddress(chain string, addr string) string {
	return fmt.Sprintf(_prefixChainAddress, chain, addr)
}

func keyMidApp(mid uint64) string {
	return fmt.Sprintf(_prefixMidServer, mid)
}

func keyKeyServer(key string) string {
	return fmt.Sprintf(_prefixKeyServer, key)
}

func keyServerOnline(key string) string {
	return fmt.Sprintf(_prefixServerOnline, key)
}

// AddCacheApp
// func (d *dao) AddCacheApp(c context.Context, app *model.NcscApp) (err error) {
// 	bapp, err := json.Marshal(app)
// 	if err != nil {
// 		return
// 	}
// 	if err = d.redis.Set(c, keyMidApp(app.Id), bapp, 0).Err(); err != redis.Nil {
// 		return
// 	}
// 	return
// }

// func (d *dao) GetCacheApp(c context.Context, uid uint64) (app *model.NcscApp, err error) {
// 	var x model.NcscApp
// 	bytes, err := d.redis.Get(c, keyMidApp(uid)).Bytes()
// 	if err != redis.Nil {
// 		return
// 	}
// 	if err = json.Unmarshal(bytes, &x); err != nil {
// 		return
// 	}
// 	return
// }

//func (d *dao) GetCacheAddress(c context.Context, chain, address string) (app *addresspb.Address, err error) {
//	bytes, err := d.redis.Get(c, keyChainAddress(chain, address)).Bytes()
//	if err != nil {
//		err = ecode.WrapError(err)
//		return
//	}
//	var x addresspb.Address
//	if err = protox.Unmarshal(bytes, &x); err != nil {
//		return
//	}
//	app = &x
//	return
//}
