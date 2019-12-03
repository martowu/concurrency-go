package main

import (
	"fmt"
	"sync"
	"time"
)

type OpenlogiItem struct {
	ItemID    string
	AccountID string
}

func main() {
	linkingItems := []OpenlogiItem{
		{
			AccountID: "A001",
			ItemID:    "I001",
		},
		{
			AccountID: "A001",
			ItemID:    "I002",
		},
		{
			AccountID: "A001",
			ItemID:    "I003",
		},
		{
			AccountID: "A001",
			ItemID:    "I004",
		},
		{
			AccountID: "B001",
			ItemID:    "I001",
		},
		{
			AccountID: "V001",
			ItemID:    "I001",
		},
	}

	accountItems := make(map[string][]string, len(linkingItems))
	// map split user
	for _, item := range linkingItems {
		accountItems[item.AccountID] = append(accountItems[item.AccountID], item.ItemID)
	}

	// fmt.Println(accountItems)
	targetID := make(chan string, 10)
	// defer close(targetID)

	wg1 := sync.WaitGroup{}
	wg := sync.WaitGroup{}

	for accountID, itemIDs := range accountItems {
		wg1.Add(1)
		accountID := accountID
		itemIDs := itemIDs
		go func() {
			defer wg1.Done()
			fmt.Println("get user by accountID: ", accountID)
			fmt.Println("auth user by access key")
			for _, i := range itemIDs {
				i := i
				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(time.Second * 1)
					fmt.Println("fetchOpenlogiItem by itemID: ", i)
					targetID <- i
				}()
			}
		}()
	}

	go func() {
		wg1.Wait()
		wg.Wait()
		close(targetID)
	}()
	for a := range targetID {
		fmt.Println(a)
	}

	// userItems := make(map[string][]OpenlogiItem)
	// userItems["A001"] = append(userItems["A001"], OpenlogiItem{ItemID: "I001"})
	// userItems["A001"] = append(userItems["A001"], "I002")
	// userItems["A002"] = append(userItems["A002"], "I001")
	// userItems["A002"] = append(userItems["A002"], "I002")
	// fmt.Println(userItems)
	// for user, items := range userItems {
	// 	fmt.Println(user, items)
	// }
}
