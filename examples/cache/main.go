package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/finalist736/gokit/cache/ramcache"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Int63()

	rndString := strconv.FormatInt(rnd, 10)

	key := fmt.Sprintf("%x", sha1.Sum([]byte(rndString)))
	fmt.Printf("key: %s; value: %d\n", key, rnd)

	cs := ramcache.New(time.Second*10, time.Second*15)

	cs.Set(key, rnd)

	time.Sleep(time.Second * 8)
	item, ok := cs.Get(key)
	if ok {
		itemInt := item.(int64)
		fmt.Printf("getted item: %d\n", itemInt)
	} else {
		fmt.Printf("no item in cache!\n")
	}

	time.Sleep(time.Second * 5)
	item, ok = cs.Get(key)
	if !ok {
		fmt.Println("item removed!")
	} else {
		fmt.Printf("error! item not removed!\n")
	}
}
