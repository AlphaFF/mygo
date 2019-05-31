package test

import (
	"testing"
)

func TestDivision1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没通过")
	} else {
		t.Log("第一个测试通过了")
	}
}

func TestDivision2(t *testing.T) {
	t.Error("就是不通过.")
}

func TestDivision3(t *testing.T) {
	if i, e := Division(6, 0); i == 0 || e != nil {
		t.Error("测试没通过.")
	} else {
		t.Log("测试通过了.")
	}
}
