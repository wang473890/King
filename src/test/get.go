package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"io/ioutil"
	"strings"
	"db"
	"encoding/json"
)

type TrainData struct {
	Flag   string            `json:"flag" bson:"flag"`
	Map    map[string]string `json:"map" bson:"map"`
	Result []string          `json:"result" bson:"result"`
}

type TrainReturn struct {
	Data       TrainData `json:"data" bson:"data"`
	HttpStatus int64     `json:"httpstatus" bson:"httpstatus"`
	Messages   string    `json:"messages" bson:"messages"`
	Status     bool      `json:"status" bson:"status"`
}

type TrainMgoData struct {
	TrainNum   string `json:"train_num" bson:"train_num"`
	TrainBegin string `json:"train_begin" bson:"train_begin"`
	TrainEnd   string `json:"train_end" bson:"train_end"`
	TimeBegin  string `json:"time_begin" bson:"time_begin"`
	TimeEnd    string `json:"time_end" bson:"time_end"`
	TimeLong   string `json:"time_long" bson:"time_long"`
	Spec       string `json:"special" bson:"special"`
	Fir        string `json:"first" bson:"first"`
	Sec        string `json:"second" bson:"second"`
	HighLie    string `json:"high_lie" bson:"high_lie"`
	SoftLie    string `json:"soft_lie" bson:"soft_lie"`
	StillLie   string `json:"still_lie" bson:"still_lie"`
	HardLie    string `json:"hard_lie" bson:"hard_lie"`
	SoftSeat   string `json:"soft_seat" bson:"soft_seat"`
	HardSeat   string `json:"hard_seat" bson:"hard_seat"`
	Stand      string `json:"stand" bson:"stand"`
	CreateTime int64  `json:"create_time" bson:"create_time"`
	UpdateTime int64  `json:"update_time" bson:"update_time"`
}

func GetData(c *gin.Context) {
	TrainDate := c.DefaultQuery("train_date", (time.Now().Format("2006-01-02")))
	_, e := time.Parse("2006-01-02", TrainDate)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":       "时间格式错误",
			"train_time": TrainDate,
			"err":        e,
		})
		return
	}
	FStation := c.DefaultQuery("from_station", "BJP")
	TStation := c.DefaultQuery("to_station", "HBB")
	curl := "https://kyfw.12306.cn/otn/leftTicket/queryO?"
	curl = curl + "leftTicketDTO.train_date=" + TrainDate
	curl = curl + "&leftTicketDTO.from_station=" + FStation
	curl = curl + "&leftTicketDTO.to_station=" + TStation
	curl = curl + "&purpose_codes=" + "ADULT"
	resp, e := http.Get(curl)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "请求错误",
			"resp": resp,
			"err":  e,
		})
		return
	}
	res, e := ioutil.ReadAll(resp.Body)
	var data TrainReturn
	e = json.Unmarshal(res, &data)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "接受错误",
			"data": data,
			"err":  e,
		})
		return
	}
	resp.Body.Close()
	defer resp.Body.Close()
	l := len(data.Data.Result)
	a := make([]string, l)
	var b []string
	mgoData := make([]TrainMgoData, l)
	for i := 0; i < l; i++ {
		a[i] = data.Data.Result[i]
		b = strings.Split(a[i], "|")
		mgoData[int(i)].TrainNum = b[3]
		mgoData[int(i)].TrainBegin = b[4]
		mgoData[int(i)].TrainEnd = b[5]
		mgoData[int(i)].TimeBegin = b[8]
		mgoData[int(i)].TimeEnd = b[9]
		mgoData[int(i)].TimeLong = b[10]
		mgoData[int(i)].Stand = b[26]
		mgoData[int(i)].HardSeat = b[29]
		mgoData[int(i)].CreateTime = time.Now().Unix()
		mgoData[int(i)].UpdateTime = time.Now().Unix()
	}
	for i := 0; i < l; i++ {
		e = db.Mgo.MgoSession.DB(db.Mgo.MgoDb).C(db.Mgo.MgoTable).Insert(&mgoData[i])
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": "存储错误",
				"mgo":  mgoData[i],
				"err":  e,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "ok",
		"data": mgoData,
		"err":  e,
	})
}
