package Decode

import (
	"encoding/binary"
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

	DecodePackage.Time = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
	Offset = Offset + 4
	DecodePackage.ByteMask = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
	Offset = Offset + 4
	for Index := Offset; Index < len(BytesArray); Index = Index + Offset {
		var NewDataBlock PositionInfoBlock
		NewDataBlock.BlockType = binary.BigEndian.Uint16(BytesArray[Offset : Offset+2])
		Offset = Offset + 2
		NewDataBlock.BlockSize = binary.BigEndian.Uint32(BytesArray[Offset : Offset+4])
		Offset = Offset + 4
		NewDataBlock.Hidden = BytesArray[Offset]
		Offset = Offset + 1
		NewDataBlock.BlockDataType = BytesArray[Offset]
		Offset = Offset + 1
		for Index := Offset; Index < len(BytesArray); Index++ {

			if BytesArray[Index] == 0x0 {
				NewDataBlock.BlockName = string(BytesArray[Offset:Index])
				Offset = Index + 1
				break

			}

		}

		switch NewDataBlock.BlockName {
		case "posinfo":
			NewDataBlock.Longitude = BytesToFloat(BytesArray[Offset : Offset+8])
			Offset = Offset + 8
			NewDataBlock.Latitude = BytesToFloat(BytesArray[Offset : Offset+8])
			Offset = Offset + 8
			NewDataBlock.Height = BytesToFloat(BytesArray[Offset : Offset+8])
			Offset = Offset + 8
			NewDataBlock.Speed = binary.BigEndian.Uint16(BytesArray[Offset : Offset+2])
			Offset = Offset + 2
			NewDataBlock.Course = binary.BigEndian.Uint16(BytesArray[Offset : Offset+2])
			Offset = Offset + 2
			NewDataBlock.SatellitesCount = BytesArray[Offset]
			Offset = Offset + 1
			DecodePackage.DataBlocks = append(DecodePackage.DataBlocks, NewDataBlock)

		}
	}
	return DecodePackage

}
