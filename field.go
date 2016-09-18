package logify

type fieldType uint8

const (
	boolType fieldType = iota + 1
	intType
	stringType
)

type Field struct {
	key       string
	value     interface{}
	fieldType fieldType
}

func Bool(key string, value bool) Field {
	return Field{key: key, value: value, fieldType: boolType}
}

func Int(key string, value int) Field {
	return Field{key: key, value: value, fieldType: intType}
}

func String(key, value string) Field {
	return Field{key: key, value: value, fieldType: stringType}
}
