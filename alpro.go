package main

import (
	"fmt"
	"math"
)

func vtabung(jari_jari,tinggi float64) float64 {
	const pi = 3.14
	volume := pi * math.Pow(jari_jari,2) * tinggi
	return volume
}

func ltabung(jari_jari,tinggi float64) float64 {
	const pi = 3.14
	luas_lingkaran := pi * math.Pow(jari_jari,2)
	luas_selimut := 2 * pi * jari_jari * tinggi
	luas_tabung := luas_lingkaran + luas_selimut
	return luas_tabung
}

func konverWaktu (jam,menit,detik int) int {
	return jam * 3600 + menit * 60 + detik
}

func segitiga(a,b,c int) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	return a+b > c && a+c > b && b+c > a
}

func diskon(total_belanja int, member bool) int {
	if total_belanja > 100000 && member {
		return total_belanja - (total_belanja * 10 / 100)
	} else if total_belanja > 1000000 && !member {
		return total_belanja - (total_belanja * 5 / 100)
	}
	return total_belanja
}

func fibonacci(n int) int {
	// if n <= 0 {
	// 	return 0
	// }

	if n <=1 { 
		return n
	}
		//pakai rekursif
	return fibonacci(n-1) + fibonacci(n-2)

	// fibo := make([]int,n+1)	
	// fibo[0]=0
	// fibo[1]=1

	// for i:=2;i<=n;i++{
	// 	fibo[i] = fibo[i-1] + fibo[i-2]
	// }
	// return fibo[n]
}

func main() {
	// var r, t float64

	// fmt.Print("Masukkan jari-jari tabung: ")
	// fmt.Scan(&r)

	// fmt.Print("Masukkan tinggi tabung: ")
	//
	// fmt.Print("Masukkan tinggi tabung: ")
	// fmt.Scan(&t)

	// hasil := vtabung(r, t)
	// hasil2 := ltabung(r,t)
	// fmt.Println("Volume tabung:", hasil)
	// fmt.Println("luas tabung:", hasil2)

	//konversi waktu
	// var j,m,d int
	// fmt.Scan(&j,&m,&d)
	// hasil := konverWaktu(j,m,d)
	// fmt.Println("hasil konversi:",hasil,"detik")

	//segitiga
	// var a,b,c int
	// fmt.Scan(&a,&b,&c)
	
	// if segitiga(a,b,c) {
	// 	fmt.Println("Segitiga")
	// }else {
	// 	fmt.Println("Bukan Segitiga")
	// }

	//diskon
	// var total int
	// var member bool
	// fmt.Scan(&total,&member)

	// hasil := diskon(total,member)
	// fmt.Println(hasil)

	//fibonacci
	var n int
	fmt.Scan(&n)

	hasil := fibonacci(n)
	fmt.Println(hasil)
}