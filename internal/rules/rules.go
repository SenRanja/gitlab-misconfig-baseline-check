package rules

import (
	"errors"
	"gitlab-misconfig/internal/log"
	"strconv"
)

type Rule struct {
	// 描述信息
	Description string

	// 规则编号
	RuleID string

	// 键值
	Keywords string
}

func CheckRule(realValue string, operation string, ruleValue string) (bool, error) {
	switch operation {
	case ">":
		realValueNumber := stringToInt(realValue)
		ruleVulueNumber := stringToInt(ruleValue)
		if realValueNumber > ruleVulueNumber {
			return true, nil
		} else {
			return false, nil
		}
	case ">=":
		realValueNumber := stringToInt(realValue)
		ruleVulueNumber := stringToInt(ruleValue)
		if realValueNumber >= ruleVulueNumber {
			return true, nil
		} else {
			return false, nil
		}
	case "<":
		realValueNumber := stringToInt(realValue)
		ruleVulueNumber := stringToInt(ruleValue)
		if realValueNumber < ruleVulueNumber {
			return true, nil
		} else {
			return false, nil
		}
	case "<=":
		realValueNumber := stringToInt(realValue)
		ruleVulueNumber := stringToInt(ruleValue)
		if realValueNumber <= ruleVulueNumber {
			return true, nil
		} else {
			return false, nil
		}
	case "=":
		if ruleValue == realValue {
			return true, nil
		} else {
			return false, nil
		}
	case "!=":
		if ruleValue != realValue {
			return true, nil
		} else {
			return false, nil
		}
	default:
		err := errors.New("check rule error")
		return false, err
	}

}

func stringToInt(value string) int {
	valueNumber, err := strconv.Atoi(value)
	if err != nil {
		log.Error("change value error")
		log.Error(err)
	}
	return valueNumber
}
