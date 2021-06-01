package coops

import "fmt"

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

// Pluralize returns a singular or plural version of a string.
func Pluralize(number, singularFormat, pluralFormat string) string {
	if number == "1" {
		return fmt.Sprintf(singularFormat, number)
	}

	return fmt.Sprintf(pluralFormat, number)
}
