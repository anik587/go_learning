package concurrency

import (
	"fmt"
	"sync"
)

var score = []int{
	0,
}

var wgrp = &sync.WaitGroup{}
var mutx = &sync.Mutex{}

func RaceCondition() {
	wgrp.Add(4)
	go func(wgr *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("From 1")
		mutx.Lock()
		score = append(score, 1)
		mutx.Unlock()
		wgrp.Done()
	}(wgrp, mutx)

	go func(wgr *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("From 2")
		mutx.Lock()
		score = append(score, 2)
		mutx.Unlock()
		wgrp.Done()
	}(wgrp, mutx)

	go func(wgr *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("From 3")
		mutx.Lock()
		score = append(score, 3)
		mutx.Unlock()
		wgrp.Done()
	}(wgrp, mutx)

	go func(wrg *sync.WaitGroup, m *sync.Mutex,) {
		fmt.Println("From 4")
		mutx.Lock()
		score = append(score, 4)
		mutx.Unlock()
		wgrp.Done()
	}(wgrp, mutx)

	wgrp.Wait()
	fmt.Println(score)
}
