package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	bytes, err := ioutil.ReadFile("./main")
	if err != nil {
		panic(err)
	}
	elfHeader := ParseEHFromBytes(bytes)

	// if magic number is not right, exit (we do not care about big endian)
	if elfHeader.EMagic != 1179403647 {
		panic("not an elf file")
	}

	var programHeaders []ProgramHeader
	elfHeader.print()
	for i := 0; i < int(elfHeader.Phnum); i++ {
		programHeaders = append(programHeaders, ParsePHFromBytes(64+i*int(elfHeader.Phentsize), bytes))
	}
	for _, programHeader := range programHeaders {
		programHeader.print()
	}

	intstructionStart := elfHeader.Phoff + elfHeader.Phentsize*elfHeader.Phnum
	fmt.Println(bytes[intstructionStart:])
}
