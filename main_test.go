package main

import (
	"testing"
)

func TestStack_pop(t *testing.T) {
	type fields struct {
		elements []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test case 1",
			fields: fields{
				elements: []int{1, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				elements: tt.fields.elements,
			}
			if got := s.pop(); got != tt.want {
				t.Errorf("Stack.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_push(t *testing.T) {
	type fields struct {
		elements []int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				elements: tt.fields.elements,
			}
			s.push(tt.args.value)
		})
	}
}
