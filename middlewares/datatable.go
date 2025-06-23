package middlewares

import (
	"ayam-geprek-backend/config"
	"strconv"
)

func Datatables(sRecursive, sTable, order, sFilter string, limit, offset int) map[string]interface{} {
	switch limit {
	case 10:
		limit = 10
	case 25:
		limit = 25
	case 50:
		limit = 50
	default:
		limit = 100
	}

	var counts []map[string]interface{}

	cte := sRecursive

	stotalRecords := cte + ` select count(*) as total
	from (
			` + sTable + `
	) as b `

	result := config.DB.Raw(stotalRecords).First(&counts)
	total := 0
	if result.Error != nil {
		total = 0
	} else {
		total = int(counts[0]["total"].(int64))
	}

	if order != "" {
		order = " order by " + order
	}

	datas := make(map[string]interface{})
	datas["recordsTotal"] = total

	var data []map[string]interface{}

	query := cte + ` ` + sTable + ` ` + sFilter + order + ` LIMIT ` + strconv.Itoa(limit) + ` OFFSET ` + strconv.Itoa(offset)

	result = config.DB.Raw(query).Scan(&data)
	if result.Error != nil {
		datas["data"] = nil
		return datas
	}
	datas["data"] = data

	var countsfilter []map[string]interface{}
	query = cte + ` select count(*) as total from (` + sTable + ` ` + sFilter + ` ) as b  `
	result = config.DB.Raw(query).First(&countsfilter)
	totalfilters := 0
	if result.Error != nil {
		totalfilters = 0
	} else {
		totalfilters = int(countsfilter[0]["total"].(int64))
	}
	datas["recordsFiltered"] = totalfilters

	return datas
}
