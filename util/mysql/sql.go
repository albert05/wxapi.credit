package mysql

import (
	"fmt"
	"strings"
	"time"
	"strconv"
)

const TYPE_FIND 	= 0
const TYPE_INSERT 	= 1
const TYPE_DELTE 	= 2
const TYPE_UPDATE 	= 3

const DEFAULT_FIELDS  = "*"

const CREATED_AT_FIELD = "created_at"
const UPDATED_AT_FIELD = "updated_at"

var TypeMap = map[int]string {
	TYPE_FIND: "select",
	TYPE_INSERT: "insert",
	TYPE_DELTE: "delete",
	TYPE_UPDATE: "update",
}

type Sql struct {
	OptType 		int
	TableName 		string
	Fields			string
	Values			map[string]string
	Conditions		map[string]string	`key => value`
	MultiConditions	[][]string			`opt, key, val`
	PreValues 		[]interface{}
	S 				string
}

func GetSqlInstance() *Sql {
	return &Sql{
		OptType: TYPE_FIND,
		Fields: DEFAULT_FIELDS,
	}
}

func (s *Sql) SetValues(opt int, table string, fields string, values map[string]string, cond map[string]string, multiCond [][]string) {
	if _, ok := TypeMap[opt]; ok {
		s.OptType = opt
	}

	if table != "" {
		s.TableName = table
	}

	if fields != "" {
		s.Fields = fields
	}

	if values != nil {
		s.Values = values
	}

	if cond != nil {
		s.Conditions = cond
	}

	if multiCond != nil {
		s.MultiConditions = multiCond
	}
}

func (s *Sql) Create() {
	s.AutoTime()
	switch s.OptType {
	case TYPE_FIND:
		s.Find()
	case TYPE_INSERT:
		s.Insert()
	case TYPE_DELTE:
		s.Delete()
	case TYPE_UPDATE:
		s.Update()
	}
}

func (s *Sql) AutoTime() {
	now := strconv.FormatInt(time.Now().Unix(), 10)

	switch s.OptType {
	case TYPE_INSERT:
		s.Values[CREATED_AT_FIELD] = now
		s.Values[UPDATED_AT_FIELD] = now
	case TYPE_UPDATE:
		s.Values[UPDATED_AT_FIELD] = now
	}
}


func(s *Sql) Find() {
	sql := fmt.Sprintf("%s %s from %s where 1=1 ", TypeMap[s.OptType], s.Fields, s.TableName)

	// key value
	if len(s.Conditions) > 0 {
		for  k, v := range s.Conditions {
			sql += fmt.Sprintf(" and %s='%s' ", k, v)
		}
	}

	// opt, key, value
	if len(s.MultiConditions) > 0 {
		for  _, item := range s.MultiConditions {
			if len(item) == 3 {
				sql += fmt.Sprintf(" and %s %s '%s' ", item[1], item[0], item[2])
			}
		}
	}

	s.S = sql
}

func(s *Sql) Insert() {
	l := len(s.Values)
	if l > 0 {
		s.PreValues = make([]interface{}, l)
		i := 0
		km := make([]string, 0)
		vm := make([]string, 0)
		for k, v := range s.Values {
			km = append(km, k)
			vm = append(vm, "?") // value pre set ?
			s.PreValues[i] = v
			i++
		}

		fields := strings.Join(km, ",")
		values := strings.Join(vm, ",")

		s.S = fmt.Sprintf("%s into %s (%s) values (%s) ", TypeMap[s.OptType], s.TableName, fields, values)
	}
}

func(s *Sql) Delete() {
	sql := fmt.Sprintf("%s from %s where 1=1 ", TypeMap[s.OptType], s.TableName)

	s.PreValues = make([]interface{}, len(s.Conditions) + len(s.MultiConditions))
	i := 0

	// key value
	if len(s.Conditions) > 0 {
		for  k, v := range s.Conditions {
			sql += fmt.Sprintf(" and %s=? ", k, v)
			s.PreValues[i] = v
			i++
		}
	}

	// opt, key, value
	if len(s.MultiConditions) > 0 {
		for  _, item := range s.MultiConditions {
			if len(item) == 3 {
				sql += fmt.Sprintf(" and %s %s ? ", item[1], item[0])
				s.PreValues[i] = item[2]
				i++
			}
		}
	}

	s.S = sql
}

func(s *Sql) Update() {
	l := len(s.Values)
	if l < 0 {
		return
	}

	sl := len(s.Conditions) + len(s.MultiConditions) + len(s.Values)
	s.PreValues = make([]interface{}, sl)
	i := 0

	vs := make([]string, 0)
	for k, v := range s.Values {
		vs = append(vs, k + "=?")
		s.PreValues[i] = v
		i++
	}

	var where string
	// key value
	if len(s.Conditions) > 0 {
		for  k, v := range s.Conditions {
			where += fmt.Sprintf(" and %s=? ", k)
			s.PreValues[i] = v
			i++
		}
	}

	// opt, key, value
	if len(s.MultiConditions) > 0 {
		for  _, item := range s.MultiConditions {
			if len(item) == 3 {
				where += fmt.Sprintf(" and %s %s ? ", item[1], item[0])
				s.PreValues[i] = item[2]
				i++
			}
		}
	}

	values := strings.Join(vs, ",")
	s.S = fmt.Sprintf("%s %s set %s where 1=1 %s", TypeMap[s.OptType], s.TableName, values, where)
}
