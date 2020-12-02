package main

import (
	"testing"
	"reflect"
)

func TestMinMaxScaler1(t *testing.T) {
	test1 := [][]float32{
		{1,2,5.5},
		{6,84,423.43},
	}
	test2 := [][]float32{
		{3.3, 4.53, 6.42},
		{3.3, 4.53, 6.42},
		{3.3, 4.53, 6.42},
	}

	ref1 := [][]float32{
		{0., 0., 0.},
		{1., 1., 1.},
	}
	ref2 := [][]float32{
		{0.46, 0.03085366, 0.00220133},
    	{0.46, 0.03085366, 0.00220133},
    	{0.46, 0.03085366, 0.00220133},
	}


	scaler := InitMinMaxScaler(3)

	scaler.PartialFit(&test1)
	scaler.PartialFit(&test2)

	res1 := scaler.Transform(&test1)
	res2 := scaler.Transform(&test2)

	RoundAllElm(res1)
	RoundAllElm(res2)
	RoundAllElm(&ref1)
	RoundAllElm(&ref2)

	if !reflect.DeepEqual(*res1, ref1) {
		t.Errorf("Error Test1")
		t.Errorf("res: %v\n", *res1)
		t.Errorf("ref: %v\n", ref1)
	}

	if !reflect.DeepEqual(*res2, ref2) {
		t.Errorf("Error Test1")
		t.Errorf("res: %v\n", *res2)
		t.Errorf("ref: %v\n", ref2)
	}
}


func TestMinMaxScaler2(t *testing.T) {
	test1 := [][]float32{
		{1, 0, 0, 5.5},
		{6, 84, 0.43, 7.7},
	}

	test2 := [][]float32{
		{3.3, 4.53, 6.42, 0},
		{3.3, 0, 6.42, 0},
	}

	test3 := [][]float32{
		{3.3, 4.53, 6.42, 1},
		{3.3, 4.53, 6.42, 1},
	}

	ref1 := [][]float32{
		{0., 0., 0., 0.71428571},
		{1., 1., 0.06697819, 1.},
	}
	ref2 := [][]float32{
		{0.46, 0.05392857, 1., 0.},
    	{0.46, 0., 1., 0.},
	}

	ref3 := [][]float32{
		{0.46, 0.05392857, 1., 0.12987013},
		{0.46, 0.05392857, 1., 0.12987013},
	}

	scaler := InitMinMaxScaler(4)

	scaler.PartialFit(&test1)
	scaler.PartialFit(&test2)
	scaler.PartialFit(&test3)

	res1 := scaler.Transform(&test1)
	res2 := scaler.Transform(&test2)
	res3 := scaler.Transform(&test3)

	RoundAllElm(res1)
	RoundAllElm(res2)
	RoundAllElm(res3)
	RoundAllElm(&ref1)
	RoundAllElm(&ref2)
	RoundAllElm(&ref3)

	if !reflect.DeepEqual(*res1, ref1) {
		t.Errorf("Error Test2")
		t.Errorf("res: %v\n", *res1)
		t.Errorf("ref: %v\n", ref1)
	}

	if !reflect.DeepEqual(*res2, ref2) {
		t.Errorf("Error Test2")
		t.Errorf("res: %v\n", *res2)
		t.Errorf("ref: %v\n", ref2)
	}

	if !reflect.DeepEqual(*res3, ref3) {
		t.Errorf("Error Test2")
		t.Errorf("res: %v\n", *res3)
		t.Errorf("ref: %v\n", ref3)
	}
}