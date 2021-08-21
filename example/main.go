package main

import (
	"fmt"
	"time"

	"github.com/spksoft/pipeline"
)

func generateData(input <-chan interface{}) <-chan interface{} {
	output := make(chan interface{})
	go func() {
		number := 0
		for {
			output <- number
			number++
			time.Sleep(5 * time.Second)
		}
	}()
	return output
}

func prefixData(input <-chan interface{}) <-chan interface{} {
	output := make(chan interface{})
	go func() {
		for data := range input {
			output <- fmt.Sprintf("prefix: %d", data.(int))
		}
	}()
	return output
}

func main() {
	p := pipeline.New("example")
	p.RegisterProcessor(generateData)
	p.RegisterProcessor(prefixData)
	c := p.Run()
	for data := range c {
		fmt.Println(data)
	}
}
