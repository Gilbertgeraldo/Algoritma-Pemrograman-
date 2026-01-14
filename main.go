package main

import (
	"fmt"
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
func main() {
	bilangan()
}
