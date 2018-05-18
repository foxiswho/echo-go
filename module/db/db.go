package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/xormplus/xorm"
	. "github.com/foxiswho/echo-go/conf"
	"github.com/foxiswho/echo-go/module/log"
	"reflect"
	"strconv"
)

//结构体
type Db struct {
	Engine        *xorm.Engine
	FilterSession *xorm.Session
}

//全局变量
var db *Db

//数据库xiangg dsn 格式化
func dsn() string {
	db_user := Conf.DB.UserName
	db_pass := Conf.DB.Pwd
	db_host := Conf.DB.Host
	db_port := Conf.DB.Port
	db_name := Conf.DB.Name
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	return dsn
}

//创建数据库连接
func newDB() (*Db, error) {
	var err error
	db = new(Db)
	db.Engine, err = xorm.NewEngine("mysql", dsn())
	if err != nil {
		fmt.Println("NewEngine", err)
		panic(err.Error())
	}
	db.Engine.ShowSQL(true)
	locat, _ := time.LoadLocation("Asia/Shanghai")
	db.Engine.TZLocation = locat
	db.Engine.DatabaseTZ = locat
	//defer db.Engine.Close()
	return db, nil
}

//快捷调研
func DB() *Db {
	if db == nil {
		log.Debugf("Model NewDB")
		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.Engine.SetMaxIdleConns(10)
		newDb.Engine.SetMaxIdleConns(100)
		//newDb.Engine.SetLogger(orm.Logger{})
		db = newDb
	}
	return db
}

//sql map XML文件
func SqlMapFile(directory, extension string) (*xorm.XmlSqlMap) {
	return xorm.Xml(directory, extension)
}

//初始化
//func Init() {
//	DB()
//}

type QuerySession struct {
	Session *xorm.Session
}

var Query *QuerySession

func Filter(where []*QueryCondition) *xorm.Session {
	db := DB()
	Query = new(QuerySession)
	if len(where) > 0 {
		i := 1
		for _, qc := range where {
			condition := qc.Condition
			key := qc.Field + qc.Operation + "?"
			fmt.Println("QuerySession", qc)
			fmt.Println("where query key=>", key)
			//fmt.Println(k, condition, reflect.TypeOf(condition))
			//fmt.Println("?号个数为", strings.Count(k, "?"))
			arrCount := 0
			var arr []string
			switch condition.(type) {
			case string:
				//是字符时做的事情
				if condition == "" {
					FilterWhereAnd(db, i, key, "")
				} else {
					FilterWhereAnd(db, i, key, condition.(string))
				}
			case int, int8, int16, int32, int64:
				//是整数时做的事情
				FilterWhereAnd(db, i, key, condition)
			case float32, float64:
				//是整数时做的事情
				FilterWhereAnd(db, i, key, condition)
			case []string:
				arr = condition.([]string)
				arrCount = len(arr)
				if arrCount == 0 {
					FilterWhereAnd(db, i, key, "")
				} else {
					str := ""
					for j, val := range arr {
						if j > 0 {
							str += ","
						}
						str += val
					}
					FilterWhereAnd(db, i, qc.Field+qc.Operation+"(?)", str)
				}

			case []int, []int8, []int16, []int32, []int64:
				objArrToIntArr(db, i, qc)
			case []float32, []float64:
				objArrToFloatArr(db, i, qc)
			default:
				fmt.Println("其他还没有收录")
				fmt.Println("其他还没有收录")
				fmt.Println("其他还没有收录")
				fmt.Println("其他还没有收录")
				fmt.Println("其他还没有收录")
				fmt.Println("其他还没有收录")
				fmt.Println("其他还没有收录")
			}
			i++
		}
	} else {
		//初始化
		Query.Session = db.Engine.Limit(20, 0)
	}

	return Query.Session
}
func FilterWhereAnd(db *Db, i int, key string, value ...interface{}) {
	fmt.Println("key", key)
	fmt.Println("value", value)
	fmt.Println("TypeOf", reflect.TypeOf(value))
	if i == 1 {
		Query.Session = DB().Engine.Where(key, value...)
	} else {
		Query.Session = Query.Session.And(key, value...)
	}
}

func objArrToIntArr(db *Db, i int, qc *QueryCondition) {
	str := ""
	arrCount := 0
	switch qc.Condition.(type) {
	case []int64:
		arr := qc.Condition.([]int64)
		arrCount = len(arr)
		for j, val := range arr {
			if j > 0 {
				str += ","
			}
			str += strconv.FormatInt(val, 10)
		}
	case []int32:
		arr := qc.Condition.([]int32)
		arrCount = len(arr)
		for j, val := range arr {
			if j > 0 {
				str += ","
			}
			str += strconv.FormatInt(int64(val), 10)
		}
	case []int16:
		arr := qc.Condition.([]int16)
		arrCount = len(arr)
		for j, val := range arr {
			if j > 0 {
				str += ","
			}
			str += strconv.FormatInt(int64(val), 10)
		}
	case []int8:
		arr := qc.Condition.([]int8)
		arrCount = len(arr)
		for j, val := range arr {
			if j > 0 {
				str += ","
			}
			str += strconv.FormatInt(int64(val), 10)
		}
	case []int:
		arr := qc.Condition.([]int)
		arrCount = len(arr)
		for j, val := range arr {
			if j > 0 {
				str += ","
			}
			str += strconv.FormatInt(int64(val), 10)
		}
	}
	if arrCount == 0 {
		FilterWhereAnd(db, i, qc.Field+qc.Operation, "")
	} else {
		FilterWhereAnd(db, i, qc.Field+qc.Operation+"(?)", str)
	}
}

func objArrToFloatArr(db *Db, i int, qc *QueryCondition) {
	str := ""
	arrCount := 0
	switch qc.Condition.(type) {
	case []float64:
		arr := qc.Condition.([]float64)
		arrCount = len(arr)
		for j, val := range arr {
			if j > 0 {
				str += ","
			}
			str += strconv.FormatFloat(val, 'E', -1, 64)
		}
	case []float32:
		arr := qc.Condition.([]float32)
		arrCount = len(arr)
		for j, val := range arr {
			if j > 0 {
				str += ","
			}
			str += strconv.FormatFloat(float64(val), 'E', -1, 64)
		}

	}
	if arrCount == 0 {
		FilterWhereAnd(db, i, qc.Field+qc.Operation, "")
	} else {
		FilterWhereAnd(db, i, qc.Field+qc.Operation+"(?)", str)
	}
}
