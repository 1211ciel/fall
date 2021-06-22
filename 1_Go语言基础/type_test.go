package __Go语言基础

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
)

type Animal struct {
	name string
}

func (a *Animal) Move() {
	fmt.Println(a.name, "移动了")
}

type Dog struct {
	Feet int8
	*Animal
}

func TestAoo(t *testing.T) {
	defer ants.Release()
	pool, err := ants.NewPool(10)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(pool)
	var wg sync.WaitGroup
	num := 0
	for i := 0; i < 10000; i++ {
		//pool.Submit(func() {
		wg.Add(1)
		pool.Submit(func() {
			fmt.Println(i)
			wg.Done()
			num++
		})
		//})
	}
	wg.Wait()
	fmt.Println("num:", num)
}
