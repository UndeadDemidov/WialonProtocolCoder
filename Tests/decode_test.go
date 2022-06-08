package Tests

import (
	"encoding/hex"
	"testing"

	"github.com/MrAmperage/WialonProtocolCoder/Decode"
)

func TestDecodePackage(t *testing.T) {
	Message := "74000000333533393736303133343435343835004B0BFB70000000030BBB000000270102706F73696E666F00A027AFDF5D9848403AC7253383DD4B400000000000805A40003601460B0BBB0000001200047077725F657874002B87ё16D9CE973B400BBB00000011010361766C5F696E707574730000000001"
	ByteMessage, _ := hex.DecodeString(Message)

	DecodePackage := Decode.Decode(ByteMessage)

	if DecodePackage.PackageSize != 116 {

		t.Error("Размер пакета декодировался некорректно")
	}

	if DecodePackage.UUID != "353976013445485" {
		t.Error("UID декодирован некорректно")

	}
	if DecodePackage.ByteMask != 3 {
		t.Error("Битовая маска декодирована неверно")
	}

	if DecodePackage.DataBlocks[0].Longitude != 49.1903648 {

		t.Error("Неверно декодирована долгота")
	}

	if DecodePackage.DataBlocks[0].Latitude != 55.7305664 {

		t.Error("Неверно декодирована широта")
	}

	if DecodePackage.DataBlocks[0].Height != 106.0 {
		t.Error("Неверно декодирована высота")
	}

}
