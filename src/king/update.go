package king

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"model"
	"log"
	"db"
	"fmt"
	"net/http"
)

/**
数据修改
 */
func Update(c *gin.Context) {
	cid := c.Query("id")
	id, err := strconv.Atoi(cid)
	person := model.Person{Id: id}
	err = c.Bind(&person)
	if err != nil {
		log.Fatalln(err)
	}
	stmt, err := db.DevContext.Db.Prepare("UPDATE tbl_person SET first_name=?, last_name=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(person.FirstName, person.LastName, person.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Update person %d successful %d", person.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
