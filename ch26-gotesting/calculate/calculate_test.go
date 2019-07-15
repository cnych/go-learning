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

func TestAdd(t *testing.T) {
	inputs:=[]struct{a,b,c int}{
		{1,2,3},
		{4,5,9},
		{10,20,30},
		{3001,4002,7003},
	}
	for _, data := range inputs {
		if result := Add(data.a, data.b); result != data.c {
			t.Errorf("Add(%d,%d) expected result=%d, actual result=%d\n",
				data.a, data.b, data.c, result)
		}
	}
}
