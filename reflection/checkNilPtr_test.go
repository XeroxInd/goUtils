package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckNilPtr(t *testing.T) {
	type nestedStruct struct {
		datePtr *string
	}

	type bigStruct struct {
		structPtr *bigStruct
	}

	type structWithPointersFail struct {
		data          string
		strPtrMap     map[string]*string
		strPtrArray   [3]*string
		strPtrSlice   []*string
		interfaceWrap interface{}
		dataStruct    nestedStruct
		dataStructPtr *nestedStruct
		strPtr        *string
		bigStruct     bigStruct
	}

	type structWithPointersSuccess struct {
		data          string
		strPtrMap     map[string]*string
		strPtrArray   [3]*string
		strPtrSlice   []*string
		interfaceWrap interface{}
		dataStruct    nestedStruct
		dataStructPtr *nestedStruct
		strPtr        *string
	}

	stringPtr := new(string)
	fail := &structWithPointersFail{
		"hello",
		map[string]*string{"stringPtr": stringPtr},
		[3]*string{stringPtr, stringPtr, stringPtr},
		[]*string{stringPtr, stringPtr},
		nestedStruct{stringPtr},
		nestedStruct{stringPtr},
		&nestedStruct{stringPtr},
		stringPtr,
		bigStruct{&bigStruct{&bigStruct{&bigStruct{&bigStruct{&bigStruct{nil}}}}}}}
	success := &structWithPointersSuccess{
		"hello",
		map[string]*string{"stringPtr": stringPtr},
		[3]*string{stringPtr, stringPtr, stringPtr},
		[]*string{stringPtr, stringPtr},
		nestedStruct{stringPtr},
		nestedStruct{stringPtr},
		&nestedStruct{stringPtr},
		stringPtr}

	hasNil, fieldName := CheckNilPtr(reflect.ValueOf(fail), "fail")

	fmt.Printf("Struct contain nil field : [%t] at [%s] \n", hasNil, fieldName)
	if !hasNil || fieldName != "structPtr" {
		t.Errorf("CheckNilPtr expected: [true] at [strPtr] got: [%t] at [%v]",
			hasNil, fieldName)
	}
	hasNil, fieldName = CheckNilPtr(reflect.ValueOf(success), "fail")

	fmt.Printf("Struct contain nil field : [%t] at [%s] \n", hasNil, fieldName)
	if hasNil || fieldName != "" {
		t.Errorf("CheckNilPtr expected: [false] at [] got: [%t] at [%v]",
			hasNil, fieldName)
	}
}
