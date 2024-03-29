package Encode

import "database/sql"

type EncodePackage struct {
	UID    string
	TS     uint32
	Flags  sql.NullInt64
	Lon    sql.NullInt64
	Lat    sql.NullInt64
	Alt    sql.NullInt64
	Speed  sql.NullInt64
	Dir    sql.NullInt64
	Vlt    sql.NullInt64
	Sat    sql.NullInt64
	U      [6]sql.NullInt64
	LLS1   sql.NullInt64
	LLS2   sql.NullInt64
	SPNId  []byte
	SPNVal []byte
	IBtn   []byte
	GPSAlt sql.NullInt64
	Unval0 float64
	Unval1 float64
	CLLS1  sql.NullInt64
}
