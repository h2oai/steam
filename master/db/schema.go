package db

//go:generate icing -idl schema.ice -go encoding.go

type Printer func(b []byte) (string, error)

var Printers = make(map[string]Printer)

type Record struct {
	ID         string
	CreatedAt  int64
	ModifiedAt int64
}
