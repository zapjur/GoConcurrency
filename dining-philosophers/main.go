package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", rightFork: 0, leftFork: 1},
	{name: "Aristotle", rightFork: 1, leftFork: 2},
	{name: "Socrates", rightFork: 2, leftFork: 3},
	{name: "Descartes", rightFork: 3, leftFork: 4},
	{name: "Confucius", rightFork: 4, leftFork: 0},
}

var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 1 * time.Second

var orderMutex = &sync.Mutex{}
var orderFinished []string

func main() {

	fmt.Println("Dining Philosophers Problem")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty")

	dine()

	fmt.Println("The table is empty")
	fmt.Printf("The order of philosophers who finished eating: %v\n", orderFinished)

}

func dine() {

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()

}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("%s is seated at the table\n", philosopher.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s has picked up the right fork\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s has picked up the left fork\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s has picked up the left fork\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s has picked up the right fork\n", philosopher.name)
		}

		fmt.Printf("\t%s is eating\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("\t%s has put down the forks\n", philosopher.name)

	}

	fmt.Printf("%s is leaving the table\n", philosopher.name)
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()

}
