package main

import (
	"errors"
	"fmt"
)

type Mahasiswa struct {
	ID       int
	Nim      string
	Nama     string
	Angkatan int
	Ipk      float64
}

func main() {
	// === Bagian 1: Struct & Slice ===
	daftar := []Mahasiswa{
		{ID: 1, Nim: "H1A123001", Nama: "Andi Prasetyo", Angkatan: 2023, Ipk: 3.45},
		{ID: 2, Nim: "H1A123002", Nama: "Bunga Lestari", Angkatan: 2023, Ipk: 3.80},
		{ID: 3, Nim: "H1A124003", Nama: "Citra Ramadhani", Angkatan: 2024, Ipk: 3.20},
	}

	fmt.Println("=== Daftar Mahasiswa ===")
	for _, m := range daftar {
		fmt.Printf("%d | %s | %s | %d | %.2f\n", m.ID, m.Nim, m.Nama, m.Angkatan, m.Ipk)
	}

	// === Bagian 2: Map (hitung per angkatan) ===
	jumlahPerAngkatan := map[int]int{}
	for i := 0; i < len(daftar); i++ {
		jumlahPerAngkatan[daftar[i].Angkatan] = jumlahPerAngkatan[daftar[i].Angkatan] + 1
	}
	fmt.Println("\nJumlah per angkatan:", jumlahPerAngkatan)

	// === Bagian 3: Penanganan Error ===
	hasil, err := cariMahasiswa(daftar, 99)
	if err != nil {
		fmt.Println("\nError:", err)
	} else {
		fmt.Println("\nDitemukan:", hasil.Nama)
	}

	hasil2, err2 := cariMahasiswa(daftar, 1)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Ditemukan:", hasil2.Nama)
	}
}

// Fungsi yang mengembalikan error sebagai nilai terakhir
func cariMahasiswa(daftar []Mahasiswa, id int) (Mahasiswa, error) {
	for _, m := range daftar {
		if m.ID == id {
			return m, nil
		}
	}
	return Mahasiswa{}, errors.New("mahasiswa tidak ditemukan")
}
