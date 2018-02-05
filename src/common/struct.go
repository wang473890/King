package common

import "database/sql"

//城市列表
type CityList struct {
	Id          int    `json:"id" form:"id"`
	ParentId    int    `json:"parent_id" form:"parent_id"`
	Code        int    `json:"code" form:"code"`
	Name        string `json:"name" form:"name"`
	LongName    string `json:"long_name" form:"long_name"`
	ShortName   string `json:"short_name" form:"short_name"`
	FirstLetter string `json:"first_letter" form:"first_letter"`
	Type        string `json:"type" form:"type"`
	WeatherId   int    `json:"weather_id" form:"weather_id"`
}

//页码返回
type Pagination struct {
	Page      int `json:"page" form:"page"`
	PageSize  int `json:"size" form:"size"`
	TotalPage int `json:"total_page" form:"total_page"`
	Total     int `json:"total" form:"total"`
}

type CityReturn struct {
	List       []CityList
	Pagination Pagination `json:"pagination"`
}

type MysqlDb struct {
	db *sql.DB //定义结构体
}
type Total struct {
	Total int `json:"total" form:"total"`
}
