package main

import (
	"testing"
	"reflect"
)

func TestStandardScaler1(t *testing.T) {
	test1 := [][]float32{
		{0., 0.},
		{0., 0.},
	}
	test2 := [][]float32{
		{1., 1.},
		{1., 1.},
	}

	ref1 := [][]float32{
		{-1., -1.},
		{-1., -1.},
	}

	ref2 := [][]float32{
		{1., 1.},
		{1., 1},
	}

	scaler := InitStandardScaler(2)

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

func TestStandardScaler2(t *testing.T) {
	test1 := [][]float32{
		{0., 0., 44, 5},
		{0., 1., 3, 4},
	}
	test2 := [][]float32{
		{1., 1., 9, 3},
		{0., 1., 1, 1},
	}

	ref1 := [][]float32{
		{-0.577, -1.732, 1.707, 1.183},
		{-0.577, 0.577, -0.646, 0.507},
	}

	ref2 := [][]float32{
		{1.732, 0.577, -0.301, -0.169},
		{-0.577, 0.577, -0.760, -1.521},
	}

	scaler := InitStandardScaler(4)

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