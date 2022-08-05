package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type ProgramHeader struct {
	Type   uint64 `gasm:"4" desc:"Type of segment"`              // type of segment
	Flags  uint64 `gasm:"4" desc:"Flags"`                        // flags
	Offset uint64 `gasm:"8" desc:"Offset to this segment"`       // offset
	Vaddr  uint64 `gasm:"8" desc:"Virtual address"`              // virtual address
	Paddr  uint64 `gasm:"8" desc:"Physical address"`             // physical address
	Filesz uint64 `gasm:"8" desc:"Size of segment in file"`      // size of segment in file
	Memsz  uint64 `gasm:"8" desc:"Size of segment in memory"`    // size of segment in memory
	Align  uint64 `gasm:"8" desc:"Alignment of segment in file"` // alignment of segment in file
}

func (p ProgramHeader) print() {
	val := reflect.ValueOf(&p).Elem()
	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("%s: %d\n", val.Type().Field(i).Tag.Get("desc"), val.Field(i).Interface())
	}
}

func ParsePHFromBytes(offset int, bytes []byte) ProgramHeader {
	a := &ProgramHeader{}
	val := reflect.ValueOf(a).Elem()
	for i := 0; i < val.NumField(); i++ {
		size, _ := strconv.Atoi(val.Type().Field(i).Tag.Get("gasm"))
		byteValue := ByteArrayToInt(bytes[offset : offset+size])
		offset += size
		val.Field(i).Set(reflect.ValueOf(byteValue))
	}
	return *a

}
