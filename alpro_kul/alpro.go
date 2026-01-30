package alpro_kul

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
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

func Io_reader() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan nama anda : ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error",nil)
	}
	fmt.Printf("Halo, %s",input)
}

func Fibonacci(n int) int {
	// if n <= 0 {
	// 	return 0
	// }

	if n <=1 { 
		return n
	}
		//pakai rekursif
	return Fibonacci(n-1) + Fibonacci(n-2)

	// fibo := make([]int,n+1)	
	// fibo[0]=0
	// fibo[1]=1

	// for i:=2;i<=n;i++{
	// 	fibo[i] = fibo[i-1] + fibo[i-2]
	// }
	// return fibo[n]
}

func Getfull() (string,string) {
	return "fayyet","ajg"
}

func Bebufio(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)

	if scanner.Scan(){
		return scanner.Text()
	}
	return ""
}

func LogInfo(pesan ...string) {
	fullMessage := strings.Join(pesan, " ")
	
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("[%s] [INFO] %s\n", timestamp, fullMessage)
}

func Saring(angka []int, kondisi func(int) bool) []int {
	var hasil []int
	for _,v  := range angka {
		if kondisi(v) {
			hasil = append(hasil,v)
		}
	}
	return hasil
}

func Filterspam(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("hello \t" + filteredName)
}

func SpamFilter(name string )string {
	if name == "anjing" {
		return "..."
	}else {
		return name
	}
}

func FuncAsParam(number []int, age func([]int)int)int {
	param := age(number)
	return param
}

func KondisiMotor(merek map[string]string, mesin func(string)bool)bool {
	for _,status := range merek {
		if mesin(status) {
			return true
		}
	}	
	return false
}