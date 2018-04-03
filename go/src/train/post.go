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
	CreateTime int64       `json:"create_time" bson:"create_time"`
}

func PostRow(c *gin.Context) {
	buf := make([]byte, 90000000)
	//参数获取json
	n, _ := c.Request.Body.Read(buf)
	var data interface{}
	if e := json.Unmarshal([]byte(string(buf[0:n])), &data); e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  1000,
			"msg":   "post fail",
			"error": e,
			"data":  string(buf[0:n]),
		})
		return
	}
	var postData PostData
	postData.Data = data
	postData.CreateTime = time.Now().Unix()
	tbs := "house_task"
	if e := db.Mgo.MgoSession.DB(db.Mgo.MgoDb).C(tbs).Insert(&postData); e != nil {
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
