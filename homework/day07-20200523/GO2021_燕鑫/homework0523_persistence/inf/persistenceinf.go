package inf

type PersisInf interface {
	Encode(string) error
	Decode(string) error
	GetObj() interface{}
}
