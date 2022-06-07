package Decode

type DataBlock PositionInfoBlock
type DecodePackage struct {
	UUID        string
	PackageSize uint32
	Time        uint32
	ByteMask    uint32
	DataBlocks  []BaseDataBlock
}

type BaseDataBlock struct {
	BlockType     uint16
	BlockSize     uint32
	Hidden        byte
	BlockDataType byte
	BlockName     string
	Block         DataBlock
}
type PositionInfoBlock struct {
	Longitude       uint32
	Latitude        uint32
	Height          uint32
	Speed           uint16
	Course          uint16
	SatellitesCount byte
}

type ValueInfoBlock struct {
	Parameter string
	Value     uint32
}
