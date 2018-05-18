package db

//查询条件
type QueryCondition struct {
	Field     string      //字段
	Operation string      //操作 =,!=,>,<,>=,<=,in,not in
	Condition interface{} //具体值
}

var expression []*QueryCondition

func NewMakeQueryCondition() []*QueryCondition {
	return make([]*QueryCondition, 0)
}

//添加条件
func AddQueryCondition(field, operation string, condition interface{}) *QueryCondition {
	return &QueryCondition{field, operation, condition}
}
