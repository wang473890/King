package king

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func AnalogGet(c *gin.Context) {
	u, _ := url.Parse("http://tyn.wiiqq.com/guide/guide/guideapi/guide_list")
	q := u.Query()
	q.Set("size", "8")
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())

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
