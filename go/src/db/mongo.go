package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DevMgo struct {
	MgoSession  *mgo.Session
	MgoSessions map[string]*mgo.Session
	MgoDb       string
	MgoTable    string
}

var Mgo = DevMgo{
	MgoDb:    "db_train",
	MgoTable: "tbl_train_daily",
}

type DevDBText struct {
	ProjectId  string `json:"ProjectId,omitempty" bson:"_id"`
	ConnString string `json:"ConnString,omitempty" bson:"connString"`
}

func InitSessions() {

	var err error
	//init mongodb
	Mgo.MgoSession, err = mgo.Dial(DevDb.MgoAddr)
	if err != nil {
		fmt.Printf("mgo.Dial(%s) failed, err=%v", DevDb.MgoAddr, err)
		return
	}
	DevDb.MgoSessions = make(map[string]*mgo.Session)

	InitProMgoSession()

	//go func() {
	//	ticker := time.NewTicker(time.Duration(30*60) * time.Second)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			InitProMgoSession()
	//		}
	//	}
	//}()
}
func InitProMgoSession() {
	newSess := Mgo.MgoSession.Clone()
	defer newSess.Close()

	var projectDb []DevDBText
	var param map[string]interface{}
	param = make(map[string]interface{})
	param["status"] = 1
	param["mode"] = gin.Mode()

	newSess.DB(Mgo.MgoDb).C(Mgo.MgoTable).Find(param).All(&projectDb)

	for i, d := range projectDb {
		id := projectDb[i].ProjectId
		if _, ok := Mgo.MgoSessions[id]; ok {
			continue
		}
		connString := d.ConnString
		var er error
		var gSession *mgo.Session
		gSession, er = mgo.Dial(connString)
		if er != nil {
			fmt.Printf("InitMongoDB | mgo.Dial(%s) failed, err=%v", DevDb.MgoAddr, er)
			return
		}
		DevDb.MgoSessions[id] = gSession
	}
}
