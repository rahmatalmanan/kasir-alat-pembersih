package main

import "fmt"

func main() {
	// Daftar barang dan harga
	barang := []string{
		"Sapu",
		"Pel Lantai",
		"Cairan Pembersih Lantai",
		"Spons Cuci",
		"Lap Microfiber",
		"Ember",
	}

	harga := []int{
		15000,
		25000,
		20000,
		5000,
		12000,
		18000,
	}

	var pilihan, jumlah int
	var totalKeseluruhan int = 0
	var lanjut string
	var isLanjut bool = true

	for isLanjut {
		// Tampilkan daftar barang
		fmt.Println("=== Daftar Barang Alat Pembersih ===")
		for i := 0; i < len(barang); i++ {
			fmt.Printf("%d. %s - Rp %d\n", i+1, barang[i], harga[i])
		}
		fmt.Println("==========================================")

		// Validasi pilihan barang
		for {
			fmt.Print("Pilih nomor barang: ")
			fmt.Scan(&pilihan)

			if pilihan >= 1 && pilihan <= len(barang) {
				break
			} else {
				fmt.Println("Error: nomor barang tidak valid.")
			}
		}

		// Validasi jumlah
		for {
			fmt.Print("Masukkan jumlah barang: ")
			fmt.Scan(&jumlah)

			if jumlah > 0 {
				break
			} else {
				fmt.Println("Error: jumlah harus lebih dari 0.")
			}
		}

		// Hitung total
		totalBarang := harga[pilihan-1] * jumlah
		totalKeseluruhan += totalBarang

		// Tanya lanjut
		fmt.Print("Tambah barang lagi? (y/n): ")
		fmt.Scan(&lanjut)

		if lanjut != "y" {
			isLanjut = false
		}
	}
	// Output total
	fmt.Println("==============================")
	fmt.Println("Total belanja:", totalKeseluruhan)
	fmt.Println("==============================")
}
