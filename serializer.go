package logify

type Serializer interface {
	Serialize(*Entry) error
}
