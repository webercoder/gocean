package coops

// ValueBasedResult has a time and value.
type ValueBasedResult struct {
	Time  string `xml:"t,attr" json:"t"`
	Value string `xml:"v,attr" json:"v"`
}

// ValueBasedResultWithFlags is a ValueBasedResult but with a flags field.
type ValueBasedResultWithFlags struct {
	ValueBasedResult
	Flags string `xml:"f,attr" json:"f"`
}
