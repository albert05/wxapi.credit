package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"unsafe"
	"reflect"
)

type Mysql struct {
	DB *sql.DB
}

type queryResult struct {
	affectedRows, insertId int64
}

const DriverNAME = "mysql"
const TableMethodName = "GetTableName"

var mysqlDB map[string]Mysql
var Conn Mysql
var DSN string

func Init(dsn string) {
	DSN = dsn
	mysqlDB = make(map[string]Mysql)
	Conn = GetInstance()
}

func GetInstance() Mysql {
	if mysql, ok := mysqlDB[DSN]; ok {
		return mysql
	}

	db, err := sql.Open(DriverNAME, DSN)
	if err != nil {
		panic(err)
		return Mysql{}
	}

	mysqlDB[DSN] = Mysql{DB: db}
	return Mysql{DB: db}
}

func GetTable(v interface{}) string {
	rv := reflect.ValueOf(v)
	return rv.MethodByName(TableMethodName).Call(nil)[0].String()
}

//单条记录
func FindCond(v interface{}, where map[string]string, fields string) error {
	s := GetSqlInstance()
	s.SetValues(TYPE_FIND, GetTable(v), fields, nil, where, nil)
	s.Create()

	return FindOne(v, s.S)
}

//单条记录
func FindMultiCond(v interface{}, multiWhere [][]string, fields string) error {
	s := GetSqlInstance()
	s.SetValues(TYPE_FIND, GetTable(v), fields, nil, nil, multiWhere)
	s.Create()

	return FindOne(v, s.S)
}

//单条记录
func FindOne(vr interface{}, sql string) error {
	rows, err := Conn.DB.Query(sql)
	if err != nil {
		return err
	}
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	count := len(columns)
	//定义输出的类型
	result := make(MapModel)
	//这个是sql查询出来的字段
	values := make([]interface{}, count)
	//保存sql查询出来的对应的地址
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		//scansql查询出来的字段的地址
		rows.Scan(valuePtrs...)

		//开始循环columns
		for i, col := range columns {
			var v interface{}
			//值
			val := values[i]
			//判读值的类型（interface类型）如果是byte，则需要转换成字符串
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			//保存
			result[col] = v
		}
	}

	result.Load(vr)
	return nil
}

//多条记录（根据上面的多条记录修改）
func FindAll(sql string) ([]MapModel, error) {
	rows, err := Conn.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make([]MapModel, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(MapModel)
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	if err != nil {
		return nil, err
	}

	return tableData, nil
}

// 插入数据
// string table
// map data 插入的数据
// return bool
func Insert(vr interface{}) bool {
	s := GetSqlInstance()
	s.SetValues(TYPE_INSERT, GetTable(vr), "", Unload(vr, true), nil, nil)
	s.Create()

	stmtIns, err := Conn.DB.Prepare(s.S) // ? = placeholder
	if err != nil {
		panic(err)
	}

	defer stmtIns.Close()

	result, err := stmtIns.Exec(s.PreValues...) // Insert tuples (i, i^2)
	if err != nil {
		panic(err)
	}

	var qResult queryResult = *(*queryResult)(unsafe.Pointer(&result))
	return qResult.affectedRows > 0
}

// 更新数据
// string table
// map data 插入的数据
// map condition 更新条件
// return bool
func Update(vr interface{}) bool {
	data := Unload(vr, false)
	id := data["id"]
	delete(data, "id")
	s := GetSqlInstance()
	s.SetValues(TYPE_UPDATE, GetTable(vr), "", data, map[string]string{"id": id}, nil)
	s.Create()

	stmtIns, err := Conn.DB.Prepare(s.S) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close()

	result, err := stmtIns.Exec(s.PreValues...) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var qResult queryResult = *(*queryResult)(unsafe.Pointer(&result))
	return qResult.affectedRows > 0
}

// 更新数据
// string table
// map data 插入的数据
// map condition 更新条件
// return bool
func UpdateCond(vr interface{}, data map[string]string, condition map[string]string) bool {
	s := GetSqlInstance()
	s.SetValues(TYPE_UPDATE, GetTable(vr), "", data, condition, nil)
	s.Create()

	stmtIns, err := Conn.DB.Prepare(s.S) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close()

	result, err := stmtIns.Exec(s.PreValues...) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var qResult queryResult = *(*queryResult)(unsafe.Pointer(&result))
	return qResult.affectedRows > 0
}

func Exec(sql string) error {
	_, err := Conn.DB.Exec(sql)
	return err
}
