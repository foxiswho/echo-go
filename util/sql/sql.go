package sql

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"
	"unicode"
)

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

func SqlParse(sql string, vars []interface{}) string {
	var (
		sqlStr          string
		formattedValues []string
	)
	for _, value := range vars {
		indirectValue := reflect.Indirect(reflect.ValueOf(value))
		if indirectValue.IsValid() {
			value = indirectValue.Interface()
			if t, ok := value.(time.Time); ok {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
			} else if b, ok := value.([]byte); ok {
				if str := string(b); isPrintable(str) {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
				} else {
					formattedValues = append(formattedValues, "'<binary>'")
				}
			} else if r, ok := value.(driver.Valuer); ok {
				if value, err := r.Value(); err == nil && value != nil {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			} else {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
			}
		} else {
			formattedValues = append(formattedValues, "NULL")
		}
	}

	// differentiate between $n placeholders or else treat like ?
	if numericPlaceHolderRegexp.MatchString(sql) {
		sqlStr = sql
		for index, value := range formattedValues {
			placeholder := fmt.Sprintf(`\$%d`, index+1)
			sqlStr = regexp.MustCompile(placeholder).ReplaceAllString(sqlStr, value)
		}
	} else {
		formattedValuesLength := len(formattedValues)
		for index, value := range sqlRegexp.Split(sql, -1) {
			sqlStr += value
			if index < formattedValuesLength {
				sqlStr += formattedValues[index]
			}
		}

	}

	return sqlStr
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
