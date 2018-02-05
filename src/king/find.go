package king

import (
	"github.com/gin-gonic/gin"
	"model"
	"db"
	"log"
	"net/http"
)

/**
单数据查询
 */
func Find(c *gin.Context) {
	id := c.Query("id")
	var person model.Person
	err := db.DevContext.Db.QueryRow("SELECT id, first_name, last_name FROM tbl_person WHERE id=?", id).Scan(
		&person.Id, &person.FirstName, &person.LastName,
	)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"person": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}
