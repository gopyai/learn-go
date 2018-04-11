package main

//func main() {
//	ReadMNISTDB(func(rows, cols int, lbl byte, b []byte) {
//		img := Bytes2Image(rows, cols, b)
//		_ = img
//	})
//}

func main() {
	fIdx := 0
	ReadMNISTDB(func(rows, cols int, lbl byte, b []byte) {
		img := Bytes2Image(rows, cols, b)
		CreatePNGAndShow(img, lbl, fIdx)
		fIdx++
	})
}
