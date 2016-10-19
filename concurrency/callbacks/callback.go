//This package is for simulating callbacks with go channels
package main

import (
	"fmt"
)

type PurchaseOrder struct {
	number int
	value  float64
}

func main() {
	po := new(PurchaseOrder)
	po.value = 56.67
	ch := make(chan *PurchaseOrder)

	go SavePO(po, ch)
	newPO := <-ch

	fmt.Printf("New Purchase Order: %v", newPO)

}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.number = 1234
	callback <- po
}
