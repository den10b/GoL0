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

func ordersCacheInit() error {
	orders, err := DB.GetAllOrders()
	if err != nil {
		return err
	}
	for _, myOrder := range orders {
		err := ordersCache.Set(myOrder.Id.String(), myOrder)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetOrders() (map[interface{}]interface{}, error) {
	all := ordersCache.GetALL(false)
	return all, nil
}
func getOrder(orderId string) (interface{}, error) {
	return ordersCache.Get(orderId)
}

func CacheInit() {
	itemsCache = gcache.New(1000).ARC().Build()
	ordersCache = gcache.New(1000).ARC().Build()
	err := ordersCacheInit()
	if err != nil {
		return
	}
	kk, _ := ordersCache.Get("0e37df36-f698-11e6-8dd4-cb9ced3df976")
	fmt.Println(kk)
	deliveryCache = gcache.New(1000).ARC().Build()
	paymentCache = gcache.New(1000).ARC().Build()

}
func TestCache() {
	value1, err := itemsCache.Get("key")
	if err != nil {
		panic(err)
	}
	all := ordersCache.GetALL(false)
	for _, myOrder := range all {
		fmt.Println(myOrder)
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
	fmt.Println("Get:", value3)
	fmt.Println("Get:", value4)
}
