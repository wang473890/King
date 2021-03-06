package train

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"time"
	"net/http"
	"db"
)

type PostData struct {
	Data       interface{} `json:"data" bson:"data"`
	ErrorStr string `json:"str" bson:"str"`
	CreateTime int64       `json:"create_time" bson:"create_time"`
}

func PostRow(c *gin.Context) {
	buf := make([]byte, 1024*1024*1024)
	//参数获取json
	n, _ := c.Request.Body.Read(buf)
	var data interface{}
	var postData PostData

	if e := json.Unmarshal([]byte(string(buf[0:n])), &data); e != nil {
		//c.JSON(http.StatusOK, gin.H{
		//	"code":  1000,
		//	"msg":   "post fail",
		//	"error": e,
		//	"data":  string(buf[0:n]),
		//})
		//return
		postData.ErrorStr = string(buf[0:n])
	} else {
		postData.Data = data
	}
	postData.CreateTime = time.Now().Unix()
	if e := db.Mgo.MgoSession.DB("db_train").C("house_task").Insert(&postData); e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  1001,
			"msg":   "storage fail",
			"error": e,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": postData,
	})
}
