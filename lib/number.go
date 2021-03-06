package lib

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "AAA=int,int16,int32,int64,float32,float64"

import (
	"errors"
	"fmt"
	"strconv"
)

// SumAAA は、与えられた値の合計値を返します.
func SumAAA(values []AAA) AAA {
	var sum AAA = 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func MeanAAA(values []AAA) AAA {
	sum := SumAAA(values)
	return sum / AAA(len(values))
}

// FilterAAA は、与えられた値それぞれを関数へ適用し、結果がtrueになる値のSliceを返します.
func FilterAAA(values []AAA, f func(v AAA) bool) (newValues []AAA) {
	for _, value := range values {
		if f(value) {
			newValues = append(newValues, value)
		}
	}
	return
}

// FilterAAASlice は、与えられたSliceそれぞれを関数へ適用し、結果がtrueになるSliceのSliceを返します.
func FilterAAASlice(values [][]AAA, f func(v []AAA) bool) (newValues [][]AAA) {
	for _, value := range values {
		if f(value) {
			newValues = append(newValues, value)
		}
	}
	return
}

// UniqAAA は、与えられた値から重複を取り除いて返します.
func UniqAAA(values []AAA) (newValues []AAA) {
	m := map[AAA]bool{}
	for _, value := range values {
		m[value] = true
	}

	for key := range m {
		newValues = append(newValues, key)
	}
	return
}

// SubtractAAABy は、values1に関数を適用した値からvalues2に関数を適用した値を引いた結果を返します.
func SubtractAAABy(values1 []AAA, values2 []AAA, f func(v AAA) AAA) (newValues []AAA, err error) {
	if len(values1) != len(values2) {
		return nil, errors.New("two values lengths are different")
	}

	for i := 0; i < len(values1); i++ {
		fValue1 := f(values1[i])
		fValue2 := f(values2[i])
		newValues = append(newValues, fValue1-fValue2)
	}
	return newValues, nil
}

// SubtractAAA は、values1それぞれの要素からvalues2それぞれの要素を引いた結果を返します.
func SubtractAAA(values1 []AAA, values2 []AAA) (newValues []AAA, err error) {
	return SubtractAAABy(values1, values2, func(v AAA) AAA {
		return v
	})
}

// RDiffAAABy は、valuesそれぞれに関数を適用した結果の、隣り合う要素の差(n-(n-1))を返します.
func RDiffAAABy(values []AAA, f func(v AAA) AAA) (newValues []AAA, err error) {
	diffValues := append([]AAA{0}, values...)
	newValues, err = SubtractAAABy(values, diffValues[:len(diffValues)-1], f)
	if err != nil {
		return nil, fmt.Errorf("failed to RDiff: %v", err)
	}
	return newValues[1:], nil
}

// RDiffAAA は、隣り合う要素の差(n-(n-1))を返します.
func RDiffAAA(values []AAA) (newValues []AAA, err error) {
	return RDiffAAABy(values, func(v AAA) AAA {
		return v
	})
}

// StringToAAASlice は、Stringの各文字をAAAへ変換したSliceを返します.
func StringToAAASlice(s string) (ValueLine []AAA, err error) {
	for _, r := range s {
		v, err := strconv.ParseInt(string(r), 10, 64)
		if err != nil {
			return nil, err
		}
		ValueLine = append(ValueLine, AAA(v))
	}
	return
}

// StringSliceToAAASlice は、String sliceをAAA sliceへ変換します.
func StringSliceToAAASlice(line []string) (ValueLine []AAA, err error) {
	newLine, err := toSpecificBitIntLine(line, 64)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		ValueLine = append(ValueLine, AAA(v))
	}
	return
}

func NewAAAGridMap(grid [][]string, defaultValue AAA) (m [][]AAA) {
	for _, line := range grid {
		var newLine []AAA
		for range line {
			newLine = append(newLine, defaultValue)
		}
		m = append(m, newLine)
	}
	return
}

// AAARange は、startからendまで、stepずつ増加する数列を返します.(endは含まない)
func AAARange(start, end, step AAA) ([]AAA, error) {
	if end < start {
		return nil, fmt.Errorf("end(%v) is bigger than start(%v)", end, start)
	}
	s := make([]AAA, 0, int(1+(end-start)/step))
	for start < end {
		s = append(s, start)
		start += step
	}
	return s, nil
}

// AAAMap は、map[AAA][AAA]に便利メソッドを追加します.
type AAAMap map[AAA]AAA

func NewAAAMap(cap int) AAAMap {
	return make(map[AAA]AAA, cap)
}

// MustGetは、指定したkeyの値を返します. 指定したkeyの値が存在しない場合panicします.
func (m AAAMap) MustGet(key AAA) AAA {
	v, ok := m[key]
	if !ok {
		panic(fmt.Sprintf("ivnalid key is specfied in AAAMap: %v", key))
	}
	return v
}

// GetOr は、指定したkeyの値が存在すればその値を、存在しなければdefaultValueを返します.
func (m AAAMap) GetOr(key, defaultValue AAA) AAA {
	v, ok := m[key]
	if !ok {
		return defaultValue
	}
	return v
}

// ChMin は、与えられた値が既に存在する値よりも小さければ代入します.
// 指定したkeyの値が存在しない場合も代入します. この場合、2つめの戻り値はfalseになります.
func (m AAAMap) ChMin(key, value AAA, values ...AAA) (replaced bool, valueAlreadyExist bool) {
	min, _ := MinAAA(append(values, value)...)
	if v, ok := m[key]; ok {
		if v > min {
			m[key] = min
			return true, true
		} else {
			return false, true
		}
	}
	m[key] = min
	return true, false
}

// MustChMin は、与えられた値が既に存在する値よりも小さければ代入します.
// 指定したkeyの値が存在しない場合はpanicします.
func (m AAAMap) MustChMin(key, value AAA) (replaced bool) {
	v, ok := m[key]
	if !ok {
		panic(fmt.Sprintf("ivnalid key is specfied in AAAMap: %v", key))
	}
	if v > value {
		m[key] = value
		return true
	}
	return false
}

// ChMax は、与えられた値が既に存在する値よりも大きれば代入します.
// 指定したkeyの値が存在しない場合も代入します. この場合、2つめの戻り値はfalseになります.
func (m AAAMap) ChMax(key, value AAA, values ...AAA) (replaced bool, valueAlreadyExist bool) {
	max, _ := MaxAAA(append(values, value)...)
	if v, ok := m[key]; ok {
		if v < max {
			m[key] = max
			return true, true
		} else {
			return false, true
		}
	}
	m[key] = max
	return true, false
}

// MustChMin は、与えられた値が既に存在する値よりも大きければ代入します.
// 指定したkeyの値が存在しない場合はpanicします.
func (m AAAMap) MustChMax(key, value AAA) (replaced bool) {
	v, ok := m[key]
	if !ok {
		panic(fmt.Sprintf("ivnalid key is specfied in AAAMap: %v", key))
	}
	if v < value {
		m[key] = value
		return true
	}
	return false
}

func (m AAAMap) Keys() (keys []AAA) {
	keys = make([]AAA, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return
}

func (m AAAMap) Values() (values []AAA) {
	values = make([]AAA, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return
}

// AAAMap は、map[AAA][AAA]に便利メソッドを追加します.
type AAA2DMap map[AAA]AAAMap

var AAAMapCapForAAA2DMap = 0

func NewAAA2DMap(cap, cap2 int) AAA2DMap {
	AAAMapCapForAAA2DMap = cap2
	return make(map[AAA]AAAMap, cap)
}

func (m AAA2DMap) Get(key1, key2 AAA) (AAA, bool) {
	v1, ok := m[key1]
	if !ok {
		return 0, false
	}
	v2, ok := v1[key2]
	if !ok {
		return 0, false
	}
	return v2, true
}

func (m AAA2DMap) Set(key1, key2, value AAA) (isNewValue bool) {
	v1, ok := m[key1]
	if !ok {
		m[key1] = NewAAAMap(AAAMapCapForAAA2DMap)
		v1 = m[key1]
	}
	_, ok = v1[key2]
	v1[key2] = value
	return !ok
}

// MustGetは、指定したkeyの値を返します. 指定したkeyの値が存在しない場合panicします.
func (m AAA2DMap) MustGet(key1, key2 AAA) AAA {
	v1, ok := m[key1]
	if !ok {
		panic(fmt.Sprintf("ivnalid key1 is specfied in AAAMap: %v", key1))
	}
	v2, ok := v1[key2]
	if !ok {
		panic(fmt.Sprintf("ivnalid key2 is specfied in AAAMap: %v", key2))
	}
	return v2
}

// GetOr は、指定したkeyの値が存在すればその値を、存在しなければdefaultValueを返します.
func (m AAA2DMap) GetOr(key1, key2, defaultValue AAA) AAA {
	v, ok := m.Get(key1, key2)
	if !ok {
		return defaultValue
	}
	return v
}

// ChMin は、与えられた値が既に存在する値よりも小さければ代入します.
// 指定したkeyの値が存在しない場合も代入します. この場合、2つめの戻り値はfalseになります.
func (m AAA2DMap) ChMin(key1, key2, value AAA, values ...AAA) (replaced bool, valueAlreadyExist bool) {
	min, _ := MinAAA(append(values, value)...)
	if v, ok := m.Get(key1, key2); ok {
		if v > min {
			m.Set(key1, key2, min)
			return true, true
		} else {
			return false, true
		}
	}
	m.Set(key1, key2, min)
	return true, false
}

// ChMax は、与えられた値が既に存在する値よりも大きれば代入します.
// 指定したkeyの値が存在しない場合も代入します. この場合、2つめの戻り値はfalseになります.
func (m AAA2DMap) ChMax(key1, key2, value AAA, values ...AAA) (replaced bool, valueAlreadyExist bool) {
	max, _ := MaxAAA(append(values, value)...)
	if v, ok := m.Get(key1, key2); ok {
		if v < max {
			m.Set(key1, key2, max)
			return true, true
		} else {
			return false, true
		}
	}
	m.Set(key1, key2, max)
	return true, false
}

func (m AAA2DMap) GetMap(key AAA) (AAAMap, bool) {
	m1, ok := m[key]
	return m1, ok
}

func (m AAA2DMap) MustGetMap(key AAA) AAAMap {
	m1, ok := m.GetMap(key)
	if !ok {
		panic(fmt.Sprintf("invalid key is given to MustGetMap: %v", key))
	}
	return m1
}

type AAAList []AAA

func NewAAAList(length int, initialValue AAA) AAAList {
	ret := make([]AAA, length, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	return ret
}

func (a AAAList) ChMin(i int, value AAA) bool {
	curV := a[i]
	if curV > value {
		a[i] = value
		return true
	}
	return false
}

func (a AAAList) ChMax(i int, value AAA) bool {
	curV := a[i]
	if curV < value {
		a[i] = value
		return true
	}
	return false
}

type AAA2DList [][]AAA

func NewAAA2DList(length1, length2 int, initialValue AAA) AAA2DList {
	ret := make([][]AAA, length1, length1)
	for i := 0; i < length1; i++ {
		ret[i] = NewAAAList(length2, initialValue)
	}
	return ret
}

func (a AAA2DList) ChMin(i, j int, value AAA) bool {
	curV := a[i][j]
	if curV > value {
		a[i][j] = value
		return true
	}
	return false
}

func (a AAA2DList) ChMax(i, j int, value AAA) bool {
	curV := a[i][j]
	if curV < value {
		a[i][j] = value
		return true
	}
	return false
}
