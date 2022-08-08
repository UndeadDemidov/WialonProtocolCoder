package Tests

import (
	"database/sql"
	"testing"

	"github.com/MrAmperage/WialonProtocolCoder/Encode"
)

func TestEncodePackage(t *testing.T) {
	NewPackage := Encode.EncodePackage{UID: "423876197", TS: 423903526, Lat: sql.NullInt64{Int64: 543304085, Valid: true}, Lon: sql.NullInt64{Int64: 867263737, Valid: true}, Sat: sql.NullInt64{Valid: true, Int64: 11}, GPSAlt: sql.NullInt64{Valid: true, Int64: 11}, Unval0: sql.NullInt64{Valid: true, Int64: 550}}
	t.Log(Encode.Encode(NewPackage, 1230768000))
}
