package structs

type ChloeWriter interface {
	WriteObject(object *ResponseObject) error
	Write([]byte) (int, error)
	Close() error
	Flush()
}
