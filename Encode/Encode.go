package Encode

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math"
)

func FloatToUint(Number float64) uint64 {
	return math.Float64bits(Number)

}
func Encode(Package EncodePackage, TimeOffset uint32) string {
	BufferPackage := new(bytes.Buffer)                          //Буфер закодированого пакета
	BufferPackage.WriteString(Package.UID)                      //Записываем UID
	BufferPackage.WriteByte(0x0)                                //Записываем окончание UID
	UnixTime := make([]byte, 4)                                 //Временное хранилище времени
	binary.BigEndian.PutUint32(UnixTime, Package.TS+TimeOffset) //Получаем дату и пребавляем сдвиг OmnicommTime

	BufferPackage.Write(UnixTime)           //Записываем время в UNIX Time
	ButeMask := make([]byte, 4)             //Временное хранилище битовой маски
	binary.BigEndian.PutUint32(ButeMask, 3) //Записываем битовую маску
	BufferPackage.Write(ButeMask)
	BlockType := make([]byte, 2)
	binary.BigEndian.PutUint16(BlockType, 3003)
	BufferPackage.Write(BlockType)
	//Начало записи блока posinfo
	BlockBuffer := new(bytes.Buffer)
	BlockBuffer.WriteByte(1)           //Записываем атрибут скрытости
	BlockBuffer.WriteByte(2)           //Записываем тип данных блока
	BlockBuffer.WriteString("posinfo") //Записываем Имя блока
	BlockBuffer.WriteByte(0x0)
	Longitude := make([]byte, 8) //Временное хранилище долготы
	binary.LittleEndian.PutUint64(Longitude, FloatToUint(float64(Package.Lon.Int64)*10e-8))
	BlockBuffer.Write(Longitude)
	Latitude := make([]byte, 8) //Временное хранилище широты
	binary.LittleEndian.PutUint64(Latitude, FloatToUint(float64(Package.Lat.Int64)*10e-8))
	BlockBuffer.Write(Latitude)
	Height := make([]byte, 8) //Временное хранилище высоты
	binary.LittleEndian.PutUint64(Height, FloatToUint(float64(Package.GPSAlt.Int64)*0.1))
	BlockBuffer.Write(Height)
	Speed := make([]byte, 2) //Временное хранилище скорости
	binary.BigEndian.PutUint16(Speed, uint16(Package.Speed.Int64/10))
	BlockBuffer.Write(Speed)
	Course := make([]byte, 2) //Временное хранилище курса
	binary.BigEndian.PutUint16(Course, uint16(Package.Dir.Int64))
	BlockBuffer.Write(Course)
	BlockBuffer.WriteByte(byte(Package.Sat.Int64)) //Записываем количество спутников
	//Запись длинны блока
	BlockLength := BlockBuffer.Len()
	BlockString := BlockBuffer.String()
	BlockBuffer.Reset()
	BlockSize := make([]byte, 4)
	binary.BigEndian.PutUint32(BlockSize, uint32(BlockLength))
	BufferPackage.Write(BlockSize)
	BufferPackage.WriteString(BlockString)

	//Блок с первым сенсором
	BlockBuffer.Reset()
	BlockBuffer.WriteByte(1)        //Записываем атрибут скрытости
	BlockBuffer.WriteByte(4)        //Записываем тип данных блока
	BlockBuffer.WriteString("adc1") //Записываем Имя блока
	BlockBuffer.WriteByte(0x0)
	Sensor := make([]byte, 4) //Временное хранилище датчика1
	binary.BigEndian.PutUint32(Sensor, uint32(Package.Unval0.Int64))
	BlockBuffer.Write(Sensor)
	//Запись длинны блока
	BlockLength = BlockBuffer.Len()
	BlockBytes := BlockBuffer.Bytes()
	BlockBuffer.Reset()
	BlockType = make([]byte, 2)
	binary.BigEndian.PutUint16(BlockType, 3003)
	BufferPackage.Write(BlockType)
	BlockSize = make([]byte, 4)
	binary.BigEndian.PutUint32(BlockSize, uint32(BlockLength))
	BufferPackage.Write(BlockSize)
	BufferPackage.Write(BlockBytes)

	//Блок с вторым сенсором
	BlockBuffer.Reset()
	BlockBuffer.WriteByte(1)        //Записываем атрибут скрытости
	BlockBuffer.WriteByte(4)        //Записываем тип данных блока
	BlockBuffer.WriteString("adc2") //Записываем Имя блока
	BlockBuffer.WriteByte(0x0)
	Sensor = make([]byte, 4) //Временное хранилище датчика2
	binary.BigEndian.PutUint32(Sensor, uint32(Package.Unval1.Int64))
	BlockBuffer.Write(Sensor)
	//Запись длинны блока
	BlockLength = BlockBuffer.Len()
	BlockBytes = BlockBuffer.Bytes()
	BlockBuffer.Reset()
	BlockType = make([]byte, 2)
	binary.BigEndian.PutUint16(BlockType, 3003)
	BufferPackage.Write(BlockType)
	BlockSize = make([]byte, 4)
	binary.BigEndian.PutUint32(BlockSize, uint32(BlockLength))
	BufferPackage.Write(BlockSize)
	BufferPackage.Write(BlockBytes)

	//Запись длинны пакета
	PackageLength := BufferPackage.Len()
	PackageString := BufferPackage.String()
	BufferPackage.Reset()
	PackageSize := make([]byte, 4)
	binary.LittleEndian.PutUint32(PackageSize, uint32(PackageLength))
	BufferPackage.Write(PackageSize)
	BufferPackage.WriteString(PackageString)
	return hex.EncodeToString(BufferPackage.Bytes())

}
