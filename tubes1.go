package main

import (
	"errors"
	"fmt"
	"strings"
)

func tambahProduk(stok map[string]int, harga map[string]int, nama string, qty int, hrg int) string {
	stok[nama] += qty
	harga[nama] = hrg

	return fmt.Sprintf("Sukses!! %s, sekarang berjumlah %d ", nama, stok[nama])	
}

func hapusProduk(stok map[string]int, harga map[string]int, nama string) bool {
	if _,ada := stok[nama];ada {
		delete(stok,nama)
		delete(harga,nama)
		return true
	}
	return false
}

func jualBarang(stok map[string]int, harga map[string]int, nama string, qty int, hrg int) (int, error) {
	if stok[nama] < qty {
		return 0, errors.New("Stok tidak mencukupi")
	}
	stok[nama] -= qty
	total := qty * harga[nama]

	return total, nil 
}

func tampilkanBarang(stok map[string]int, harga map[string]int) {
	fmt.Println("\n===list barang gudang===\n")
	fmt.Printf("%-15s | %-10s | %-10s\n", "Nama", "Harga", "Stok")
	fmt.Println("-------------------------------------------")

	for nama,jumlah := range stok {
		fmt.Printf("%-15s | Rp%-9d | %-10d\n", nama, harga[nama], jumlah)
	}
}

func cariBarang(stok map[string]int, harga map[string]int, keyword string ) []string {
	hasil := []string{}
	for nama := range stok {
		if strings.Contains(strings.ToLower(nama),strings.ToLower(keyword)) {
			hasil = append(hasil, nama)
		}
	}
	return hasil
}