package main

import "fmt"

func ubahMap(data map[string]string) {
	data["status"] = "sudah paham"
}

func ubahSlice(angka []int) {
	angka[0] = 200
}

func main() {
	catatanBelajar := map[string]string {
		"status":"belum paham",
	}
	ubahMap(catatanBelajar)
	fmt.Println(catatanBelajar)

	kumpulanAngka := []int{1, 2, 3}
	ubahSlice(kumpulanAngka)
	fmt.Println(kumpulanAngka)
}