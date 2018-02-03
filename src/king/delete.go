package king

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"db"
	"fmt"
	"strconv"
)

/**
数据删除
 */
func Delete(c *gin.Context) {
	cid := c.Query("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("111")
	rs, err := db.DevContext.Db.Exec("DELETE FROM person WHERE id=?", id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
