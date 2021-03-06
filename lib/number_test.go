package lib

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSumValue(t *testing.T) {
	type args struct {
		values []AAA
	}
	tests := []struct {
		name string
		args args
		want AAA
	}{
		{
			name: "can calc sum",
			args: args{
				values: []AAA{0, 1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAAA(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAAA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterValue(t *testing.T) {
	type args struct {
		values []AAA
		f      func(v AAA) bool
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []AAA
	}{
		{
			name: "can filter",
			args: args{
				values: []AAA{1, 2, 3},
				f: func(v AAA) bool {
					return v == 2
				},
			},
			wantNewValues: []AAA{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := FilterAAA(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("FilterAAA() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestFilterAAASlice(t *testing.T) {
	type args struct {
		values [][]AAA
		f      func(v []AAA) bool
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues [][]AAA
	}{
		{
			name: "FilterAAASlice",
			args: args{
				values: [][]AAA{{1, 2, 3}, {4, 5, 6}},
				f: func(v []AAA) bool {
					return v[0] == 1
				},
			},
			wantNewValues: [][]AAA{{1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := FilterAAASlice(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("FilterAAASlice() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}
func TestUniqValue(t *testing.T) {
	type args struct {
		values []AAA
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []AAA
	}{
		{
			name: "uniq",
			args: args{
				values: []AAA{1, 2, 2, 3, 1},
			},
			wantNewValues: []AAA{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues := UniqAAA(tt.args.values)
			sort.SliceStable(gotNewValues, func(i, j int) bool {
				return gotNewValues[i] < gotNewValues[j]
			})
			sort.SliceStable(tt.wantNewValues, func(i, j int) bool {
				return tt.wantNewValues[i] < tt.wantNewValues[j]
			})
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("UniqAAA() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestSubtractValueBy(t *testing.T) {
	type args struct {
		values1 []AAA
		values2 []AAA
		f       func(v AAA) AAA
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []AAA
		wantErr       bool
	}{
		{
			name: "SubtractAAABy",
			args: args{
				values1: []AAA{4, 5, 6},
				values2: []AAA{3, 2, 1},
				f: func(v AAA) AAA {
					return v
				},
			},
			wantNewValues: []AAA{1, 3, 5},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues, err := SubtractAAABy(tt.args.values1, tt.args.values2, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubtractAAABy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("SubtractAAABy() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestRDiffValueBy(t *testing.T) {
	type args struct {
		values []AAA
		f      func(v AAA) AAA
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []AAA
		wantErr       bool
	}{
		{
			name: "RDiffAAABy",
			args: args{
				values: []AAA{1, 2, 4},
				f: func(v AAA) AAA {
					return v
				},
			},
			wantNewValues: []AAA{1, 2},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues, err := RDiffAAABy(tt.args.values, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("RDiffAAABy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("RDiffAAABy() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestStringToValueLine(t *testing.T) {
	type args struct {
		line []string
	}
	tests := []struct {
		name          string
		args          args
		wantValueLine []AAA
		wantErr       bool
	}{
		{
			name: "StringSliceToAAASlice",
			args: args{
				line: []string{"1", "2", "3"},
			},
			wantValueLine: []AAA{1, 2, 3},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueLine, err := StringSliceToAAASlice(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringtoValueLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValueLine, tt.wantValueLine) {
				t.Errorf("StringtoValueLine() = %v, want %v", gotValueLine, tt.wantValueLine)
			}
		})
	}
}

func TestStringToAAASlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name          string
		args          args
		wantValueLine []AAA
		wantErr       bool
	}{
		{
			name: "StringToAAASlice",
			args: args{
				s: "012",
			},
			wantValueLine: []AAA{0, 1, 2},
			wantErr:       false,
		},
		{
			name: "StringToAAASlice",
			args: args{
				s: "a",
			},
			wantValueLine: nil,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueLine, err := StringToAAASlice(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToAAASlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValueLine, tt.wantValueLine) {
				t.Errorf("StringToAAASlice() = %v, want %v", gotValueLine, tt.wantValueLine)
			}
		})
	}
}

func TestNewAAAGridMap(t *testing.T) {
	type args struct {
		grid         [][]string
		defaultValue AAA
	}
	tests := []struct {
		name  string
		args  args
		wantM [][]AAA
	}{
		{
			name: "NewAAAGridMap",
			args: args{
				grid:         [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
				defaultValue: 1,
			},
			wantM: [][]AAA{{1, 1, 1}, {1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := NewAAAGridMap(tt.args.grid, tt.args.defaultValue); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("NewAAAGridMap() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestAAARange(t *testing.T) {
	type args struct {
		start AAA
		end   AAA
		step  AAA
	}
	tests := []struct {
		name    string
		args    args
		want    []AAA
		wantErr bool
	}{
		{
			name: "AAARange",
			args: args{
				start: 0,
				end:   1,
				step:  1,
			},
			want: []AAA{0},
		},
		{
			name: "AAARange",
			args: args{
				start: 1,
				end:   3,
				step:  1,
			},
			want: []AAA{1, 2},
		},
		{
			name: "AAARange",
			args: args{
				start: 0,
				end:   1,
				step:  0.5,
			},
			want: []AAA{0, 0.5},
		},
		{
			name: "AAARange",
			args: args{
				start: 1,
				end:   0,
				step:  1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AAARange(tt.args.start, tt.args.end, tt.args.step)
			if (err != nil) != tt.wantErr {
				t.Errorf("AAARange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AAARange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAAAMap_MustGet(t *testing.T) {
	type args struct {
		key AAA
	}
	tests := []struct {
		name string
		m    AAAMap
		args args
		want AAA
	}{
		{
			m: map[AAA]AAA{
				0: 1,
				1: 2,
			},
			args: args{key: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.MustGet(tt.args.key); got != tt.want {
				t.Errorf("MustGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAAAMap_GetOr(t *testing.T) {
	type args struct {
		key          AAA
		defaultValue AAA
	}
	tests := []struct {
		name string
		m    AAAMap
		args args
		want AAA
	}{
		{
			m: map[AAA]AAA{
				0: 1,
				1: 2,
			},
			args: args{key: 0, defaultValue: 9},
			want: 1,
		},
		{
			m: map[AAA]AAA{
				0: 1,
				1: 2,
			},
			args: args{key: 3, defaultValue: 9},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.GetOr(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAAAMap_ChMin(t *testing.T) {
	type args struct {
		key         AAA
		value       AAA
		extraValues []AAA
	}
	tests := []struct {
		name                  string
		m                     AAAMap
		args                  args
		wantReplaced          bool
		wantValueAlreadyExist bool
		wantNewValue          AAA
	}{
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 1},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          1,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 1, extraValues: []AAA{2, 3}},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          1,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 2},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantNewValue:          2,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 3},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantNewValue:          2,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 3, extraValues: []AAA{1, 2, 4}},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          1,
		},
		{
			m:                     map[AAA]AAA{},
			args:                  args{key: 0, value: 1},
			wantReplaced:          true,
			wantValueAlreadyExist: false,
			wantNewValue:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReplaced, gotValueAlreadyExist := tt.m.ChMin(tt.args.key, tt.args.value, tt.args.extraValues...)
			if gotReplaced != tt.wantReplaced {
				t.Errorf("ChMin() gotReplaced = %v, want %v", gotReplaced, tt.wantReplaced)
			}
			if gotValueAlreadyExist != tt.wantValueAlreadyExist {
				t.Errorf("ChMin() gotValueAlreadyExist = %v, want %v", gotValueAlreadyExist, tt.wantValueAlreadyExist)
			}
			if tt.m[tt.args.key] != tt.wantNewValue {
				t.Errorf("ChMin() new value = %v, want %v", tt.m[tt.args.key], tt.wantNewValue)
			}
		})
	}
}

func TestAAAMap_MustChMin(t *testing.T) {
	type args struct {
		key   AAA
		value AAA
	}
	tests := []struct {
		name         string
		m            AAAMap
		args         args
		wantReplaced bool
		wantNewValue AAA
	}{
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:         args{key: 0, value: 1},
			wantReplaced: true,
			wantNewValue: 1,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:         args{key: 0, value: 2},
			wantReplaced: false,
			wantNewValue: 2,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:         args{key: 0, value: 3},
			wantReplaced: false,
			wantNewValue: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReplaced := tt.m.MustChMin(tt.args.key, tt.args.value); gotReplaced != tt.wantReplaced {
				t.Errorf("MustChMin() = %v, want %v", gotReplaced, tt.wantReplaced)
			}
			if tt.m[tt.args.key] != tt.wantNewValue {
				t.Errorf("ChMin() new value = %v, want %v", tt.m[tt.args.key], tt.wantNewValue)
			}
		})
	}
}

func TestAAAMap_ChMax(t *testing.T) {
	type args struct {
		key         AAA
		value       AAA
		extraValues []AAA
	}
	tests := []struct {
		name                  string
		m                     AAAMap
		args                  args
		wantReplaced          bool
		wantValueAlreadyExist bool
		wantNewValue          AAA
	}{
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 1},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantNewValue:          2,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 2},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantNewValue:          2,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 3},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          3,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 3, extraValues: []AAA{1, 2}},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          3,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:                  args{key: 0, value: 1, extraValues: []AAA{2, 3, 4}},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          4,
		},
		{
			m:                     map[AAA]AAA{},
			args:                  args{key: 0, value: 1},
			wantReplaced:          true,
			wantValueAlreadyExist: false,
			wantNewValue:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReplaced, gotValueAlreadyExist := tt.m.ChMax(tt.args.key, tt.args.value, tt.args.extraValues...)
			if gotReplaced != tt.wantReplaced {
				t.Errorf("ChMax() gotReplaced = %v, want %v", gotReplaced, tt.wantReplaced)
			}
			if gotValueAlreadyExist != tt.wantValueAlreadyExist {
				t.Errorf("ChMax() gotValueAlreadyExist = %v, want %v", gotValueAlreadyExist, tt.wantValueAlreadyExist)
			}
			if tt.m[tt.args.key] != tt.wantNewValue {
				t.Errorf("ChMin() new value = %v, want %v", tt.m[tt.args.key], tt.wantNewValue)
			}
		})
	}
}

func TestAAAMap_MustChMax(t *testing.T) {
	type args struct {
		key   AAA
		value AAA
	}
	tests := []struct {
		name         string
		m            AAAMap
		args         args
		wantReplaced bool
		wantNewValue AAA
	}{
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:         args{key: 0, value: 1},
			wantReplaced: false,
			wantNewValue: 2,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:         args{key: 0, value: 2},
			wantReplaced: false,
			wantNewValue: 2,
		},
		{
			m: map[AAA]AAA{
				0: 2,
			},
			args:         args{key: 0, value: 3},
			wantReplaced: true,
			wantNewValue: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReplaced := tt.m.MustChMax(tt.args.key, tt.args.value); gotReplaced != tt.wantReplaced {
				t.Errorf("MustChMax() = %v, want %v", gotReplaced, tt.wantReplaced)
			}
			if tt.m[tt.args.key] != tt.wantNewValue {
				t.Errorf("ChMin() new value = %v, want %v", tt.m[tt.args.key], tt.wantNewValue)
			}
		})
	}
}

func TestAAA2DMap_ChMin(t *testing.T) {
	type args struct {
		key1   AAA
		key2   AAA
		value  AAA
		values []AAA
	}
	tests := []struct {
		name                  string
		m                     AAA2DMap
		args                  args
		wantReplaced          bool
		wantValueAlreadyExist bool
		wantNewValue          AAA
	}{
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 1},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          1,
		},
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 1, values: []AAA{2, 3}},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          1,
		},
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 2},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantNewValue:          2,
		},
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 3},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantNewValue:          2,
		},
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 3, values: []AAA{1, 2, 4}},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          1,
		},
		{
			m:                     map[AAA]AAAMap{},
			args:                  args{key1: 1, key2: 1, value: 1},
			wantReplaced:          true,
			wantValueAlreadyExist: false,
			wantNewValue:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReplaced, gotValueAlreadyExist := tt.m.ChMin(tt.args.key1, tt.args.key2, tt.args.value, tt.args.values...)
			if gotReplaced != tt.wantReplaced {
				t.Errorf("ChMin() gotReplaced = %v, want %v", gotReplaced, tt.wantReplaced)
			}
			if gotValueAlreadyExist != tt.wantValueAlreadyExist {
				t.Errorf("ChMin() gotValueAlreadyExist = %v, want %v", gotValueAlreadyExist, tt.wantValueAlreadyExist)
			}
			if v := tt.m[tt.args.key1][tt.args.key2]; v != tt.wantNewValue {
				t.Errorf("ChMin() new value = %v, want %v", v, tt.wantNewValue)
			}
		})
	}
}

func TestAAA2DMap_ChMax(t *testing.T) {
	type args struct {
		key1   AAA
		key2   AAA
		value  AAA
		values []AAA
	}
	tests := []struct {
		name                  string
		m                     AAA2DMap
		args                  args
		wantReplaced          bool
		wantValueAlreadyExist bool
		wantNewValue          AAA
	}{
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 1},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantNewValue:          2,
		},
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 1, values: []AAA{2, 3}},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          3,
		},
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 2},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantNewValue:          2,
		},
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 3},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          3,
		},
		{
			m: map[AAA]AAAMap{
				1: {
					1: 2,
				},
			},
			args:                  args{key1: 1, key2: 1, value: 3, values: []AAA{1, 2}},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantNewValue:          3,
		},
		{
			m:                     map[AAA]AAAMap{},
			args:                  args{key1: 1, key2: 1, value: 1},
			wantReplaced:          true,
			wantValueAlreadyExist: false,
			wantNewValue:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReplaced, gotValueAlreadyExist := tt.m.ChMax(tt.args.key1, tt.args.key2, tt.args.value, tt.args.values...)
			if gotReplaced != tt.wantReplaced {
				t.Errorf("ChMax() gotReplaced = %v, want %v", gotReplaced, tt.wantReplaced)
			}
			if gotValueAlreadyExist != tt.wantValueAlreadyExist {
				t.Errorf("ChMax() gotValueAlreadyExist = %v, want %v", gotValueAlreadyExist, tt.wantValueAlreadyExist)
			}
		})
	}
}

func TestAAA2DList_ChMin(t *testing.T) {
	type args struct {
		i     int
		j     int
		value AAA
	}
	tests := []struct {
		name      string
		a         AAA2DList
		args      args
		want      bool
		wantValue AAA
	}{
		{
			a: AAA2DList{
				{1, 2, 3},
				{4, 5, 6},
			},
			args: args{
				i:     0,
				j:     0,
				value: -1,
			},
			want:      true,
			wantValue: -1,
		},
		{
			a: AAA2DList{
				{1, 2, 3},
				{4, 5, 6},
			},
			args: args{
				i:     1,
				j:     2,
				value: -1,
			},
			want:      true,
			wantValue: -1,
		},
		{
			a: AAA2DList{
				{1, 2, 3},
				{4, 5, 6},
			},
			args: args{
				i:     0,
				j:     0,
				value: 2,
			},
			want:      false,
			wantValue: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.ChMin(tt.args.i, tt.args.j, tt.args.value); got != tt.want {
				fmt.Println(tt.a[tt.args.i][tt.args.j])
				t.Errorf("ChMin() = %v, want %v", got, tt.want)
			}
			if gotValue := tt.a[tt.args.i][tt.args.j]; gotValue != tt.wantValue {
				t.Errorf("new value = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestAAA2DList_ChMax(t *testing.T) {
	type args struct {
		i     int
		j     int
		value AAA
	}
	tests := []struct {
		name      string
		a         AAA2DList
		args      args
		want      bool
		wantValue AAA
	}{
		{
			a: AAA2DList{
				{1, 2, 3},
				{4, 5, 6},
			},
			args: args{
				i:     0,
				j:     0,
				value: 9,
			},
			want:      true,
			wantValue: 9,
		},
		{
			a: AAA2DList{
				{1, 2, 3},
				{4, 5, 6},
			},
			args: args{
				i:     1,
				j:     2,
				value: 9,
			},
			want:      true,
			wantValue: 9,
		},
		{
			a: AAA2DList{
				{1, 2, 3},
				{4, 5, 6},
			},
			args: args{
				i:     0,
				j:     0,
				value: -1,
			},
			want:      false,
			wantValue: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.ChMax(tt.args.i, tt.args.j, tt.args.value); got != tt.want {
				fmt.Println(tt.a[tt.args.i][tt.args.j])
				t.Errorf("ChMax() = %v, want %v", got, tt.want)
			}
			if gotValue := tt.a[tt.args.i][tt.args.j]; gotValue != tt.wantValue {
				t.Errorf("new value = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
