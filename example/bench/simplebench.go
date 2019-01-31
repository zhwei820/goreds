package main

import (
	"bufio"
	"github.com/blainsmith/goreds"
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"strings"
	"time"
)

func readFile(filepath string) []string {
	// 打开将要搜索的文件
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 逐行读入
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		var text string
		data := strings.Split(scanner.Text(), "||||")
		if len(data) != 10 {
			continue
		}
		text = data[9]
		if text != "" {
			lines = append(lines, text)
		}
	}
	return lines
}

func main() {

	redis, _ := redis.DialURL("redis://localhost:6379")
	search := goreds.NewClient(redis, "namespace")
	//lines := readFile("weibo_data.txt")
	//for id, line := range lines {
	//	search.Index(line, strconv.Itoa(id))
	//}
	ids, _ := search.Query("非洲")
	println(len(ids), strings.Join(ids, ", "))
	// 19 837, 9798, 9757, 9207, 8544, 7968, 7779, 5694, 5322, 5208, 5123, 5077, 4959, 4856, 4776, 4582, 4496, 3578, 2682
	start := time.Now()
	for i := 0; i < 10000; i++ {
		_, _ = search.Query("非洲")
	}
	elapsed := time.Since(start)

	println("search time: ")
	println(float64(elapsed)/1.0e9, " s")
	// +1.111377e+000  s

}
