package Encode

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

func Encode(Package EncodePackage, TimeOffset uint32) string {
	BufferPackage := new(bytes.Buffer)                          //Буфер закодированого пакета
	PackageSize := make([]byte, 4)                              //Временное хранилище для размера пакета
	binary.LittleEndian.PutUint32(PackageSize, uint32(116))     //Записывается верно , но высчитывается неверно
	BufferPackage.Write(PackageSize)                            //Записываем размер блока
	BufferPackage.WriteString(Package.UID)                      //Записываем UID
	BufferPackage.WriteByte(0x0)                                //Записываем окончание UID
	UnixTime := make([]byte, 4)                                 //Временное хранилище времени
	binary.BigEndian.PutUint32(UnixTime, Package.TS+TimeOffset) //Получаем дату и пребавляем сдвиг OmnicommTime

	BufferPackage.Write(UnixTime) //Записываем время в UNIX Time
	ButeMask := make([]byte, 4)   //Временное хранилище битовой маски

	binary.BigEndian.PutUint32(ButeMask, 3) //Записываем битовую маску
	BufferPackage.Write(ButeMask)
	return hex.EncodeToString(BufferPackage.Bytes())

}
