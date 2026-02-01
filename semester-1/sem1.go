package semester1

import (
	"math"
	"strings"
)

//SOAL 1!!
func Tamasya(n int)int {
	return int(math.Ceil(float64(n) / 7.0))
}

//SOAL 2!!
func Diskon(tot int, IsExam bool)(int) {
	var TotalAkhir int
	if IsExam {
		TotalAkhir = tot - int(float64(tot) * 0.35)
		IsExam = true
	}else {
		TotalAkhir = tot
		IsExam = false
	}
	return TotalAkhir
}

//SOAL 3!!
func Konsonan(H string) (string) {
	vokal := []string{"a","i","u","e","o"}

	Hlow := strings.ToLower(H)

	for _,v := range vokal {
		if Hlow == v {
			return "Bukan konsonan"
		}
	}
	return "konsonan"
}

//KELIPATAN 3 & 5

func TriEndFaiv(n int)string  {
	tri5 := n % 5
	tri3 := n % 3

	if tri3 == 0 && tri5 == 0 {
		return "Kelipatan 5 dan 3"
	}

	if tri5 == 0{
		return "kelipatan 5"
	}
	
	if tri3 == 0 {
		return "kelipatan 3"
	}

	return "bukan kelipatan 5 atau 3"
}

//SOAL 5!!
func Max(a,b,c,d int)int {
	max := a
	if b > max {
		max = b
	}

	if c > max {
		max = c
	}

	if d > max {
		max = d
	}
	return max
}

func Min(a,b,c,d int)int {
	min := a

	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	if d < min {
		min = d
	}

	return min
}
