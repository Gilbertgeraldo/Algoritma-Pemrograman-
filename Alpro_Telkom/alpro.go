package alpro

import (
	"math"
)

//MATDIS ALPRO TELKOM UNIVERSITY
func Faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func Permutasi(n,r int) int {
	var hasil int
	hasil = Faktorial(n) / Faktorial(n - r)
	return hasil
}

func Combinasi(n,r int) int {
	var val int
	val = Faktorial(n) / (Faktorial(r) * Faktorial(n - r))
	return val
}

//SOAL 2!!
func Komposisi(f,g,h func(int)int) func(int) (int, int, int) {
    return func(x int) (int, int, int) {
        a := f(g(h(x)))
        b := g(h(f(x)))
        c := h(f(g(x)))
        
        return a, b, c
    }
}

func F(x int)int {return x * x}
func G(x int)int {return x - 2}
func H(x int)int {return x + 1}

//SOAL 3!!
func Lingkaran(x, y, cx1, cy1, cx2, cy2, r1, r2 int) string {
	d1 := (x-cx1)*(x-cx1) + (y-cy1)*(y-cy1)
	d2 := (x-cx2)*(x-cx2) + (y-cy2)*(y-cy2)

	r1sq := r1 * r1
	r2sq := r2 * r2

	if d1 <= r1sq && d2 <= r2sq {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if d1 <= r1sq {
		return "Titik di dalam lingkaran 1"
	} else {
		return "Titik di dalam lingkaran 2"
	}
}

//SOAL 4!!

func ZoomIn(w,h,s func(int)int)func(int)(int,int) {
	return func(i int) (int, int) {
		a := w(i)
		b := h(i)

		return a,b

	}
}

func ZoomOut(w,h,s func(int)int)func(int)(int,int) {
	return func(i int) (int, int) {
		a := w(i)
		b := h(i)

		return a,b

	}
}

//ZOOM IN FUNCTION

func W(w,i int)int {
	var re int

	re = w * i
	return re
}

func C(h,i int)int {
	var re int

	re = h * i
	return re
}

//ZOOM OUT FUNCTION

func A(w,i int)int {
	var re int

	re = int(math.Abs(float64(w / i)))
	return re
}

func E(h,i int)int {
	var re int

	re = int(math.Abs(float64(h / i)))
	return re
}

//SOAL 5!!
func Volume(t,r float64)float64 {
	pi := math.Floor(math.Pi * 100) / 100
	return  pi * math.Pow(r,2) * t
}

func Massa(rho,t,r float64)float64 {
	return  rho * Volume(t,r)
}
