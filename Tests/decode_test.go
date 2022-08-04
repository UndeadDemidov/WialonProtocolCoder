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
		t.Log("Размер пакета", DecodePackage.PackageSize)
	}

	if DecodePackage.UUID != "353976013445485" {
		t.Error("UID декодирован некорректно")
		t.Log("UID", DecodePackage.UUID)

	}
	if DecodePackage.UnixTime != 1565613499 {
		t.Error("Дата декодирована неверно")
		t.Log("UnixTime", DecodePackage.UnixTime)
	}
	if DecodePackage.ByteMask != 3 {

		t.Error("Битовая маска декодирована неверно")
		t.Log("Битовая маска", DecodePackage.ByteMask)
	}
	for _, DataBlock := range DecodePackage.DataBlocks {

		switch DataBlock.(type) {
		case Decode.PositionInfoBlock:
			if DataBlock.(Decode.PositionInfoBlock).BlockType != 3003 {
				t.Error("Тип блока декодирован неверно")
				t.Log("Тип блока", DataBlock.(Decode.PositionInfoBlock).BlockType)
			}
			if DataBlock.(Decode.PositionInfoBlock).BlockSize != 39 {
				t.Error("Размер блока задан неверно")
				t.Log("Размер блока", DataBlock.(Decode.PositionInfoBlock).BlockSize)
			}
			if DataBlock.(Decode.PositionInfoBlock).BlockSize != 1 {
				t.Error("Атрибут скрытности задан неверно")
				t.Log("Атрибут скрытности ", DataBlock.(Decode.PositionInfoBlock).Hidden)
			}
			if DataBlock.(Decode.PositionInfoBlock).BlockDataType != 2 {

				t.Error("Тип данных блока задан неверно")
				t.Log("Тип данных блока ", DataBlock.(Decode.PositionInfoBlock).BlockDataType)

			}
			if DataBlock.(Decode.PositionInfoBlock).BlockName != "posinfo" {
				t.Error("Имя блока  задано неверно")
				t.Log("Имя блока", DataBlock.(Decode.PositionInfoBlock).BlockName)
			}
			if DataBlock.(Decode.PositionInfoBlock).Longitude != 867263737 {
				t.Error("Долгота  задана неверно")
				t.Log("Долгота", DataBlock.(Decode.PositionInfoBlock).Longitude)
			}
			if DataBlock.(Decode.PositionInfoBlock).Latitude != 543304085 {
				t.Error("Широта  задана неверно")
				t.Log("Широта", DataBlock.(Decode.PositionInfoBlock).Latitude)
			}
			if DataBlock.(Decode.PositionInfoBlock).Height != 2843 {
				t.Error("Высота над уровнем моря неверна")
				t.Log("Высота над уровнем моря", DataBlock.(Decode.PositionInfoBlock).Height)
			}
			if DataBlock.(Decode.PositionInfoBlock).Speed != 122 {
				t.Error("Скорость неверна")
				t.Log("Скорость", DataBlock.(Decode.PositionInfoBlock).Speed)
			}
			if DataBlock.(Decode.PositionInfoBlock).Course != 343 {
				t.Error("Курс неверен")
				t.Log("Курс", DataBlock.(Decode.PositionInfoBlock).Course)
			}
			if DataBlock.(Decode.PositionInfoBlock).SatellitesCount != 19 {
				t.Error("Количество спутников неверно")
				t.Log("Количество спутников", DataBlock.(Decode.PositionInfoBlock).SatellitesCount)
			}

		case Decode.SensorsInfoBlock:
			if DataBlock.(Decode.SensorsInfoBlock).BlockName != "pwr_ext" {
				t.Error("Имя блока  задано неверно")
				t.Log("Имя блока", DataBlock.(Decode.SensorsInfoBlock).BlockName)
			}
		}

	}

}
