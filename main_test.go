package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, -1, 0, 1, 8, 1000} {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLength := size
			if wantLength < 0 {
				wantLength = 0
			}
			if len(data) != wantLength {
				t.Fatalf("len(generateRandomElements(%d)) = %d, want %d", size, len(data), wantLength)
			}
			for i, value := range data {
				if value <= 0 || value > maxRandomValue {
					t.Errorf("data[%d] = %d, want a value in [1, %d]", i, value, maxRandomValue)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{42}, want: 42},
		{name: "single negative element", data: []int{-42}, want: -42},
		{name: "maximum at start", data: []int{9, 3, 1}, want: 9},
		{name: "maximum in middle", data: []int{3, 9, 1}, want: 9},
		{name: "maximum at end", data: []int{1, 3, 9}, want: 9},
		{name: "equal elements", data: []int{5, 5, 5}, want: 5},
		{name: "zero elements", data: []int{0, 0, 0}, want: 0},
		{name: "negative elements", data: []int{-8, -2, -5}, want: -2},
		{name: "mixed signs", data: []int{-8, 0, 5, -2}, want: 5},
		{name: "integer limits", data: []int{minInt, 0, maxInt}, want: maxInt},
		{name: "minimum integer", data: []int{minInt}, want: minInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.data)
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
			if !slices.Equal(tt.data, original) {
				t.Error("maximum changed the input slice")
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{42}, want: 42},
		{name: "single negative element", data: []int{-42}, want: -42},
		{name: "short slice", data: []int{3, 9, 1}, want: 9},
		{name: "eight elements", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "maximum in remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, want: 99},
		{name: "negative elements", data: []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, want: -1},
		{name: "equal elements", data: []int{5, 5, 5, 5, 5, 5, 5, 5}, want: 5},
		{name: "zero elements", data: make([]int, CHUNKS), want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.data)
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
			if !slices.Equal(tt.data, original) {
				t.Error("maxChunks changed the input slice")
			}
		})
	}
}

func TestMaxChunksPartitions(t *testing.T) {
	for _, size := range []int{CHUNKS - 1, CHUNKS, CHUNKS + 1, 2*CHUNKS - 1, 2 * CHUNKS, 2*CHUNKS + 1, 8*CHUNKS + 1} {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			for maxIndex := 0; maxIndex < size; maxIndex++ {
				data := make([]int, size)
				for i := range data {
					data[i] = -2
				}
				data[maxIndex] = -1

				if got := maxChunks(data); got != -1 {
					t.Fatalf("maximum at index %d: maxChunks returned %d, want -1", maxIndex, got)
				}
			}
		})
	}
}
