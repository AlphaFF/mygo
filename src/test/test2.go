/*
* @Author: wangfeng1
* @Date:   2019-04-08 17:37:38
* @Last Modified by:   wangfeng1
* @Last Modified time: 2019-04-09 09:37:49
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world.")
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)

	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item
			}
		}
	}(theMine)

	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel
			fmt.Println("From Finder:", foundOre)
			minedOreChan <- "mineOre"
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan
			fmt.Println("From Miner:", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	<-time.After(time.Second * 5)
}
