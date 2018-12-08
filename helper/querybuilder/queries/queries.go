package qbqueries

import (
	"fmt"
	"reflect"

	"github.com/wa-gtwy/helper/querybuilder/model"
)

// this function below, to decide is the last loop or not
func isLastIteration(currNumb int, endNumb int) bool {
	if (currNumb + 1) == endNumb {
		return true
	}
	return false
}

// check reflect kind is string or not
func isString(dType reflect.Kind) bool {
	if dType == reflect.String {
		return true
	}
	return false
}

// check reflect kind is int or not
func isInt(dType reflect.Kind) bool {
	if dType == reflect.Int || dType == reflect.Int8 || dType == reflect.Int16 || dType == reflect.Int32 || dType == reflect.Int64 {
		return true
	}
	return false
}

// check reflect kind is float or not
func isFloat(dType reflect.Kind) bool {
	if dType == reflect.Float32 || dType == reflect.Float64 {
		return true
	}
	return false
}

// check reflect kind is bool or not
func isBool(dType reflect.Kind) bool {
	if dType == reflect.Bool {
		return true
	}
	return false
}

// check reflect kind is struct or not
func isStruct(dType reflect.Kind) bool {
	if dType == reflect.Struct {
		return true
	}
	return false
}

// CreateQueriesInsert ...
func CreateQueriesInsert(data interface{}) (string, string) {
	var cols string
	var injection string

	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		var separComma string
		if isLastIteration(i, v.NumField()) == true {
			separComma = ""
		} else {
			separComma = ","
		}

		// set column
		columnName := v.Type().Field(i).Tag.Get("db")
		cols = fmt.Sprintf("%s\"%s\"%s ", cols, columnName, separComma)

		// set injection
		injection = fmt.Sprintf("%s$%d%s ", injection, i+1, separComma)
	}
	return cols, injection
}

// CreateQueriesUpdate ...
func CreateQueriesUpdate(data interface{}) string {
	var sets string

	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		separComma := ","
		if isLastIteration(i, v.NumField()) == true {
			separComma = ""
		}
		// get column
		column := v.Type().Field(i).Tag.Get("db")

		// set
		sets = fmt.Sprintf("%s%s=$%d%s", sets, column, i+1, separComma)

	}
	return sets
}

// CreateQueriesWhere ...
func CreateQueriesWhere(conditions []*qbmodel.Condition) string {
	var where string
	for index, condition := range conditions {
		nextCondition := condition.NextCond
		if isLastIteration(index, len(conditions)) {
			nextCondition = ""
		}

		if isString(reflect.TypeOf(condition.Value).Kind()) == true || isStruct(reflect.TypeOf(condition.Value).Kind()) {
			where = fmt.Sprintf("%s \"%s\"%s'%v' %s", where, condition.Key, condition.Operator, condition.Value, nextCondition)
		} else {
			where = fmt.Sprintf("%s \"%s\"%s%v %s", where, condition.Key, condition.Operator, condition.Value, nextCondition)
		}
	}
	return where
}

// CreateQueriesWith ...
func CreateQueriesWith(joins []*qbmodel.Join) string {
	var with string
	for _, join := range joins {
		with = fmt.Sprintf(" %s %v %s %s ON %s=%s", with, join.TableFrom, join.Type, join.TableWith, join.ColumnTableFrom, join.ColumnTableWith)
	}
	return with
}
