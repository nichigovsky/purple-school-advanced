package main

import (
	"fmt"
	"math/rand"
)

func double(num int, res chan<- int) {
	res <- num * num
}

func main(){
	numValues := 10

	res := make(chan int)

	go func(){
		sli := make([]int, 0, numValues)

		for range numValues {
			randV := rand.Intn(100)
			sli = append(sli, randV)
		}

		for _, el := range sli {
			go double(el, res)
		}
	}()

	for range numValues {
		num := <-res
		fmt.Println(num)
	}
}