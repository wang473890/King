package king

import (
	"db"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"model"
)

func FindAll(c *gin.Context) {
	rows, err := db.DevContext.Db.Query("SELECT id, first_name, last_name FROM person")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	persons := make([]model.Person, 0)

	for rows.Next() {
		var person model.Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
