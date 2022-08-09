package Tests

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/MrAmperage/WialonProtocolCoder/Decode"
)

func TestDecodePackage(t *testing.T) {
	Message, Error := hex.DecodeString("7b0000003330333031393034390062f1bb70000000030bbb000000270102706f73696e666f0091e22da14caa554058b4ef4053ff4a409a99999999e96340000000be130bbb0000000f0104616463310000000000008066400bbb0000000f010461646332000000000000408f400bbb0000000c0103434c4c53310000000000")
	if Error != nil {
		fmt.Println(Error)
	}
	DecodePackage := Decode.Decode(Message)
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
			if DataBlock.(Decode.AdditionalValueBlock[uint32]).BlockName != "adc1" {
				t.Error("Имя блока  задано неверно")
				t.Log("Имя блока", DataBlock.(Decode.AdditionalValueBlock[uint32]).BlockName)
			}
		}

	}

}
