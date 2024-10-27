package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {

	rand.Seed(time.Now().UnixNano())

	color.Yellow("The sleeping barber problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := Barbershop{
		ShopCapacity:    seatingCapacity,
		HaircutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}

	color.Green("The barbershop is now open")

	shop.addBarber("Frank")
	shop.addBarber("Gerard")
	shop.addBarber("Tom")
	shop.addBarber("John")
	shop.addBarber("Peter")
	shop.addBarber("Max")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShop()
		closed <- true
	}()

	i := 1

	go func() {
		for {
			randMs := rand.Int() % (2 * arrivalRate)

			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randMs)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	<-closed

	close(shopClosing)
	close(closed)

}
