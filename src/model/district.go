package model

type District struct {
	Id          int    `json:"id" form:"id"`
	ParentId    int    `json:"parent_id" form:"parent_id"`
	Code        int    `json:"code" form:"code"`
	Name        string `json:"name" form:"name"`
	LongName    string `json:"long_name" form:"long_name"`
	ShortName   string `json:"short_name" form:"short_name"`
	FirstLetter string `json:"first_letter" form:"first_letter"`
	Type        string `json:"type" form:"type"`
	OmgId       int    `json:"omg_id" form:"omg_id"`
	OmgPid      int    `json:"omg_pid" form:"omg_pid"`
	WeatherId   int    `json:"weather_id" form:"weather_id"`
	Disabled    int    `json:"disabled" form:"disabled"`
}
