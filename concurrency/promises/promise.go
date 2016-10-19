//This package simulates promises
package main

import (
	"errors"
	"fmt"
	"time"
)

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

type PurchaseOrder struct {
	Number int
	Value  float64
}

func main() {
	po := new(PurchaseOrder)
	po.Value = 78.34

	SavePO(po, false).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrder)
		fmt.Printf("Purchase Order saved with id: '%d'\n", po.Number)
		return nil
	}, func(err error) {
		fmt.Printf("Failed to save purchase order: ", err.Error()+"\n")
	}).Then(func(obj interface{}) error {
		fmt.Println("Second promise success")
		return nil
	}, func(err error) {
		fmt.Println("Second promise failure: " + err.Error())
	})
	fmt.Scanln()
}

func SavePO(po *PurchaseOrder, shouldFail bool) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		time.Sleep(3 * time.Second)
		if shouldFail {
			result.failureChannel <- errors.New("The Po saving is Unsuccessful")
		} else {
			po.Number = 1234
			result.successChannel <- po
		}
	}()
	return result
}

func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	//timeout := time.After(1 * time.Second)

	go func() {
		select {
		case obj := <-this.successChannel:
			newErr := success(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-this.failureChannel:
			failure(err)
			result.failureChannel <- err
			//case <-timeout:
			//failure(errors.New("Promise time out"))
		}
	}()
	return result
}
