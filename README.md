# goreds

[![Build Status](https://travis-ci.org/blainsmith/goreds.svg?branch=master)](https://travis-ci.org/blainsmith/goreds)
[![GoDoc](https://godoc.org/github.com/blainsmith/goreds?status.svg)](https://godoc.org/github.com/blainsmith/goreds)
![](https://img.shields.io/badge/license-MIT-blue.svg)

goreds is a Go port of [tj/reds](http://github.com/tj/reds) for Node.js

## About

via [https://github.com/tj/reds#about](https://github.com/tj/reds#about)

Currently reds strips stop words and applies the metaphone and porter stemmer algorithms to the remaining words before mapping the constants in Redis sets. For example the following text:

    Tobi is a ferret and he only wants four dollars

Converts to the following constant map:

```go
map[Tobi:TB ferret:FRT wants:WNTS four:FR dollars:DLRS]
```

This also means that phonetically similar words will match, for example "stefen", "stephen", "steven" and "stefan" all resolve to the constant "STFN". Reds takes this further and applies the porter stemming algorithm to "stem" words, for example "counts", and "counting" become "count".

Consider we have the following bodies of text:

    Tobi really wants four dollars

    For some reason tobi is always wanting four dollars

The following search query will then match _both_ of these bodies, and "wanting", and "wants" both reduce to "want".

    tobi wants four dollars

## Example

```go
redis, _ := redis.DialURL("redis://localhost:6379")
search := goreds.NewClient(redis, "namespace")

// index some text and assign to their id
search.Index("example index text", "1")
search.Index("example text being indexed for the sake of this example", "2")
search.Index("example testing with a lot more text", "3")
search.Index("index some more sample text data", "4")

// remove an item from the search index
search.Remove("2")

// query the search index that should match ids 1 and 4
ids, _ := search.Query("index text", goreds.AND)
fmt.Println(ids)
// Output: [4 1]
```
