package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order for pizza #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d will take %d seconds \n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d ***", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** We burned pizza #%d ***", pizzaNumber)
		} else {
			msg = fmt.Sprintf("Pizza #%d is ready!", pizzaNumber)
			success = true
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {

	i := 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func main() {

	rand.Seed(time.Now().UnixNano())

	color.Cyan("The pizza shop is open!")
	color.Cyan("-----------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d completed successfully", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Order #%d failed", i.pizzaNumber)
			}
		} else {
			color.Cyan("We are closing the pizza shop!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing channel: %v", err)
			}
		}
	}

	color.Cyan("-----------------------")
	color.Cyan("Done for the day!")
	color.Cyan("Total pizzas made: %d", pizzasMade)
	color.Cyan("Total pizzas failed: %d", pizzasFailed)
	color.Cyan("Total pizzas: %d", total)

	switch {
	case pizzasFailed > 9:
		color.Red("We need to improve our service!")
	case pizzasFailed > 5:
		color.Red("We need to improve our service!")
	case pizzasFailed > 2:
		color.Yellow("We are doing good but we can do better!")
	default:
		color.Green("We are doing great!")
	}

}
