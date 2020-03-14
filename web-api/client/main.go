package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	start := time.Now()
	for i := 0; i < n; i++ {
		a, _ := http.Get("http://localhost:8090/hello")
		log.Printf("get a: %v", a.Status)
	}
	ms := time.Since(start)
	log.Printf("mil: %v\tmic: %v\tmin:%v\n", ms.Milliseconds(), ms.Microseconds(), ms.Minutes())
	log.Printf("total time: %v\ttotal req: %v\ttime per req:%v\n", ms, n, ms.Microseconds()/int64(n))
}
