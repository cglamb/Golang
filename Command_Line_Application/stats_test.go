package main

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []float64{3, 6, 0, 7, 4, 6, 3, 2, 1, 3}
	expected := 35.0
	result := sum(numbers)
	if result != expected {
		t.Errorf("Expected sum %f, but got %f", expected, result)
	}
}

func TestAvg(t *testing.T) {

	numbers := []float64{3, 6, 0, 7, 4, 6, 3, 2, 1, 3}
	expected := 3.5
	result := avg(numbers)
	if result != expected {
		t.Errorf("Expected sum %f, but got %f", expected, result)
	}
}

func TestCnt(t *testing.T) {

	numbers := []float64{3, 6, 0, 7, 4, 6, 3, 2, 1, 3}
	expected := float64(10)
	result := count(numbers)
	if result != expected {
		t.Errorf("Expected sum %f, but got %f", expected, result)
	}
}

func TestStd(t *testing.T) {

	numbers := []float64{3, 6, 0, 7, 4, 6, 3, 2, 1, 3}
	expected := 2.273030283
	result := std(numbers)
	if math.Abs(result-expected) > .0001 {
		t.Errorf("Expected sum %f, but got %f", expected, result)
	}
}

func TestMin(t *testing.T) {

	numbers := []float64{3, 6, 0, 7, 4, 6, 3, 2, 1, 3}
	expected := float64(0)
	result := min(numbers)
	if result != expected {
		t.Errorf("Expected sum %f, but got %f", expected, result)
	}
}

func TestMax(t *testing.T) {

	numbers := []float64{3, 6, 0, 7, 4, 6, 3, 2, 1, 3}
	expected := float64(7)
	result := max(numbers)
	if result != expected {
		t.Errorf("Expected sum %f, but got %f", expected, result)
	}
}

func TestPercentile(t *testing.T) {

	numbers := []float64{3, 6, 0, 7, 4, 6, 3, 2, 1, 3}

	expected1 := float64(2.25)
	expected2 := float64(3)
	expected3 := float64(5.5)

	result1 := Percentile(numbers, .25)
	result2 := Percentile(numbers, .50)
	result3 := Percentile(numbers, .75)

	if result1 != expected1 {
		t.Errorf("Expected sum %f, but got %f", expected1, result1)
	}

	if result2 != expected2 {
		t.Errorf("Expected sum %f, but got %f", expected2, result2)
	}

	if result3 != expected3 {
		t.Errorf("Expected sum %f, but got %f", expected3, result3)
	}
}
