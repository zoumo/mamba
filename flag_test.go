package mamba

import (
	"testing"
	time "time"
)

type testCase struct {
	flag     string
	defValue interface{}
	want     interface{}
}

func getTestCase(t string) testCase {
	switch t {
	case "Bool":
		return testCase{
			flag:     "true",
			defValue: false,
			want:     true,
		}
	case "BoolSlice":
		return testCase{
			flag:     "true,false",
			defValue: []bool{},
			want:     []bool{true, false},
		}
	case "Duration":
		return testCase{
			flag:     "1s",
			defValue: time.Duration(0),
			want:     time.Second,
		}
	case "Float32", "Float64":
		return testCase{
			flag:     "1.32",
			defValue: 1,
			want:     1.32,
		}
	case "Int", "Int16", "Int32", "Int64", "Uint", "Uint16", "Uint32", "Uint64":
		return testCase{
			flag:     "1",
			defValue: 0,
			want:     1,
		}
	case "IntSlice":
		return testCase{
			flag:     "1,2,3",
			defValue: []int{},
			want:     []int{1, 2, 3},
		}
	case "String":
		return testCase{
			flag:     "dev",
			defValue: "",
			want:     "dev",
		}
	case "StringSlice":
		return testCase{
			flag:     "for,dev",
			defValue: []string{},
			want:     []string{"for", "dev"},
		}
	default:
		return testCase{
			flag:     "",
			defValue: "",
			want:     "",
		}
	}
}

func Test_mergeWithEnvPrefix(t *testing.T) {

	type T struct {
		name string
		key  string
		want string
	}

	tests := []T{
		{"no prefix", "Key", "KEY"},
		{"no prefix", "KEY", "KEY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeWithEnvPrefix(tt.key); got != tt.want {
				t.Errorf("mergeWithEnvPrefix() = %v, want %v", got, tt.want)
			}
		})
	}

	envKeyReplacer = UnderlineReplacer
	tests = []T{
		{"with replacer", "Key-Key", "KEY_KEY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeWithEnvPrefix(tt.key); got != tt.want {
				t.Errorf("mergeWithEnvPrefix() = %v, want %v", got, tt.want)
			}
		})
	}

	SetEnvPrefix("test")
	tests = []T{
		{"with prefix and replacer", "Key-key", "TEST_KEY_KEY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeWithEnvPrefix(tt.key); got != tt.want {
				t.Errorf("mergeWithEnvPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
