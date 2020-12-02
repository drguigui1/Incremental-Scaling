package main

import (
	"testing"
	"reflect"
)

func TestToStringSlices1(t *testing.T) {
	test1 := [][]float32{
		{1.2, 1.4, 1.53},
		{3.4, 4.5, 6.6},
		{1.1, 7.7, 4.4},
	}

	ref1 := [][]string{
		{"1.200000", "1.400000", "1.530000"},
		{"3.400000", "4.500000", "6.600000"},
		{"1.100000", "7.700000", "4.400000"},
	}

	res1 := ToStringSlices(&test1)

	if !reflect.DeepEqual(*res1, ref1) {
		t.Errorf("Error Test1")
		t.Errorf("res: %v\n", res1)
		t.Errorf("ref: %v\n", ref1)
	}
}

func TestToStringSlices2(t *testing.T) {
	test1 := [][]float32{
		{1.2, 1.4, 1.53, 5.5555},
		{3.4, 4.5, 6.6, 6.66},
		{1.1, 7.7, 4.4, 8.88},
	}

	ref1 := [][]string{
		{"1.200000", "1.400000", "1.530000", "5.555500"},
		{"3.400000", "4.500000", "6.600000", "6.660000"},
		{"1.100000", "7.700000", "4.400000", "8.880000"},
	}

	res1 := ToStringSlices(&test1)

	if !reflect.DeepEqual(*res1, ref1) {
		t.Errorf("Error Test1")
		t.Errorf("res: %v\n", res1)
		t.Errorf("ref: %v\n", ref1)
	}
}

func TestToFloatSlices(t *testing.T) {
	test1 := [][]string{
		{"1.123", "1.42", "5.643"},
		{"7.13", "2.4542", "5.6643"},
		{"9.23", "3.4254", "9.64653"},
	}

	ref1 := [][]float32{
		{1.123, 1.42, 5.643},
		{7.13, 2.4542, 5.6643},
		{9.23, 3.4254, 9.64653},
	}

	res1 := ToFloatSlices(&test1)

	if !reflect.DeepEqual(*res1, ref1) {
		t.Errorf("Error Test1")
		t.Errorf("res: %v\n", res1)
		t.Errorf("ref: %v\n", ref1)
	}

}

func TestRoundAllElm(t *testing.T) {
	test1 := [][]float32{
		{ 1.321112, 1.3423432, 1.4343},
		{ 1.321112, 1.3423432, 1.4343},
	}

	RoundAllElm(&test1)
	ref1 := [][]float32{
		{1.321, 1.342, 1.434},
		{1.321, 1.342, 1.434},
	}

	if !reflect.DeepEqual(test1, ref1) {
		t.Errorf("Error test1")
		t.Errorf("res: %v\n", test1)
		t.Errorf("ref: %v\n", ref1)
	}
}