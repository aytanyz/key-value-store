package main

import (
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
}

var DB map[string]Data
var DB_history map[string][]Data

func getValue(context *gin.Context) {
	key := context.Param("key")
	if item, exist := DB[key]; exist {
		if item.Value == "null" {
			context.String(204, "No Content")
		} else {
			context.String(200, item.Value)
		}
	} else {
		context.String(404, "Not Found")
	}
}

func setValue(context *gin.Context) {
	key := context.Param("key")
	value, _ := ioutil.ReadAll(context.Request.Body)
	DB[key] = Data{
		string(value),
		time.Now().UnixNano() / int64(time.Millisecond)}
	// add to histroy
	DB_history[key] = append(
		DB_history[key],
		Data{
			string(value),
			time.Now().UnixNano() / int64(time.Millisecond)})
	context.String(204, "No Content")
}

func deleteValue(context *gin.Context) {
	key := context.Param("key")
	if _, exist := DB[key]; exist {
		DB[key] = Data{
			"null",
			time.Now().UnixNano() / int64(time.Millisecond)}
		// add to histroy
		DB_history[key] = append(
			DB_history[key],
			Data{
				"null",
				time.Now().UnixNano() / int64(time.Millisecond)})
	}
	context.String(204, "No Content")
}

func getHistory(context *gin.Context) {
	key := context.Param("key")
	currTime := time.Now().UnixNano() / int64(time.Millisecond)
	if itemList, exist := DB_history[key]; exist {
		var localHistory []Data
		for i := len(itemList) - 1; i >= 0; i-- {
			if currTime-itemList[i].Timestamp <= 120000 {
				localHistory = append(localHistory, itemList[i])
			}
		}
		context.IndentedJSON(200, localHistory)
	} else {
		context.String(204, "No Content")
	}
}

func main() {
	router := gin.Default()

	DB = make(map[string]Data)
	DB_history = make(map[string][]Data)

	router.GET("/:key", getValue)
	router.PUT("/:key", setValue)
	router.DELETE("/:key", deleteValue)
	router.GET("/:key/history", getHistory)

	router.Run(":3000")
}

// curl -i -X PUT 'http://localhost:3000/foo' -H 'Content-Type: application/octet-stream' --data-binary 'hello world!'
// curl -i 'http://localhost:3000/foo'
// curl -i 'http://localhost:3000/foo/history'
// curl -i -X DELETE 'http://localhost:3000/foo'
// curl -i -X PUT 'http://localhost:3000/boo' -H 'Content-Type: application/octet-stream' --data-binary 'aytan'
