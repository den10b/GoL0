package Cache

import (
	"GoL0/DB"
	"fmt"
	"github.com/bluele/gcache"
)

var itemsCache gcache.Cache
var ordersCache gcache.Cache
var deliveryCache gcache.Cache
var paymentCache gcache.Cache

func itemsCacheLoad(key interface{}) (interface{}, error) {
	fmt.Println("key:", key)
	items := `[{"uid":"ab0fe0e7-7819-4001-afb0-39c93da1c961","order_uid":"0b35dd36-f698-11e6-8dd4-ca8ced2df976","chrt_id":12,"track_number":"BAB22","price":48,"rid":"vava","name":"jacket","sale":222,"size":"13s","total_price":450,"nm_id":1121,"brand":"loiVeton","status":1},{"uid":"b97d5fa5-2600-4127-b45e-b422c8f495c6","order_uid":"0b35dd36-f698-11e6-8dd4-ca8ced2df976","chrt_id":12,"track_number":"AFK33","price":35,"rid":"chaha","name":"Shoes","sale":333,"size":"12a","total_price":400,"nm_id":1232,"brand":"guci","status":1}]`
	return items, nil
}
func ordersCacheInit() error {
	orders, err := DB.GetAllOrders()
	if err != nil {
		return err
	}
	for _, myOrder := range orders {
		//mo, _ := json.Marshal(myOrder)
		err := ordersCache.Set(myOrder.Id.String(), myOrder)
		if err != nil {
			return err
		}
	}
	return nil
}
func ordersCacheLoad(key interface{}) (interface{}, error) {
	fmt.Println("key:", key)
	return "value", nil
}
func deliveryCacheLoad(key interface{}) (interface{}, error) {
	fmt.Println("key:", key)
	return "value", nil
}
func paymentCacheLoad(key interface{}) (interface{}, error) {
	fmt.Println("key:", key)
	return "value", nil
}
func CacheInit() {
	itemsCache = gcache.New(1000).ARC().LoaderFunc(itemsCacheLoad).Build()
	ordersCache = gcache.New(1000).ARC().Build()
	err := ordersCacheInit()
	if err != nil {
		return
	}
	kk, _ := ordersCache.Get("0e37df36-f698-11e6-8dd4-cb9ced3df976")
	fmt.Println(kk)
	deliveryCache = gcache.New(1000).ARC().LoaderFunc(deliveryCacheLoad).Build()
	paymentCache = gcache.New(1000).ARC().LoaderFunc(paymentCacheLoad).Build()
	all := ordersCache.GetALL(false)
	for _, myOrder := range all {
		fmt.Println(myOrder)
	}
}
func TestCache() {
	value1, err := itemsCache.Get("key")
	if err != nil {
		panic(err)
	}
	value2, err := ordersCache.Get("key")
	if err != nil {
		panic(err)
	}
	value3, err := deliveryCache.Get("key")
	if err != nil {
		panic(err)
	}
	value4, err := paymentCache.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value1)
	fmt.Println("Get:", value2)
	fmt.Println("Get:", value3)
	fmt.Println("Get:", value4)
}
