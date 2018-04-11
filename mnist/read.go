package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadMNISTDB(imgFunc func(rows, cols int, label byte, b []byte)) {
	fLbl, e := os.Open("train-labels-idx1-ubyte")
	if e != nil {
		log.Panic(e)
	}
	defer fLbl.Close()

	fImg, e := os.Open("train-images-idx3-ubyte")
	if e != nil {
		log.Panic(e)
	}
	defer fImg.Close()

	// Read magic number
	if magic := ReadInt32(fLbl); magic != 2049 {
		log.Panicf("Salah magic number")
	}

	if magic := ReadInt32(fImg); magic != 2051 {
		log.Panicf("Salah magic number")
	}

	// Read number of data
	num := ReadInt32(fLbl)
	if n := ReadInt32(fImg); n != num {
		log.Panicf("Jml label dan image beda: %d != %d", num, n)
	}

	// Read image properties (rows and columns)
	rows := ReadInt32(fImg)
	cols := ReadInt32(fImg)

	// Read data
	size := rows * cols
	b := make([]byte, size)
	for i := 0; i < int(num); i++ {
		fmt.Println(i)

		// Read label
		lbl := Read1Byte(fLbl)

		// Read image
		n, e := fImg.Read(b)
		if e != nil {
			log.Panic(e)
		}
		if n != int(size) {
			log.Panicf("Yang dibaca %d != %d", n, size)
		}

		// Call
		imgFunc(int(rows), int(cols), lbl, b)
	}
}

func ReadInt32(rd io.Reader) (i int32) {
	if e := binary.Read(rd, binary.BigEndian, &i); e != nil {
		log.Panic(e)
	}
	return
}

func Read1Byte(rd io.Reader) byte {
	lbl := []byte{0}
	n, e := rd.Read(lbl)
	if e != nil {
		log.Panic(e)
	}
	if n != 1 {
		log.Panicf("Yang dibaca cuma 1 byte saja kok ga bisa?")
	}
	return lbl[0]
}
