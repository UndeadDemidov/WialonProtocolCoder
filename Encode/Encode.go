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
	PackageSize := make([]byte, 4)                              //Временное хранилище для размера пакета
	binary.LittleEndian.PutUint32(PackageSize, uint32(116))     //Записывается верно , но высчитывается неверно
	BufferPackage.Write(PackageSize)                            //Записываем размер блока
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
	BlockSize := make([]byte, 4) //Временной хранилище размера блока
	binary.BigEndian.PutUint32(BlockSize, 39)
	BufferPackage.Write(BlockSize)
	BufferPackage.WriteByte(1) //Записываем атребут скрытости
	BufferPackage.WriteByte(2) //Записываем тип данных блока

	BufferPackage.WriteString("posinfo") //Записываем Имя блока
	BufferPackage.WriteByte(0x0)
	Longitude := make([]byte, 8) //Временное хранилище долготы
	binary.LittleEndian.PutUint64(Longitude, FloatToUint(867263737))
	BufferPackage.Write(Longitude)
	Latitude := make([]byte, 8) //Временное хранилище широты
	binary.LittleEndian.PutUint64(Latitude, FloatToUint(543304085))
	BufferPackage.Write(Latitude)
	Height := make([]byte, 8) //Временное хранилище высоты
	binary.LittleEndian.PutUint64(Height, 2843)
	BufferPackage.Write(Height)
	Speed := make([]byte, 2) //Временное хранилище скорости
	binary.BigEndian.PutUint16(Speed, 122)
	BufferPackage.Write(Speed)
	Course := make([]byte, 2) //Временное хранилище курса
	binary.BigEndian.PutUint16(Course, 343)
	BufferPackage.Write(Course)
	BufferPackage.WriteByte(19) //Записываем количество спутников

	return hex.EncodeToString(BufferPackage.Bytes())

}
