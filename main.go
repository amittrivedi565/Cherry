package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OperationName string

type BaseRequest struct {
	Op OperationName `json:"op"`
}

func main() {

	db := GetInstance()

	r := gin.Default()

	r.POST("/exec", func(c *gin.Context) {

		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var base BaseRequest
		if err := json.Unmarshal(body, &base); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var query string

		switch base.Op {

		case "create":
			var obj Create
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = BuildCreate(obj)

		case "insert":
			var obj Insert
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = BuildInsert(obj)

		case "update":
			var obj Update
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = BuildUpdate(obj)

		case "delete":
			var obj Delete
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = BuildDelete(obj)

		case "select":
			var obj Select
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = BuildSelect(obj)

		default:
			c.JSON(400, gin.H{"error": "invalid operation"})
			return
		}

		data, err := db.ExecuteQuery(string(base.Op), query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"data": data,
		})
	})

	r.Run(":8080")
}
