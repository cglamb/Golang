package main

import (
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {

	data, _ := readCSV("testdata/housesInput.csv")
	result := data[0]
	expected := []string{"value", "income", "age", "rooms", "bedrooms", "pop", "hh"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result %v is not equal to expected %v", result, expected)
	}

}

func TestReadHeaders(t *testing.T) {

	data, _ := readCSV("testdata/housesInput.csv")
	result := readHeaders(data)

	expected := []string{"value", "income", "age", "rooms", "bedrooms", "pop", "hh"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result %v is not equal to expected %v", result, expected)
	}

}

func TestReadData(t *testing.T) {

	data, _ := readCSV("testdata/housesInput.csv")
	result_all := readData(data)
	result := result_all[0]

	expected := []string{"452600", "8.3252", "41", "880", "129", "322", "126"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result %v is not equal to expected %v", result, expected)
	}

}

func TestReadCol(t *testing.T) {

	numbers := [][]float64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}
	expected := []float64{1, 6, 11, 16}

	result := readCol(numbers, 0)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result %v is not equal to expected %v", result, expected)
	}

}
