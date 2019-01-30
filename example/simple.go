package main

import (
	"fmt"
	"github.com/blainsmith/goreds"
	"github.com/garyburd/redigo/redis"
)

func main() {
	redis, _ := redis.DialURL("redis://localhost:6379")
	search := goreds.NewClient(redis, "namespace")

	// index some text and assign to their id
	search.Index("example index text", "1")
	search.Index("example text being indexed for the sake of this example", "2")
	search.Index("阅读、发现和分享：8小时外的健康生活！example testing with a lot more text", "3")
	search.Index("日本双语幼儿园「Kids Duo International Center index text 南」导入了俄罗斯Universal Terminal Systems开发的AR沙场「iSandBOX」，今年4月AR沙场将作为学校的一部分课程计划开始实施。", "4")

	// remove an item from the search index
	search.Remove("2")

	// query the search index that should match ids 1 and 4
	ids, _ := search.Query("index text", goreds.AND)
	fmt.Println(ids)
	// Output: [4 1]
	ids, _ = search.Query("幼儿园", goreds.AND)
	fmt.Println(ids)
	// Output: [4 1]
}
