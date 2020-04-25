package strutil

import "testing"

func BenchmarkString2Intv1(b *testing.B) {
	for i := 0; i < 10000; i++ {
		String2Intv1(i)
	}
}

func BenchmarkString2Intv2(b *testing.B) {
	for i := 0; i < 10000; i++ {
		String2Intv2(i)
	}
}
