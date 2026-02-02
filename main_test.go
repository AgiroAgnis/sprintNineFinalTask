package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"positive size", 10, 10},
		{"zero size", 0, 0},
		{"negative size", -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			require.Len(t, result, tt.want)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", nil, 0},
		{"one element", []int{42}, 42},
		{"all equal", []int{7, 7, 7}, 7},
		{"mixed", []int{-10, 0, 5, -3, 8}, 8},
		{"only negative", []int{-10, -20, -5}, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maximum(tt.data))
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", nil, 0},
		{"one element", []int{42}, 42},
		{"less than chunks", []int{3, 1}, 3},
		{"equal chunks", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"not divisible by chunks", []int{1, 5, 2, 9, 3}, 9},
		{"all equal", []int{7, 7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maxChunks(tt.data))
		})
	}
}

func TestLargeSlice(t *testing.T) {
	data := generateRandomElements(1_000_000)

	require.Equal(t, maximum(data), maxChunks(data))
}
