package rest

import (
	"kintai/mock"
	"testing"
)

func Test_GetStaffAll(t *testing.T) {
	ary := GetStaffAll()
	isEqual := len(ary) == len(mock.MultiStaff)+1
	if isEqual {
		if ary[0] != mock.OneStaff {
			isEqual = false
		}
	}
	if isEqual {
		for i := 0; i < len(mock.MultiStaff); i++ {
			if ary[i+1] != mock.MultiStaff[i] {
				isEqual = false
				break
			}
		}
	}
	if !isEqual {
		t.Errorf("get=%v, org=%v", ary, mock.MultiStaff)
		t.Errorf("rest.GetStaffAll() failed;")
	} else {
		t.Log("rest.GetStaffAll() is ok")
	}
}
