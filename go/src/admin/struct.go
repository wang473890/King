package admin

//管理员表
type TplAdmin struct {
	FId         int64  `json:"FId" bson:"FId"'`
	FName       string `json:"FName" bson:"FName"'`
	FPass       string `json:"FPass" bson:"FPass"'`
	FSale       string `json:"FSale" bson:"FSale"'`
	FCreateTime int64  `json:"FCreateTime" bson:"FCreateTime"'`
	FUpdateTime int64  `json:"FUpdateTime" bson:"FUpdateTime"'`
	FLevel      int64  `json:"FLevel" bson:"FLevel"`
}
