//This package simulates events
//and uses goroutines to listen to them
package main

import (
	"fmt"
)

type Button struct {
	//this will have string keys for the name of the event
	//one key will have slice of string type channels as value that will be listening to that event
	eventListeners map[string][]chan string
}

func main() {

	btn := MakeButton()

	handler1 := make(chan string, 2)
	handler2 := make(chan string, 2)

	btn.AddEventListener("click", handler1)
	btn.AddEventListener("click", handler2)

	go func() {
		for {
			msg := <-handler1
			fmt.Println("HandlerOne: ", msg)
		}
	}()

	go func() {
		for {
			msg := <-handler2
			fmt.Println("HandlerTwo: ", msg)
		}
	}()
	btn.TriggerEvent("click", "Button Clicked!")
	btn.RemoveEventListener("click", handler2)

	btn.TriggerEvent("click", "Button clicked again!")

	fmt.Scanln()
}

func MakeButton() *Button {
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

func (this *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		this.eventListeners[event] = append(this.eventListeners[event], responseChannel)
	} else {
		this.eventListeners[event] = []chan string{responseChannel}
	}
}

func (this *Button) RemoveEventListener(event string, eventListener chan string) {
	if _, present := this.eventListeners[event]; present {
		for idx, _ := range this.eventListeners[event] {
			if this.eventListeners[event][idx] == eventListener {
				this.eventListeners[event] = append(this.eventListeners[event][:idx], this.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (this *Button) TriggerEvent(event string, response string) {
	if _, present := this.eventListeners[event]; present {
		for _, handler := range this.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
