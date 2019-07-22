package calculate

import (
	"testing"
)

func TestEven(t *testing.T) {
	//t.Log("Start TestEvent")
	//t.Error("This is a T.error function")
	if !Even(10) {
		t.Log("10 must be even")
		t.Fail()
	}
	if Even(77) {
		t.Error("77 is not even")
	}
}

//func TestOdd(t *testing.T) {
//	if !Odd(11) {
//		t.Error("11 must be odd")
//	}
//	if Odd(70) {
//		t.Error("70 is not odd")
//	}
//}

func TestAdd(t *testing.T) {
	inputs := []struct{ a, b, c int }{
		{1, 2, 3},
		{4, 5, 9},
		{10, 20, 30},
		{3001, 4002, 7003},
	}
	for _, data := range inputs {
		if result := Add(data.a, data.b); result != data.c {
			t.Errorf("Add(%d,%d) expected result=%d, actual result=%d\n",
				data.a, data.b, data.c, result)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	input := struct{ a, b, c int }{
		3001, 4002, 7003,
	}
	for i := 0; i < b.N; i++ {
		if result := Add(input.a, input.b); result != input.c {
			b.Errorf("Add(%d,%d) expected result=%d, actual result=%d\n",
				input.a, input.b, input.c, result)
		}
	}
}
