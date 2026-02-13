package main

import (
	// "alpro_telkom/alpro_kul"
	// "alpro_telkom/Alpro_Telkom"
	// "strings"
	// "alpro_telkom/semester-1"
	"alpro_telkom/sorting"
	"math"
	// "bufio"
	"fmt"
	// "os"
)

func soal() {
	var n int
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}

func balok() {
	var x int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	for i := x - 1; i >= 1; i-- {
		fmt.Print(i)
		for j := 1; j <= 2*x-3; j++ {
			fmt.Print(" ")
		}
		fmt.Print(i)
	}
}

func biner() {
	var n int
	fmt.Scan(&n)

	for n > 0 {
		fmt.Print(n % 2)
		n = n / 2
	}
}

func daun() {
	var n, i, x, max int
	fmt.Print("masukan angka : ")
	fmt.Scan(&n)
	fmt.Scan(&x)
	max = x
	i = 1
	for n != i {
		fmt.Scan(&x)
		if x > max {
			max = x
		}
		i++
	}

	fmt.Println("max : ", max)
}

func digit_terbesar() {
	var n int
	fmt.Scan(&n)

	maxDigit := 0
	num := n

	for num > 0 {
		digit := num % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		num /= 10
	}

	fmt.Printf("Bilangan terbesarnya adalah %d\n", maxDigit)
}

func cari_digit() {
	var x, y int
	fmt.Println("masukan angka yang anda cari : ")
	fmt.Scan(&x)
	fmt.Println("masukan angka nya : ")
	fmt.Scan(&y)

	found := false

	for y > 0 {
		sisa := y % 10
		if sisa == x {
			found = true
			break
		}
		y /= 10
	}

	if found {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func gerbang_tol() {
	var input string
	fmt.Println("masukan huruf : ")
	fmt.Scan(&input)

	countA, countB, countC := 0, 0, 0

	for _, ch := range input {
		switch ch {
		case 'A':
			countA++
		case 'B':
			countB++
		case 'C':
			countC++
		}
	}

	fmt.Println("Tipe A = ", countA)
	fmt.Println("TIpe B = ", countB)
	fmt.Println("Tipe C = ", countC)

}

func tempretatur() {
	var i, x, max, min, xi, cek int
	var mean float64
	max = x
	min = x
	for {
		fmt.Print("masukan angka : ")
		fmt.Scan(&x)

		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
		xi += x
		i++
		if x == 0 {
			fmt.Print("masukan angka : ")
			fmt.Scan(&cek)
			if cek == 0 {
				break
			} else {
				xi += cek
				i++
				continue
			}
		}
	}

	mean = float64(xi) / float64(i)

	fmt.Println("max : ", max)
	fmt.Println("min : ", min)
	fmt.Println("mean : ", mean)

}

func polbil() {
	var n int
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
			// if j == 1 || i == n || j == i {
			// 	fmt.Print("* ")
			// } else {
			// 	fmt.Print("  ")
			// }
		}
		fmt.Println("")
	}
	for i := 2; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
			// if j == 1 || i == n || j == i {
			// 	fmt.Print("* ")
			// } else {
			// 	fmt.Print("  ")
			// }
		}
		fmt.Println("")
	}
}

func polb() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}

func pol2() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n {
				fmt.Print(i, " ")
			} else if j == 1 || j == n {
				fmt.Print(i, " ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println(" ")
	}
}

func konsekutif() {
	var x int
	var sel bool
	fmt.Scan(&x)

	sel = true
	for x >= 10 && sel != false {
		digit_1 := x % 10
		digit_2 := (x / 10) % 10
		selisih := digit_1 - digit_2
		if selisih != 1 && selisih != -1 {
			sel = false
		}
		x /= 10
	}
	fmt.Println(sel)
}

func bilangan() {
	var n int 
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Println("*")
		}
		fmt.Println()
	}
}

func biblo() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		for j := 1; j <= n; j++{
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func triangle() {
	var x int
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		for j := i; j <= x; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()	 
	}
}

func param(d float64) (float64,float64) {
	var area = math.Pi * math.Pow(d / 2,2 )

	var circumences = math.Pi * d

	return area,circumences
}

func paridic(numbers ...int) float64 {
	var total int = 0

	for _,number := range numbers {
		total += number

	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func main() {
	// var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	// var Numbers = func(min int) []int {
	// 	var r []int
	// 	for _,e := range numbers {
	// 		if e < min {
	// 			continue
	// 		}
	// 		r = append(r, e)
	// 	}
	// 	return r
	// }(3)
	// fmt.Println("Original numbers : ",numbers)
	// fmt.Println("Filtered numbers : ",Numbers),

	// age := []int{1,2,3,4,5,6}

	// tambah := func(n[]int)int {
	// 	total := 0
	// 	for _,v := range n {
	// 		total += v
	// 	}
	// 	return total
	// }
	// hasil := alpro_kul.FuncAsParam(age,tambah)
	// fmt.Print(hasil)
	// daftar_motor := map[string]string{
	// 	"yamaha":"bagus",
	// 	"honda":"rusak",
	// 	"hrv" : "sedikit kerusakan",
	// }

	// rusak := alpro_kul.KondisiMotor(daftar_motor, func(s string) bool {
	// 	return s == "rusak"
	// })

	// if rusak {
	// 	fmt.Print("waspada! ada motor rusak")
	// }else {
	// 	fmt.Print("semua motor dalam keadaan aman")
	// }

//SOAL 1!!
	// var a,b,c,d int
	// fmt.Print("input : ")
	// _,err :=fmt.Scan(&a,&b,&c,&d)
	// if err != nil {
	// 	fmt.Println("error")
	// 	return
	// }
	// per1 := alpro.Permutasi(a,c)
	// com1 := alpro.Combinasi(a,c)
	// fmt.Println(per1,com1)

	// per2 := alpro.Permutasi(b,d)
	// com2 := alpro.Combinasi(b,d)
	// fmt.Println(per2,com2)

//SOAL 2!!
	// var a,b,c int
	// fmt.Scanln(&a,&b,&c)

	// fmt.Println(alpro.F(alpro.G(alpro.H(a))))
	// fmt.Println(alpro.G(alpro.H(alpro.F(b))))
	// fmt.Println(alpro.H(alpro.F(alpro.G(c))))

//SOAL 3!!

// var cx1,cy1,qr1 int 
// var cx2,cy2,qr2 int
// var x,y int

// fmt.Scan(&cx1,&cy1,&qr1)
// fmt.Scan(&cx2,&cy2,&qr2)
// fmt.Scan(&x,&y)

// hasil := alpro.Lingkaran(x,y,cx1,cx2,cy1,cy1,qr1,qr2)
// fmt.Println(hasil)

//SOAL 4!!

// var w,h int
// var s int

// fmt.Scan(&w,&h)
// fmt.Scan(&s)

// if s > 0 {
// 	fmt.Print(alpro.W(w,s),alpro.C(h,s))
// } else {
// 	fmt.Print(alpro.A(w,s),alpro.E(h,s))
// }

//SOAL 5!!

// var r,t1,rho1,t2,rho2 float64

// fmt.Scan(&r)
// fmt.Scan(&t1,&rho1)
// fmt.Scan(&t2,&rho2)

// m1 := alpro.Massa(rho1,t1,r)
// m2 := alpro.Massa(rho2,t2,r)

// selisih := math.Abs(m1 - m2)

// if selisih == 0 { 
// 	fmt.Print("BALANCE")
// }else {
// 	fmt.Printf("%.3f\n",selisih)
// }


//SOAL SEMESTER 1!!!(1-2-2026)


//SOAL 1!!

// var a int
// fmt.Scan(&a)
// fmt.Print(semester1.Tamasya(a))

//SOAL 2!!

// var tot int
// var IsE bool
// fmt.Scan(&tot,&IsE)
// fmt.Print(semester1.Diskon(tot,IsE))

//SOAL 3!!

// var n string
// fmt.Scan(&n)
// fmt.Print(semester1.Konsonan(n))

//SOAL 4!!

// var n int
// fmt.Scan(&n)
// fmt.Print(semester1.TriEndFaiv(n))

//SOAL 5!!

// var a,b,c,d int
// fmt.Scan(&a,&b,&c,&d)
// maxi := semester1.Max(a,b,c,d)
// mini := semester1.Min(a,b,c,d)
// fmt.Println(maxi,mini)

//SOAL 6!!

// var a,b,c int
// fmt.Scan(&a,&b,&c)
// fmt.Print(semester1.Segitiga(a,b,c))

// var a,b float64
// fmt.Scan(&a,&b)
// fmt.Print(semester1.Profit(a,b))


//SOAL 7!!

// var h1,m1,h2,m2 int
// var per string
// fmt.Scan(&h1,&m1,&h2,&m2)
// fmt.Println(semester1.Parki(h1,m1,h2,m2,per))

//BAB REKURSIF

//SOAL 1!!

// var n int
// fmt.Scan(&n)
// fmt.Println(alpro.Ganjil(n))

//SOAL 2!!

// var n int
// fmt.Scan(&n)
// alpro.Faktor(n)

//SOAL 3!!

// var x,y int
// fmt.Scan(&x,&y)
// fmt.Print(alpro.Powa(x,y))




//BUBBLE SORT!!
var arr = []int{1,39,2,9,7,54,11}

sorting.Bubblesort(arr)
fmt.Println(arr)

sorting.RecBubbleSOrt(arr,5)
fmt.Println(arr)

sorting.InsertionSort(arr)
fmt.Println(arr)

sorting.SelectionSort(arr)
fmt.Println(arr)

}