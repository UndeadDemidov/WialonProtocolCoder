package Tests

import (
	"encoding/hex"
	"testing"

	"github.com/MrAmperage/WialonProtocolCoder/Decode"
)

func TestDecodePackage(t *testing.T) {
	Message := "740000003230353030313731310062a048a6000000030bbb000000270102706f73696e666f000000807cb2d8c941000080ca1431c041000000000036a640007a015713"
	ByteMessage, _ := hex.DecodeString(Message)

	DecodePackage := Decode.Decode(ByteMessage)

	if DecodePackage.PackageSize != 116 {

		t.Error("Размер пакета декодировался некорректно")
		t.Log("Размер пакета", DecodePackage.PackageSize)
	}

	if DecodePackage.UUID != "205001711" {
		t.Error("UID декодирован некорректно")

	}
	if DecodePackage.UnixTime != 1654671526 {
		t.Error("Дата декодирована неверно")
		t.Log("UnixTime", DecodePackage.UnixTime)
	}
	if DecodePackage.ByteMask != 3 {

		t.Error("Битовая маска декодирована неверно")
		t.Log("Битовая маска", DecodePackage.ByteMask)
	}
	if DecodePackage.DataBlocks[0].BlockType != 3003 {

		t.Error("Тип блока декодирован неверно")
		t.Log("Битовая маска", DecodePackage.DataBlocks[0].BlockType)
	}
	if DecodePackage.DataBlocks[0].BlockSize != 39 {

		t.Error("Размер блока задан неверно")
		t.Log("Размер блока", DecodePackage.DataBlocks[0].BlockSize)
	}
	if DecodePackage.DataBlocks[0].Hidden != 1 {

		t.Error("Атрибут скрытности задан неверно")
		t.Log("Атрибут скрытности ", DecodePackage.DataBlocks[0].Hidden)

	}
	if DecodePackage.DataBlocks[0].BlockDataType != 2 {

		t.Error("Тип данных блока задан неверно")
		t.Log("Тип данных блока ", DecodePackage.DataBlocks[0].BlockDataType)
	}

	if DecodePackage.DataBlocks[0].BlockName != "posinfo" {

		t.Error("Имя блока  задано неверно")
		t.Log("Имя блока", DecodePackage.DataBlocks[0].BlockName)
	}

	if DecodePackage.DataBlocks[0].Longitude != 867263737 {

		t.Error("Долгота  задана неверно")
		t.Log("Долгота", DecodePackage.DataBlocks[0].Longitude)
	}

	if DecodePackage.DataBlocks[0].Latitude != 543304085 {
		t.Error("Широта  задана неверно")
		t.Log("Широта", DecodePackage.DataBlocks[0].Longitude)
	}

	if DecodePackage.DataBlocks[0].Height != 2843 {
		t.Error("Высота над уровнем моря неверна")
		t.Log("Высота над уровнем моря", DecodePackage.DataBlocks[0].Height)
	}
	if DecodePackage.DataBlocks[0].Speed != 122 {
		t.Error("Скорость неверна")
		t.Log("Скорость", DecodePackage.DataBlocks[0].Speed)
	}
	if DecodePackage.DataBlocks[0].Course != 343 {
		t.Error("Курс неверен")
		t.Log("Курс", DecodePackage.DataBlocks[0].Course)

	}
	if DecodePackage.DataBlocks[0].SatellitesCount != 19 {
		t.Error("Количество спутников неверно")
		t.Log("Количество спутников", DecodePackage.DataBlocks[0].SatellitesCount)

	}

}
