package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	// v0が10個確保されて11個目以降にappendされる
	// slice := make([]int, 10)
	// slice := make([]int, 0)
	// slice := []int{}
	slice := make([]string, 10000)
	// capacityをとってlenを0にしておけば10個の0が入ることはない。
	slice2 := make([]int, 0, 10000)
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for i := 0; i < 10000; i++ {
		wg.Add(2)
		i := i
		go appendSlice(&wg, slice, i)
		go func() {
			defer wg.Done()
			m.Lock()
			// 10000個の並行処理が、1つの共有のリソースに対して、配列を読み込む、変数を読み込む、新しい配列の作成、既存の配列への代入を行っている。appendはアトミックな操作ではないためデータ壊れる可能性がある。Lockして防ぐ
			slice2 = append(slice2, i)
			m.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(len(slice))
	fmt.Println(len(slice2))
}

func appendSlice(wg *sync.WaitGroup, slice []string, i int) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	slice[i] = strconv.Itoa(i)
	// m.Lock()
	// slice = append(slice, strconv.Itoa(i))
	// m.Unlock()
}
