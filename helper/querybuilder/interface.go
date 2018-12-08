package qb

import "github.com/wa-gtwy/helper/querybuilder/model"

// QueryBuilderInterface ...
type QueryBuilderInterface interface {
	Query(string, int, int) string
	QueryWhere(string, []*qbmodel.Condition) string
	QueryWith(string, []*qbmodel.Join) string
	QueryWhereWith(string, []*qbmodel.Join, []*qbmodel.Condition) string
	Create(string, interface{}) string
	Update(string, interface{}, []*qbmodel.Condition) string
	UpdateWhere(string, interface{}, []*qbmodel.Condition) string
	Delete(string, []*qbmodel.Condition) string
}
