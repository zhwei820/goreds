package goreds // import "github.com/blainsmith/goreds"

import (
	"github.com/blainsmith/goreds/words"
	"github.com/garyburd/redigo/redis"
)

//  CliEn is a goreds client that uses a Redis client and Namespace to perform searches
type ClientEn struct {
	Redis     redis.Conn
	Namespace string
}

// Index will store the `id` within the database and use the `text` as the searchable text
func (client *ClientEn) Index(text, id string) (interface{}, error) {
	w := words.StemsArray(words.StripStopWords(words.Split(text)))
	c := words.Count(w)
	mm := words.MetaphoneMap(w)

	client.Redis.Send("MULTI")
	for i, _ := range mm {
		client.Redis.Send("ZADD", client.Namespace+":word:"+mm[i], c[i], id)
		client.Redis.Send("ZADD", client.Namespace+":object:"+id, c[i], mm[i])
	}

	return client.Redis.Do("EXEC")
}

// Remove will delete the `id` from the database so it is not longer searchable
func (client *ClientEn) Remove(id string) (interface{}, error) {
	reply, err := redis.Strings(client.Redis.Do("ZREVRANGEBYSCORE", client.Namespace+":object:"+id, "+inf", 0))
	if err != nil {
		return nil, err
	}

	client.Redis.Send("MULTI")
	for _, m := range reply {
		client.Redis.Send("DEL", client.Namespace+":object:"+id)
		client.Redis.Send("ZREM", client.Namespace+":word:"+m, id)
	}

	return client.Redis.Do("EXEC")
}

// Query performs a search against the database and returns a slice of ids that match
func (client *ClientEn) Query(text string, operator Operator) ([]string, error) {
	if operator == "" {
		operator = AND
	}

	w := words.StemsArray(words.StripStopWords(words.Split(text)))
	mk := words.MetaphoneKeys(client.Namespace, w)

	lmk := len(mk)
	if lmk <= 0 {
		return nil, nil
	}

	tk := client.Namespace + "temp"
	search := []interface{}{tk, lmk}
	for _, k := range mk {
		search = append(search, k)
	}

	client.Redis.Send("MULTI")
	client.Redis.Send(string(operator), search...)
	client.Redis.Send("ZREVRANGE", tk, 0, -1)
	client.Redis.Send("ZREMRANGEBYRANK", tk, 0, -1)

	values, err := redis.Values(client.Redis.Do("EXEC"))
	if err != nil {
		return nil, err
	}

	var ids []string
	count := 0
	for len(values) > 0 {
		values, err = redis.Scan(values, &count, &ids)
	}

	return ids, nil
}
