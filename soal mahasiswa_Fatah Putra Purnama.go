package main

import "fmt"

type mahasiswa struct {
	nama           string
	kuis, uts, uas float64
	nilai_akhir    float64
}

const nmax int = 1000

type kelas [nmax]mahasiswa

func main() {
	var data kelas
	var N, i int
	var nama string

	input_data(&data, &N)

	fmt.Scan(&nama)
	i = cari_data(data, N, nama)
	hitung_nilai_akhir(&data, N)
	if i != -1 {
		fmt.Println("Nilai akhir dari", data[i].nama, "adalah", data[i].nilai_akhir)
	}
	sort_data(&data, N)
	fmt.Println("urutan ranking 1-3")
	print_data(data, N)

}
func input_data(T *kelas, N *int) {
	//I.S data siswa telah siap pada piranti masukan
	//proses : melakukan proses masukan dari user hingga nama yg diinputkan adalah "NONE" atau array telah penuh
	//F.S array T berisi sejumlah N mahasiswa
	var nama string
	*N = 0
	fmt.Scan(&nama)
	for nama != "NONE" && *N < nmax {
		T[*N].nama = nama
		fmt.Scan(&T[*N].kuis, &T[*N].uts, &T[*N].uas)
		*N = *N + 1
		fmt.Scan(&nama)
	}
}

func hitung_nilai_akhir(T *kelas, N int) {
	//I.S terdefinisi array T yang berisi N data mahasiswa
	//proses, nilai akhir = 20% kuis + 40%uts + 40%uas
	//F.S field NilaiAkhir pada array T berisi nilai akhir yg dihutung dari nilai kuis, uts, uas
	var i int
	i = 0
	for i < N {
		T[i].nilai_akhir = 0.2*T[i].kuis + 0.4*T[i].uts + 0.4*T[i].uas
		i = i + 1
	}
}

func cari_data(T kelas, N int, nama string) int {
	//mengembalikan indeks data mahasiswa yang dicari berdasarkan inputan nama, atau -1 apabila tidak ditemukan
	var ketemu, i int
	ketemu = -1
	i = 0
	for ketemu == -1 && i < N {
		if T[i].nama == nama {
			ketemu = i
		}
		i = i + 1
	}
	return ketemu
}

func sort_data(T *kelas, N int) {
	//I.S terdefinisi array T yang berisi N data mahasiswa
	//F.S data array T terurut mengecil berdasarkan nilai akhir dengan algoritma insertion sort atau selection sort
	var i, pass int
	var temp mahasiswa

	for pass = 1; pass <= N-1; pass++ {
		//mencari posisi dengan proses copy
		i = pass
		temp = T[pass]
		for i > 0 && temp.nilai_akhir > T[i-1].nilai_akhir {
			T[i] = T[i-1]
			i = i - 1
		}
		//proses penyisipan
		T[i] = temp
	}
}

func print_data(T kelas, N int) {
	//I.S terdefinisi array T yang berisi N data mahasiswa
	//F.S tercetak top-3 data siswa dengan nilai akhir tertinggi
	var i int
	i = 0
	for i < 3 {
		fmt.Println(T[i].nama, " ranking ", i+1, " dengan nilai akhir ", T[i].nilai_akhir)
		i = i + 1
	}
}
