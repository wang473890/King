package city

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"common"
	"db"
	"log"
	"math"
)

func CityList(c *gin.Context) {
	//参数获取
	_page := c.DefaultQuery("page", "1")
	_size := c.DefaultQuery("size", "10")
	key := c.Query("key")
	//参数校验
	var page, size int
	//页码校验
	page, err1 := strconv.Atoi(_page)
	if err1 != nil {
		errorData := ""
		common.Return(c, 100000, "page必须是整数", errorData)
		return
	}
	//每页数量校验
	size, err2 := strconv.Atoi(_size)
	if err2 != nil {
		errorData := ""
		common.Return(c, 100001, "size必须是整数", errorData)
		return
	}
	//总数查询
	var sqlAll string = "select count(*) as total from tbl_district where disabled = 0 and (name like '%" + key + "%' or long_name like '%" + key + "%')"
	rows, err3 := db.DevContext.Db.Query(sqlAll)
	if err3 != nil {
		log.Fatalln(err3)
	}
	var total common.Total
	rows.Next()
	rows.Scan(&total.Total)
	var totalPage float32
	totalPage = float32(total.Total) / float32(size)
	var pagination common.Pagination
	pagination.Page = page
	pagination.TotalPage = int(math.Ceil(float64(totalPage)))
	pagination.PageSize = size
	pagination.Total = total.Total
	//mysql语句拼装
	var offset, limit string
	offset = strconv.Itoa(size * (page - 1))
	limit = strconv.Itoa(size)
	var sqlList string = "select id,parent_id,code,name,long_name,short_name,first_letter,type,weather_id from tbl_district where disabled = 0 and (name like '%" + key + "%' or long_name like '%" + key + "%') limit " + offset + "," + limit
	row, err3 := db.DevContext.Db.Query(sqlList)
	if err3 != nil {
		log.Fatalln(err3)
	}
	defer func() {
		row.Close()
	}()
	list := make([]common.CityList, 0)
	for row.Next() {
		var city common.CityList
		row.Scan(&city.Id, &city.ParentId, &city.Code, &city.Name, &city.LongName, &city.ShortName, &city.FirstLetter, &city.Type, &city.WeatherId)
		list = append(list, city)
	}
	if err3 = row.Err(); err3 != nil {
		log.Fatalln(err3)
	}
	var data common.CityReturn
	data.Pagination = pagination
	data.List = list
	common.Return(c, 0, "ok", data)

}
