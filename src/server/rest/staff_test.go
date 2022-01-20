package rest

import (
	"kintai/mock"
	"testing"
)

func Test_GetStaffAll(t *testing.T) {
	ary := GetStaffAll()
	isEqual := len(ary) == len(mock.MultiStaff)
	if isEqual {
		for i := 0; i < len(ary); i++ {
			if ary[i] != mock.MultiStaff[i] {
				isEqual = false
				break
			}
		}
	}
	if !isEqual {
		t.Errorf("rest.GetStaffAll() failed;")
	} else {
		t.Log("rest.GetStaffAll() is ok")
	}
}
