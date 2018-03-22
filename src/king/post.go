package king

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"net/url"
)

type Server struct {
	StartTime string `json:"s_time"`
	EndTime   string `json:"e_time"`
	CityId    string `json:"city_id"`
}

func AnalogPost(c *gin.Context) {
	curl := "http://dev.jsga.wii.qq.com/show/showapi/get_move_total"
	u, _ := url.Parse(curl)
	q := u.Query()
	q.Set("s_time", "2017-01-01")
	q.Set("e_time", "2018-01-01")
	q.Set("city_id", "0")
	res, err := http.PostForm(curl, q)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
	c.JSON(http.StatusOK, gin.H{
		"data": string(result),
	})
}
