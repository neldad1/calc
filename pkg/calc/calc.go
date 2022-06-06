package calc

import (
	"errors"
	"strconv"
)

type operand struct {
	num1 float64
	num2 float64
}

func Addition(args []string) (float64, error) {
	var sum float64
	for _, s := range args {
		i, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, errors.New("input not supported")
		}
		sum += i
	}
	return sum, nil
}

func Subtraction(arg1 string, arg2 string) (float64, error) {
	res, err := argsToOperand(arg1, arg2)
	if res == nil {
		return 0, err
	}
	return res.num1 - res.num2, nil
}

func Multiplication(arg1 string, arg2 string) (float64, error) {
	res, err := argsToOperand(arg1, arg2)
	if res == nil {
		return 0, err
	}
	return res.num1 * res.num2, nil
}

func Division(arg1 string, arg2 string) (float64, error) {
	res, err := argsToOperand(arg1, arg2)
	if res == nil {
		return 0, err
	}
	return res.num1 / res.num2, nil
}

func argsToOperand(arg1 string, arg2 string) (oper *operand, retErr error) {
	n1, err1 := strconv.ParseFloat(arg1, 64)
	n2, err2 := strconv.ParseFloat(arg2, 64)
	if err1 == nil && err2 == nil {
		oper = new(operand)
		oper.num1 = n1
		oper.num2 = n2
	} else {
		retErr = errors.New("input not supported")
	}
	return oper, retErr
}
