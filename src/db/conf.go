package db

import "gopkg.in/mgo.v2"

type DevContext struct {
	MgoAddr string
	MgoSessions map[string]*mgo.Session
}

var DevDb = DevContext{
	MgoAddr:"mongodb://wg:wg2019@140.143.234.207:27000",
}