package main

import (
	"errors"
	"fmt"
	"strings"
)

func array1() {
	var xs = "123"
	for i, v := range xs {
		fmt.Println("Index ", i, "Value :", v)
	}
}

func array2() {
	var ys = [5]int{10, 20, 30, 40, 50}
	for _, s := range ys {
		fmt.Println("value = ", s)
	}
}

func array3() {
	var arr = [...]string{"abed", "agape", "atun", "titit"}
	for _, val := range arr {
		fmt.Println(val)
	}
}

func ASCII_comp() {
	kodeASCII := []byte{71, 111, 108, 97, 110, 103}
	strng := string(kodeASCII)
	fmt.Println(strng)
}

func make_array() {
	var input int
	fmt.Print("masukan total array yang ingin anda tambahkan : ")
	fmt.Scan(&input)
	arrN := make([]int, 0)

	for i := 1; i <= input; i++ {
		var nilai int
		fmt.Print("masukan array ke-", i, ": ")
		fmt.Scan(&nilai)
		arrN = append(arrN, nilai)
	}
	fmt.Println("array anda sekarang adalah : ", arrN)
}

func array4() {
	var numbers1 = [2][3]int{{3, 2, 5}, {3, 2, 5}}
	var numbers2 = [2][3]int{{3, 2, 5}, {3, 2, 5}}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
}

func slicing() {
	var fruits = []string{"apple", "pepet", "agape", "pipit"}
	var afruits = fruits[1:3]
	var bfruits = fruits[1:4]

	fmt.Println(fruits)
	fmt.Println(afruits)
	fmt.Println(bfruits)
}

func mapping() {
	chicken := map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari : ", chicken["januari"])
	fmt.Println("februari : ", chicken["februari"])
}

func checking() {
	urls := []string{"site1.com", "site2.com"}
	for _, val := range urls {
		fmt.Println(val)
	}
}

func mapping_12() {
	identity := map[string]int{
		"john":                   15,
		"kennedy":                24,
		"luhut binsar panjaitan": 70,
	}
	fmt.Println(identity["luhut binsar panjaitan"])
	fmt.Println(identity["kennedy"])
	fmt.Println(identity["john"])
}

func projek_1() {
	var pilihan int

	for {
		fmt.Println("\nSelamat datang di proyek pertama saya...")
		fmt.Println("projek saya ini sekarang saya ingin melatih pemahaman saya tentang bahasa golang ")
		fmt.Println("Silahkan pilih jenis projek dibawah ini ya... : ")
		fmt.Println("1. Membuat array")
		fmt.Println("2. Membuat slicing")
		fmt.Println("3. Mapping dalam bahasa golang")
		fmt.Println("4. Keluar")
		fmt.Print("Silahkan pilih (1/2/3/4) : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var input, arrbar int
			array_baru := []int{}
			fmt.Println("======Projek membuat array======")
			fmt.Print("Berapa banyak angka yang ingin diinput: ")
			fmt.Scan(&input)
			for i := 1; i <= input; i++ {
				fmt.Printf("masukan array ke-%d: ", i)
				fmt.Scan(&arrbar)
				array_baru = append(array_baru, arrbar)
			}
			fmt.Println("array baru anda adalah : ", array_baru)

		case 2:
			var slicingIdx, sliceDep, sliceBel, subPilihan int
			var kalimat string
			fmt.Println("=====Projek membuat slicing=====")
			fmt.Println("Sekarang aku minta kamu untuk memilih slicing nya mau dari mana..")
			fmt.Println("1.Slicing dari depan")
			fmt.Println("2.Slicing dari belakang")
			fmt.Println("3.Slicing keduanya")
			fmt.Print("Pilihan: ")
			fmt.Scan(&subPilihan)

			if subPilihan == 1 {
				fmt.Print("Masukan kalimat : ")
				fmt.Scan(&kalimat)
				fmt.Print("masukan indeks awal slicing: ")
				fmt.Scan(&slicingIdx)
				if slicingIdx < len(kalimat) {
					slice_baru := kalimat[slicingIdx:]
					fmt.Println("slicing kalimat : ", kalimat, "adalah : ", slice_baru)
				}
			} else if subPilihan == 2 {
				fmt.Print("Masukan kalimat : ")
				fmt.Scan(&kalimat)
				fmt.Print("masukan indeks akhir slicing: ")
				fmt.Scan(&slicingIdx)
				if slicingIdx <= len(kalimat) {
					slice_baru := kalimat[:slicingIdx]
					fmt.Println("slicing kalimat : ", kalimat, "adalah : ", slice_baru)
				}
			} else if subPilihan == 3 {
				fmt.Print("Masukan kalimat : ")
				fmt.Scan(&kalimat)
				fmt.Print("masukan slicing dari depan : ")
				fmt.Scan(&sliceDep)
				fmt.Print("Masukan slicing dari belakang : ")
				fmt.Scan(&sliceBel)
				if sliceDep < sliceBel && sliceBel <= len(kalimat) {
					slice_baru := kalimat[sliceDep:sliceBel]
					fmt.Println("slicing kalimat : ", kalimat, "adalah : ", slice_baru)
				}
			} else {
				fmt.Println("Error: Pilihan tidak valid")
			}

		case 3:
			var tipeInput string
			fmt.Println("=====Mapping dalam bahasa golang=====")
			fmt.Print("Apakah anda ingin menginput string/integer/both? ")
			fmt.Scan(&tipeInput)

			if tipeInput == strings.ToLower("string") || tipeInput == strings.ToUpper("string") {
				map_bar := make(map[string]string)
				var jumlah int
				fmt.Print("Masukan banyaknya indeks: ")
				fmt.Scan(&jumlah)
				for i := 1; i <= jumlah; i++ {
					var k, v string
					fmt.Printf("Masukan key ke-%d: ", i)
					fmt.Scan(&k)
					fmt.Printf("Masukan value ke-%d: ", i)
					fmt.Scan(&v)
					map_bar[k] = v
				}
				fmt.Println("Map string yang sudah dibuat : ", map_bar)

			} else if tipeInput == strings.ToLower("integer") || tipeInput == strings.ToUpper("integer") {
				map_bar1 := make(map[int]int)
				var jumlah int
				fmt.Print("Masukan banyaknya indeks : ")
				fmt.Scan(&jumlah)
				for i := 1; i <= jumlah; i++ {
					var k, v int
					fmt.Printf("Masukan key ke-%d: ", i)
					fmt.Scan(&k)
					fmt.Printf("Masukan value ke-%d: ", i)
					fmt.Scan(&v)
					map_bar1[k] = v
				}
				fmt.Println("Map integer : ", map_bar1, " Panjang: ", len(map_bar1))

			} else if tipeInput == strings.ToLower("both") || tipeInput == strings.ToUpper("both") {
				map_bar2 := make(map[string]int)
				var jumlah int
				fmt.Print("Masukan banyaknya indeks : ")
				fmt.Scan(&jumlah)
				for i := 1; i <= jumlah; i++ {
					var k string
					var v int
					fmt.Printf("Masukan key (string) ke-%d: ", i)
					fmt.Scan(&k)
					fmt.Printf("Masukan value (int) ke-%d: ", i)
					fmt.Scan(&v)
					map_bar2[k] = v
				}
				fmt.Println("Map campuran : ", map_bar2, " Panjang: ", len(map_bar2))
			} else {
				fmt.Println("Error: Tipe tidak dikenali")
			}
		case 4:
			fmt.Println("Keluar sistem...")
			return
		default:
			fmt.Println("Tidak ada dalam sistem")
		}
	}
}

func belajar_fungsi(nama,kampus string, umur int) (string,string,int) {
	return nama,kampus,umur
}

func validasi_pendaftaran(nama,kampus string, umur int)(string,string,error) {
	if umur < 17 { 
		return "","",errors.New("Maaf,umur anda belum cukup untuk melakukan pendaftaran ini...")
	}

	if nama == "" || kampus == "" {
		return "","",errors.New("Nama dan kampus tidak boleh kosong...")
	}
	namaFormat := strings.ToUpper(nama)
	kampusFormat := "Universitas " + strings.Title(strings.ToLower(kampus))

	return namaFormat, kampusFormat, nil	
}

func variadic(angka ...int)int {
	total := 0
	for _, v := range angka {
		total += v
	}
	return total
}


func main() {
	// nama,kampus,umur := belajar_fungsi("Gilbert Geraldo","ITB",19)
	// fmt.Println("nama",nama)
	// fmt.Println("Kampus",kampus)
	// fmt.Println("umur",umur)

	// var inputNama, inputKampus string
	// var inputUmur int

	// fmt.Println("=====Formulir pendaftaran mahasiswa=====")
	// fmt.Print("Masukan nama lengkap anda : ")
	// fmt.Scan(&inputNama)
	// fmt.Print("Masukan nama kampus anda : ")
	// fmt.Scan(&inputKampus)
	// fmt.Println("Masukan umur anda : ")
	// fmt.Scan(&inputUmur)

	// namaRes,kampusRes,err := validasi_pendaftaran(inputNama,inputKampus,inputUmur)

	// if err != nil {
	// 	fmt.Println("\n[GAGAL DAFTAR]:", err)
	// } else {
	// 	fmt.Println("\n[BERHASIL DAFTAR]")
	// 	fmt.Println("Selamat datang,", namaRes)
	// 	fmt.Println("Anda telah terdaftar di", kampusRes)
	// 	fmt.Printf("Data anda telah tersimpan dengan umur %d tahun.\n", inputUmur)
	// }

}