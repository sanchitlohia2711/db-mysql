package filter

//Query struct
type Query struct {
	filters []*Filter
}

//Filter filter struct
type Filter struct {
	Field    string
	Value    interface{}
	Operator string
}

//Update struct
type Update struct {
	Field string
	Value interface{}
}

//Sort struct
type Sort struct {
	field string
	order string
}
