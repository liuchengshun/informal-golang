package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dataOrigin []string = []string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
}

var dataObj []string

var wg sync.WaitGroup

var rwmt sync.RWMutex

func main() {

	wg.Add(3)

	go WriteData()
	go ReadData()
	go WriteData()

	wg.Wait()
	fmt.Println("End")
}

func WriteData() {
	rand.Seed(time.Now().Unix())
	defer wg.Done()

	for {
		rwmt.Lock()
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		data := dataOrigin[rand.Intn(len(dataOrigin))]
		dataObj = append(dataObj, data)
		fmt.Println("added a string into dataObj")
		rwmt.Unlock()

		if len(dataObj) == 20 {
			break
		}
	}
}

func ReadData() {
	n := 0
	defer wg.Done()
	
	for {
		rwmt.RLock()
		fmt.Printf("%s", dataOrigin[n])
		n++
	}
}
