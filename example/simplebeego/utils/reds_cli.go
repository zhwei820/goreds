package utils

import (
	"github.com/blainsmith/goreds"
	"github.com/garyburd/redigo/redis"
)

var SearchCli *goreds.Client

func init() {
	redis, _ := redis.DialURL("redis://localhost:6379")
	SearchCli = goreds.NewClient(redis, "namespace")

}
