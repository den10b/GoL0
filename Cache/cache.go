package Cache

import (
	"GoL0/DB"
	"github.com/bluele/gcache"
)

var myCache gcache.Cache

func CacheInit() error {
	myCache = gcache.New(1000).ARC().Build()
	orders, err := DB.GetAllOrders()
	if err != nil {
		return err
	}
	for _, myOrder := range orders {
		err := myCache.Set(myOrder.Id.String(), myOrder)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetOrders() (map[interface{}]interface{}, error) {
	all := myCache.GetALL(false)
	return all, nil
}
func GetOrder(orderId string) (interface{}, error) {
	return myCache.Get(orderId)
}
func AddOrder(order DB.Orders) error {
	return myCache.Set(order.Id.String(), order)
}

//func TestCache() {
//	value1, err := itemsCache.Get("key")
//	if err != nil {
//		panic(err)
//	}
//	all := ordersCache.GetALL(false)
//	for _, myOrder := range all {
//		fmt.Println(myOrder)
//	}
//	value3, err := deliveryCache.Get("key")
//	if err != nil {
//		panic(err)
//	}
//	value4, err := paymentCache.Get("key")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Get:", value1)
//	fmt.Println("Get:", value3)
//	fmt.Println("Get:", value4)
//}
