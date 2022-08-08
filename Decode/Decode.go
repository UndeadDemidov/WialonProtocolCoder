package Decode

import (
	"encoding/binary"
	"fmt"
	"math"
)

func BytesToFloat(Bytes []byte) float64 {

	Uint := binary.LittleEndian.Uint64(Bytes)

	return math.Float64frombits(Uint)
}
func Decode(BytesArray []byte) DecodePackage {

	var DecodePackage DecodePackage
	var Offset = 0
	DecodePackage.PackageSize = binary.LittleEndian.Uint32(BytesArray[:4])
	Offset = 4

	for Index := Offset; Index < len(BytesArray); Index++ {

		if BytesArray[Index] == 0x0 {
			DecodePackage.UUID = string(BytesArray[Offset:Index])
			Offset = Index + 1
			break

		}
	}

	DecodePackage.UnixTime = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
	Offset = Offset + 4
	DecodePackage.ByteMask = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
	Offset = Offset + 4
	for Offset < len(BytesArray) {
		var NewBaseBlock BaseBlock
		NewBaseBlock.BlockType = binary.BigEndian.Uint16(BytesArray[Offset : Offset+2])
		Offset = Offset + 2
		NewBaseBlock.BlockSize = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
		Offset = Offset + 4
		NewBaseBlock.Hidden = BytesArray[Offset]
		Offset = Offset + 1
		NewBaseBlock.BlockDataType = BytesArray[Offset]
		Offset = Offset + 1
		for Index := Offset; Index < len(BytesArray); Index++ {

			if BytesArray[Index] == 0x0 {
				NewBaseBlock.BlockName = string(BytesArray[Offset:Index])
				Offset = Index + 1
				break

			}

		}
		fmt.Println(NewBaseBlock.BlockName)
		switch NewBaseBlock.BlockName {
		case "posinfo":

			NewPositionInfoBlock := PositionInfoBlock{BaseBlock: NewBaseBlock}

			NewPositionInfoBlock.Longitude = BytesToFloat(BytesArray[Offset : Offset+8])
			Offset = Offset + 8
			NewPositionInfoBlock.Latitude = BytesToFloat(BytesArray[Offset : Offset+8])
			Offset = Offset + 8
			NewPositionInfoBlock.Height = BytesToFloat(BytesArray[Offset : Offset+8])
			Offset = Offset + 8
			NewPositionInfoBlock.Speed = binary.BigEndian.Uint16(BytesArray[Offset : Offset+2])
			Offset = Offset + 2
			NewPositionInfoBlock.Course = binary.BigEndian.Uint16(BytesArray[Offset : Offset+2])
			Offset = Offset + 2
			NewPositionInfoBlock.SatellitesCount = BytesArray[Offset]
			Offset = Offset + 1
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewPositionInfoBlock)
			break
		case "adc1":

			NewSensorsInfoBlock := AdditionalValueBlock[uint32]{BaseBlock: NewBaseBlock}
			NewSensorsInfoBlock.Value = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
			fmt.Println(NewSensorsInfoBlock.Value)
			Offset = Offset + 4
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewSensorsInfoBlock)
			break

		case "adc2":

			NewSensorsInfoBlock := AdditionalValueBlock[uint32]{BaseBlock: NewBaseBlock}
			NewSensorsInfoBlock.Value = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
			fmt.Println(NewSensorsInfoBlock.Value)
			Offset = Offset + 4
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewSensorsInfoBlock)
			break

		case "mileage":

			NewSensorsInfoBlock := AdditionalValueBlock[float64]{BaseBlock: NewBaseBlock}
			NewSensorsInfoBlock.Value = BytesToFloat(BytesArray[Offset : Offset+8])
			fmt.Println(NewSensorsInfoBlock.Value)
			Offset = Offset + 8
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewSensorsInfoBlock)
			break

		case "batlife":

			NewSensorsInfoBlock := AdditionalValueBlock[uint32]{BaseBlock: NewBaseBlock}
			NewSensorsInfoBlock.Value = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
			fmt.Println(NewSensorsInfoBlock.Value)
			Offset = Offset + 4
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewSensorsInfoBlock)
			break

		case "TImp":

			NewSensorsInfoBlock := AdditionalValueBlock[uint32]{BaseBlock: NewBaseBlock}
			NewSensorsInfoBlock.Value = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
			fmt.Println(NewSensorsInfoBlock.Value)
			Offset = Offset + 4
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewSensorsInfoBlock)
			break

		case "Uboard":

			NewSensorsInfoBlock := AdditionalValueBlock[float64]{BaseBlock: NewBaseBlock}
			NewSensorsInfoBlock.Value = BytesToFloat(BytesArray[Offset : Offset+8])
			fmt.Println(NewSensorsInfoBlock.Value)
			Offset = Offset + 8
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewSensorsInfoBlock)
			break

		case "CLLS1":

			NewSensorsInfoBlock := AdditionalValueBlock[uint32]{BaseBlock: NewBaseBlock}
			NewSensorsInfoBlock.Value = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
			fmt.Println(NewSensorsInfoBlock.Value)
			Offset = Offset + 4
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewSensorsInfoBlock)
			break
		}

	}
	return DecodePackage

}
