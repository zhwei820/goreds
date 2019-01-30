package goreds // import "github.com/blainsmith/goreds"

import (
	"github.com/garyburd/redigo/redis"
	"github.com/yanyiwu/gojieba"
)

const (
	AND Operator = "zinterstore"
	OR  Operator = "zunionstore"
)

type Operator string

//  Client is a goreds client that uses a Redis client and Namespace to perform searches
type Client struct {
	CliCh *ClientCh
	CliEn *ClientEn
}

// NewClient will create a new search client with the given redigo connection and namespace (defaults to "goreds").
// You may create multiple clients with different namespaces for separate search indexes.
func NewClient(redis redis.Conn, namespace string) *Client {
	if namespace == "" {
		namespace = "goreds"
	}

	return &Client{
		&ClientCh{redis, namespace, gojieba.NewJieba()},
		&ClientEn{redis, namespace}}
}

// Index will store the `id` within the database and use the `text` as the searchable text
func (client *Client) Index(text, id string) (interface{}, error) {
	if ContainsCh(text) { // 汉字
		client.CliEn.Index(SplitEnFromCh(text), id)
		return client.CliCh.Index(text, id)
	}
	return client.CliEn.Index(text, id)

}

// Remove will delete the `id` from the database so it is not longer searchable
func (client *Client) Remove(id string) (interface{}, error) {
	reply, err := redis.Strings(client.CliEn.Redis.Do("ZREVRANGEBYSCORE", client.CliEn.Namespace+":object:"+id, "+inf", 0))
	if err != nil {
		return nil, err
	}

	client.CliEn.Redis.Send("MULTI")
	for _, m := range reply {
		client.CliEn.Redis.Send("DEL", client.CliEn.Namespace+":object:"+id)
		client.CliEn.Redis.Send("ZREM", client.CliEn.Namespace+":word:"+m, id)
	}

	return client.CliEn.Redis.Do("EXEC")
}

// Query performs a search against the database and returns a slice of ids that match
func (client *Client) Query(text string) ([]string, error) {
	if ContainsCh(text) { // 汉字
		return client.CliCh.Query(text, AND)
	}
	return client.CliEn.Query(text, AND)

}
