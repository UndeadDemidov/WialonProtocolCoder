package Decode

type DecodePackage struct {
	UUID        string
	PackageSize uint32
	UnixTime    uint32
	ByteMask    uint32
	DataBlocks  []PositionInfoBlock
}

type PositionInfoBlock struct {
	BlockType       uint16
	BlockSize       uint32
	Hidden          byte
	BlockDataType   byte
	BlockName       string
	Longitude       float64
	Latitude        float64
	Height          float64
	Speed           uint16
	Course          uint16
	SatellitesCount byte
}
