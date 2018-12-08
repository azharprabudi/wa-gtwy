package qbmodel

// Join ...
type Join struct {
	Type            string
	TableFrom       string
	ColumnTableFrom string
	TableWith       string
	ColumnTableWith string
}
