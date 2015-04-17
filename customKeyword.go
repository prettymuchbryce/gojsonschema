package gojsonschema

type CustomKeyword interface {
	GetKeyword() string
	Validate(keywordValue interface{}, documentValue interface{}) error
}

type customKeywordValue struct {
	value         interface{}
	customKeyword CustomKeyword
}
