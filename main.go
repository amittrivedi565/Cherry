package main

import (
	"Cherry/builder"
	"Cherry/db"
	"Cherry/dsl"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseRequest struct {
	Op string `json:"op"`
}

func main() {

	db := db.GetInstance()
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
			
		case "insert":
			var obj dsl.Insert
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = builder.BuildInsert(obj)

		case "update":
			var obj dsl.Update
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = builder.BuildUpdate(obj)

		case "delete":
			var obj dsl.Delete
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = builder.BuildDelete(obj)

		case "select":
			var obj dsl.Select
			if err := json.Unmarshal(body, &obj); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			query = builder.BuildSelect(obj)

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
