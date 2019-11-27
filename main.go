package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// v0が10個確保されて11個目以降にappendされる
	// slice := make([]int, 10)
	// slice := make([]int, 0)
	// slice := []int{}
	slice := make([]string, 10)
	wg := sync.WaitGroup{}
	// m := sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			slice[i] = strconv.Itoa(i)
			// m.Lock()
			// slice = append(slice, strconv.Itoa(i))
			// m.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(slice)
}

/**
appendを失敗させる
配列を用意する
sync.Waitを用意する
forを10回回す
group.Addする
appendをする
group.Doneをする
group.Waitする。
*/
