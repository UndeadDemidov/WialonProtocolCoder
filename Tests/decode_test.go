package Tests

import (
	"testing"

	"github.com/MrAmperage/WialonProtocolCoder/Decode"
)

func TestDecodePackage(t *testing.T) {

	DecodePackage := Decode.Decode([]byte{15, 1, 0, 0, 56, 54, 49, 55, 55, 52, 48, 53, 56, 48, 51, 52, 56, 57, 49, 0, 98, 242, 54, 7, 0, 0, 0, 3, 11, 187, 0, 0, 0, 39, 1, 2, 112, 111, 115, 105, 110, 102, 111, 0, 84, 141, 94, 13, 80, 170, 85, 64, 1, 106, 106, 217, 90, 255, 74, 64, 0, 0, 0, 0, 0, 32, 101, 64, 0, 0, 0, 180, 11, 11, 187, 0, 0, 0, 28, 0, 4, 83, 65, 84, 69, 76, 76, 73, 84, 69, 83, 95, 66, 121, 116, 101, 95, 48, 0, 0, 0, 0, 0, 0, 0, 40, 64, 11, 187, 0, 0, 0, 23, 0, 4, 79, 84, 72, 69, 82, 95, 66, 121, 116, 101, 95, 49, 0, 0, 0, 0, 0, 0, 0, 28, 64, 11, 187, 0, 0, 0, 26, 0, 4, 86, 79, 76, 84, 65, 71, 69, 95, 83, 104, 111, 114, 116, 95, 53, 0, 0, 0, 0, 0, 128, 180, 218, 64, 11, 187, 0, 0, 0, 28, 0, 4, 76, 76, 83, 95, 70, 85, 69, 76, 95, 83, 104, 111, 114, 116, 95, 49, 51, 0, 0, 0, 0, 0, 0, 100, 185, 64, 11, 187, 0, 0, 0, 32, 0, 4, 76, 76, 83, 95, 70, 85, 69, 76, 95, 83, 104, 111, 114, 116, 95, 49, 51, 95, 116, 97, 114, 0, 0, 0, 0, 0, 0, 192, 152, 64, 11, 187, 0, 0, 0, 29, 0, 4, 66, 85, 67, 75, 69, 84, 65, 78, 71, 95, 83, 104, 111, 114, 116, 95, 49, 52, 0, 0, 0, 0, 0, 0, 64, 89, 64})
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
			if DataBlock.(Decode.PositionInfoBlock).BlockSize != 391 {
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

		case Decode.AdditionalValueBlock[uint32]:
			if DataBlock.(Decode.AdditionalValueBlock[uint32]).BlockName != "CLLS" {
				t.Error("Имя блока  задано неверно")
				t.Log("Имя блока", DataBlock.(Decode.AdditionalValueBlock[uint32]).BlockName)
			}

		case Decode.AdditionalValueBlock[float64]:
			if DataBlock.(Decode.AdditionalValueBlock[float64]).BlockName != "CLLS1" {
				t.Error("Имя блока  задано неверно")
				t.Log("Имя блока", DataBlock.(Decode.AdditionalValueBlock[float64]).BlockName)
			}
		}

	}

}
