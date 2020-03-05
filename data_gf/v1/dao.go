package data

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type Condition struct {
	Expression string
	Value      interface{}
	Type       string
}

type Args struct {
	Type        string
	Where       g.Map
	Or          g.Map
	Page        int
	Order       string
	Limit       int
	DataPerPage int
	Fields      string
	FieldsEx    string
}

type DAO struct {
}

func (d *DAO) Get(args Args, model Model) (interface{}, error) {
	m := g.DB().Table(model.TableName())

	// where 条件
	if args.Where != nil {
		for k, v := range args.Where {
			m = m.Where(k, v)
		}
	}

	// or 条件
	if args.Or != nil {
		for k, v := range args.Or {
			m = m.Or(k, v)
		}
	}

	// Field和FieldEx
	if args.Fields != "" {
		m = m.Fields(args.Fields)
	}

	if args.FieldsEx != "" {
		m = m.Fields(args.FieldsEx)
	}

	switch args.Type {
	case "all":
		// 获取全部
		// page 条件
		if args.Page > 0 {
			m = m.Offset((args.Page - 1) * args.DataPerPage).Limit(args.DataPerPage)
		}

		// order 条件
		if args.Order != "" {
			m = m.Order(args.Order)
		}

		// limit 条件
		if args.Limit > 0 {
			m = m.Limit(args.Limit)
		}

		res, err := m.All()

		if err != nil {
			return nil, err
		}

		return res.List(), nil

	case "one":
		// 获取一个
		res, err := m.One()
		if err != nil{
			return nil, err
		}

		if res == nil{
			return nil, NotFound
		}

		err = res.Struct(model)
		if err != nil{
			return nil, err
		}

		return model, nil

	case "count":
		res, err := m.Count()
		if err != nil{
			return nil, err
		}

		return res, nil
	}

	return nil, GetTypeError
}

// 增加数据
func (d *DAO) Create(args g.Map, model Model) (interface{}, error) {
	m := g.DB().Table(model.TableName())

	res, err := m.Data(args).Insert()

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil{
		return nil, err
	}

	return id, nil
}

// 增加数据（事务）
func (d *DAO) CreateWithTX(args g.Map, model Model, tx *gdb.TX) (interface{}, error) {
	m := tx.Table(model.TableName()).Safe(true)

	res, err := m.Data(args).Insert()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	newId, err := res.LastInsertId()
	if err != nil{
		return nil, err
	}

	return int(newId), nil
}

// 修改数据
func (d *DAO) Update(args g.Map, model Model) error {
	m := g.DB().Table(model.TableName())

	_, err := m.Data(args).Update()
	if err != nil {
		return err
	}
	return nil
}

// 修改数据（事务）
func (d *DAO) UpdateWithTX(args g.Map, model Model, tx *gdb.TX) error {
	m := tx.Table(model.TableName()).Safe(true)

	_, err := m.Data(args).Update()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}

		return err
	}
	return nil
}

// 根据id删除数据
func (d *DAO) DeleteByID(id int, model Model) error {
	m := g.DB().Table(model.TableName())

	_, err := m.WherePri(id).Delete()
	if err != nil {
		return err
	}

	return nil
}

// 根据id删除数据(事务)
func (d *DAO) DeleteByIDWithTX(id int, model Model, tx *gdb.TX) error {
	m := tx.Table(model.TableName()).Safe(true)

	_, err := m.WherePri(id).Delete()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}

		return err
	}

	return nil
}