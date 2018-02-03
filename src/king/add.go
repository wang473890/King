package king

import (
	"github.com/gin-gonic/gin"
	"db"
	"log"
	"fmt"
	"net/http"
)

/**
插入数据
 */
func Add(c *gin.Context) {
	firstName := c.DefaultQuery("first_name", "Guest")
	lastName := c.Query("last_name")
	//c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	rs, err := db.DevContext.Db.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", firstName, lastName)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("insert person Id {}", id)
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
