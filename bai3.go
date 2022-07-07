package main

import (
	
	"log"
	"time"
	"sync"
)

var r = sync.RWMutex{}

func errFunc() {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {  // main() thực chất cũng là 1 goroutine nếu for hơi ít thì nó chạy đến hàm dòng cuối và lỗi sẽ không sảy ra nên phải for nhiều (vì ko dùng time.Sleep để những goroutine kia chạy)
		r.Lock()
		go func() {
			for j := 1; j < 10000; j++ {
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10  // Nhiều luồng cùng ghi 1 giá trị của map thì sẽ lỗi
			}
			r.Unlock()
		}()
	}

	log.Print("done")
	time.Sleep(0 * time.Second)
}

func main() {

	errFunc()
}