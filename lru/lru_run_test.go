package lru

import (
	"reflect"
	"testing"
)

func TestRunLRU(t *testing.T) {
	type args struct {
		actions []string
		value   [][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{actions: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				value: [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			},
			want: []string{"null", "null", "null", "1", "null", "-1", "null", "-1", "3", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunLRU(tt.args.actions, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunLRU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunLRUV2(t *testing.T) {
	type args struct {
		actions []string
		value   [][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{actions: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				value: [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			},
			want: []string{"null", "null", "null", "1", "null", "-1", "null", "-1", "3", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunLRUV2(tt.args.actions, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunLRUV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
