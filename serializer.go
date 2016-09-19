package logify

type Serializer interface {
	Serialize(*Entry, []Field) error
}
