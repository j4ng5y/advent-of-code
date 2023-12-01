package main

import "testing"

func test_calculate_calibration_value(t *testing.T) {
	expected := 29
	i := calculate_calibration_value("two1nine")
	if i != expected {
		t.Fatalf("Expected %d, got %d", expected, i)
	}
}

func test_total_calculation(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := 281

	acc := 0
	for _, v := range input {
		acc += calculate_calibration_value(v)
	}
	if acc != expected {
		t.Fatalf("Expected %d, got %d", expected, acc)
	}
}
