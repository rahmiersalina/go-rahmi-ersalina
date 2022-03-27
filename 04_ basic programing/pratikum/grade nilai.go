package main

import "fmt"

func main() {
	var nilai int
	nilai = 90
	fmt.Println("nilai : ", nilai)
	kategorinilai(nilai)
}
func kategorinilai(nilai int) (hasil string) {
	if nilai >= 78 {
		hasil = "A"
	} else if nilai >= 65 {
		hasil = "B"
	} else if nilai >= 50 {
		hasil = "C"
	} else if nilai >= 35 {
		hasil = "D"
	} else if nilai >= 0 {
		hasil = "E"
	}
	fmt.Println("kategori : ", hasil)
	return hasil
}
