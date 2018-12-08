package qbmodel

// Condition ...
type Condition struct {
	Key      string
	Operator string
	Value    interface{}
	NextCond string
}
