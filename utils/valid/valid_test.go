// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package valid_test

import (
	"strings"
	"testing"
	"wxapp/utils/valid"
)

func Test_Regex(t *testing.T) {
	rule := `regex:\d{6}|\D{6}|length:6,16`
	if m := valid.Check("123456", rule, nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("abcde6", rule, nil); m == nil {
		t.Error("校验失败")
	}
}

func Test_CheckMap(t *testing.T) {
	kvmap := map[string]string{
		"id":   "0",
		"name": "john",
	}
	rules := map[string]string{
		"id":   "required|between:1,100",
		"name": "required|length:6,16",
	}
	msgs := map[string]interface{}{
		"id": "ID不能为空|ID范围应当为:min到:max",
		"name": map[string]string{
			"required": "名称不能为空",
			"length":   "名称长度为:min到:max个字符",
		},
	}
	if m := valid.CheckMap(kvmap, rules, msgs); m == nil {
		t.Error("CheckMap校验失败")
	}

	kvmap = map[string]string{
		"id":   "1",
		"name": "john",
	}
	rules = map[string]string{
		"id":   "required|between:1,100",
		"name": "required|length:4,16",
	}
	msgs = map[string]interface{}{
		"id": "ID不能为空|ID范围应当为:min到:max",
		"name": map[string]string{
			"required": "名称不能为空",
			"length":   "名称长度为:min到:max个字符",
		},
	}
	if m := valid.CheckMap(kvmap, rules, msgs); m != nil {
		t.Error(m)
	}
}

func Test_Required(t *testing.T) {
	if m := valid.Check("1", "required", nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("", "required", nil); m == nil {
		t.Error(m)
	}
	if m := valid.Check("", "required-if:id,1,age,18", nil, map[string]string{"id": "1", "age": "19"}); m == nil {
		t.Error("Required校验失败")
	}
	if m := valid.Check("", "required-if:id,1,age,18", nil, map[string]string{"id": "2", "age": "19"}); m != nil {
		t.Error("Required校验失败")
	}
}

func Test_Ip(t *testing.T) {
	if m := valid.Check("10.0.0.1", "ipv4", nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("0.0.0.0", "ipv4", nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("1920.0.0.0", "ipv4", nil); m == nil {
		t.Error("ipv4校验失败")
	}
	if m := valid.Check("fe80::5484:7aff:fefe:9799", "ipv6", nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("fe80::5484:7aff:fefe:9799123", "ipv6", nil); m == nil {
		t.Error(m)
	}
}

func Test_Length(t *testing.T) {
	rule := "length:6,16"
	if m := valid.Check("123456", rule, nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("12345", rule, nil); m == nil {
		t.Error("长度校验失败")
	}
}

func Test_MinLength(t *testing.T) {
	rule := "min-length:6"
	if m := valid.Check("123456", rule, nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("12345", rule, nil); m == nil {
		t.Error("长度校验失败")
	}
}

func Test_MaxLength(t *testing.T) {
	rule := "max-length:6"
	if m := valid.Check("12345", rule, nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("1234567", rule, nil); m == nil {
		t.Error("长度校验失败")
	}
}

func Test_Between(t *testing.T) {
	rule := "between:6.01, 10.01"
	if m := valid.Check("10", rule, nil); m != nil {
		t.Error(m)
	}
	if m := valid.Check("10.02", rule, nil); m == nil {
		t.Error("大小范围校验失败")
	}
}

func Test_SetDefaultErrorMsgs(t *testing.T) {
	rule := "integer|length:6,16"
	msgs := map[string]string{
		"integer": "请输入一个整数",
		"length":  "参数长度不对啊老铁",
	}
	valid.SetDefaultErrorMsgs(msgs)
	m := valid.Check("6.66", rule, nil)
	if len(m) != 2 {
		t.Error("规则校验失败")
	} else {
		if v, ok := m["integer"]; ok {
			if strings.Compare(v, msgs["integer"]) != 0 {
				t.Error("错误信息不匹配")
			}
		}
		if v, ok := m["length"]; ok {
			if strings.Compare(v, msgs["length"]) != 0 {
				t.Error("错误信息不匹配")
			}
		}
	}
}

func Test_CustomError1(t *testing.T) {
	rule := "integer|length:6,16"
	msgs := map[string]string{
		"integer": "请输入一个整数",
		"length":  "参数长度不对啊老铁",
	}
	m := valid.Check("6.66", rule, msgs)
	if len(m) != 2 {
		t.Error("规则校验失败")
	} else {
		if v, ok := m["integer"]; ok {
			if strings.Compare(v, msgs["integer"]) != 0 {
				t.Error("错误信息不匹配")
			}
		}
		if v, ok := m["length"]; ok {
			if strings.Compare(v, msgs["length"]) != 0 {
				t.Error("错误信息不匹配")
			}
		}
	}
}

func Test_CustomError2(t *testing.T) {
	rule := "integer|length:6,16"
	msgs := "请输入一个整数|参数长度不对啊老铁"
	m := valid.Check("6.66", rule, msgs)
	if len(m) != 2 {
		t.Error("规则校验失败")
	} else {
		if v, ok := m["integer"]; ok {
			if strings.Compare(v, "请输入一个整数") != 0 {
				t.Error("错误信息不匹配")
			}
		}
		if v, ok := m["length"]; ok {
			if strings.Compare(v, "参数长度不对啊老铁") != 0 {
				t.Error("错误信息不匹配")
			}
		}
	}
}
