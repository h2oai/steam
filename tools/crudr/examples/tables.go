//go:generate crudr $GOFILE

package examples

import "database/sql"

type Pk struct {
	Num   int            `db:"num,arg"`
	Name  string         `db:"name,arg",json:"omitempty"`
	Type1 string         `db:"type_1,def="normal""`
	Type2 sql.NullString `db:"type_2"`
}
