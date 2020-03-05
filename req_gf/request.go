package req

import (
	"encoding/json"
	"github.com/chenyingdi/my-toolkits/data_gf/v1"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GetArgs(r *ghttp.Request) (interface{}, error) {
	var args = r.GetMap()
	var sqlArgs data.Args

	for k, v := range args {
		switch k {
		case "type":
			sqlArgs.Type = v.(string)

		case "where":
			where := g.Map{}

			err := json.Unmarshal([]byte(v.(string)), &where)
			if err != nil {
				return nil, err
			}

			sqlArgs.Where = where

		case "or":
			or := g.Map{}

			err := json.Unmarshal([]byte(v.(string)), &or)
			if err != nil {
				return nil, err
			}

			sqlArgs.Or = or

		case "page":
			sqlArgs.Page = v.(int)

		case "fields":
			sqlArgs.Fields = v.(string)

		case "fieldsEx":
			sqlArgs.FieldsEx = v.(string)

		case "limit":
			sqlArgs.Limit = v.(int)

		case "dataPerPage":
			sqlArgs.DataPerPage = v.(int)
		}

	}

	return sqlArgs, nil
}
