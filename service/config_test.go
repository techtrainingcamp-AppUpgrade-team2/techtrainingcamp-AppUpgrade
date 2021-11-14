package service

import (
	"app_upgrade/model"
	"testing"
)

// @title             TestDeleteRule
// @description       删除规则单测
// @auth              刘晶玉         2021/11/14
func TestDeleteRule(t *testing.T) {
	cases := []struct {
		input  model.RuleIdType
		Expected bool
	}{
		//正常
		{
			input: model.RuleIdType(1),
			Expected:true,
		},
		//异常
		{
			input: model.RuleIdType(3),
			Expected:false,
		},
	}

	for _, c := range cases {

		if  res := DeleteRule(c.input);res != c.Expected{
			t.Errorf("expect is %+v , response is %+v", c.Expected, res)
		}
	}
}

// @title             TestPauseRule
// @description       暂停规则单测
// @auth              刘晶玉         2021/11/14
func TestPauseRule(t *testing.T) {
	cases := []struct {
		input  model.RuleIdType
		Expected bool
	}{
		//正常
		{
			input: model.RuleIdType(1),
			Expected:true,
		},
		//异常
		{
			input: model.RuleIdType(2),
			Expected:false,
		},
	}

	for _, c := range cases {
		if  PauseRule(c.input) != c.Expected{
			t.Errorf("expect is %+v , response is %+v", c.Expected, PauseRule(c.input))
		}
	}
}

