package train

import (
	"net/http"
	"time"
	"io/ioutil"
	"strings"
	"db"
	"encoding/json"
	"os"
	"fmt"
	"log"
)

type Data struct {
	Flag   string            `json:"flag" bson:"flag"`
	Map    map[string]string `json:"map" bson:"map"`
	Result []string          `json:"result" bson:"result"`
}

type ReturnData struct {
	Data       Data `json:"data" bson:"data"`
	HttpStatus int64     `json:"httpstatus" bson:"httpstatus"`
	Messages   string    `json:"messages" bson:"messages"`
	Status     bool      `json:"status" bson:"status"`
}

type MgoData struct {
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

func TimeTask() {
	//新建日志文件
	logfile, err := os.OpenFile("/data/code/website/king/logs/time_task.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	var TrainDate, FStation, TStation string
	//拼接访问url
	TrainDate = time.Now().Format("2006-01-02")
	FStation = "BJP"
	TStation = "HBB"
	curl := "https://kyfw.12306.cn/otn/leftTicket/queryO?"
	curl = curl + "leftTicketDTO.train_date=" + TrainDate
	curl = curl + "&leftTicketDTO.from_station=" + FStation
	curl = curl + "&leftTicketDTO.to_station=" + TStation
	curl = curl + "&purpose_codes=" + "ADULT"
	logger.Println("***************************获取数据路径: " + curl + "***************************")
	resp, e := http.Get(curl)
	if e != nil {
		logger.Println("***************************请求错误***************************")
		return
	}
	res, e := ioutil.ReadAll(resp.Body)
	var data ReturnData
	e = json.Unmarshal(res, &data)
	if e != nil {
		logger.Println("***************************接收错误***************************")
		return
	}
	resp.Body.Close()
	defer resp.Body.Close()
	l := len(data.Data.Result)
	a := make([]string, l)
	var b []string
	mgoData := make([]MgoData, l)
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
	logger.Println("***************************", mgoData, "***************************")
	for i := 0; i < l; i++ {
		e = db.Mgo.MgoSession.DB(db.Mgo.MgoDb).C(db.Mgo.MgoTable).Insert(&mgoData[i])
		if e != nil {
			logger.Println("***************************存储错误***************************")
			return
		}
	}
	logger.Println("***************************存储成功***************************")
}
