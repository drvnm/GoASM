package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type ELFHeader struct {
	EMagic      uint64 `gasm:"4" desc:"Magic elf number"`            // magic number
	EClass      uint64 `gasm:"1" desc:"32/64 bit"`                   // 1:32bit, 2:64bit
	EData       uint64 `gasm:"1" desc:"Endiannes"`                   // 1:little endian, 2:big endian
	EVersion    uint64 `gasm:"1" desc:"ELF version"`                 // version
	EOSABI      uint64 `gasm:"1" desc:"ABI type"`                    // os/abi
	EABIVersion uint64 `gasm:"1" desc:"ABI version"`                 // abi version
	EPad        uint64 `gasm:"7" desc:"padding"`                     // padding
	Type        uint64 `gasm:"2" desc:"Type of file"`                // type of file (1:relocatable, 2:executable, 3:shared object, 4:core file)
	Machine     uint64 `gasm:"2" desc:"Machine type"`                // machine
	Version     uint64 `gasm:"4" desc:"Version"`                     // elf Version
	Entry       uint64 `gasm:"8" desc:"Entry point instructions"`    // entry point
	Phoff       uint64 `gasm:"8" desc:"Program header offset"`       // program header offset
	Shoff       uint64 `gasm:"8" desc:"Section header offset"`       // section header offset
	Flags       uint64 `gasm:"4" desc:"Flags"`                       // flags
	Ehsize      uint64 `gasm:"2" desc:"ELF header size"`             // elf header size
	Phentsize   uint64 `gasm:"2" desc:"Program header size"`         // program header size
	Phnum       uint64 `gasm:"2" desc:"Program header count"`        // program header count
	Shentsize   uint64 `gasm:"2" desc:"Section header size"`         // section header size
	Shnum       uint64 `gasm:"2" desc:"Section header count"`        // section header count
	Shstrndx    uint64 `gasm:"2" desc:"Section header string index"` // section header string index
}

func (e ELFHeader) print() {
	val := reflect.ValueOf(&e).Elem()
	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("%s: %d\n", val.Type().Field(i).Tag.Get("desc"), val.Field(i).Interface())
	}
}

func ParseEHFromBytes(bytes []byte) ELFHeader {
	a := &ELFHeader{}
	val := reflect.ValueOf(a).Elem()
	offset := 0
	for i := 0; i < val.NumField(); i++ {
		size, _ := strconv.Atoi(val.Type().Field(i).Tag.Get("gasm"))
		offset += size
		byteValue := ByteArrayToInt(bytes[offset-size : offset])
		val.Field(i).Set(reflect.ValueOf(byteValue))
	}
	return *a
}
