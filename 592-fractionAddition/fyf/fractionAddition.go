package fyf

import (
	"fmt"
	"strconv"
)

func fractionAddition(expression string) string {
	if len(expression) == 0 {
		return ""
	}

	if expression[0] != '-' {
		expression = "+" + expression
	}

	l := len(expression)

	denominators := []int64{}
	moleculars := []string{}
	maxDenominator := int64(1)

	for i := 0; i < l; i++ {
		if expression[i] == '/' {
			for j := i + 1; j < l; j++ {
				if expression[j] == '-' || expression[j] == '+' || j == l-1 {
					end := j
					if j == l-1 {
						end = l
					}
					denominator, _ := strconv.ParseInt(expression[i+1:end], 10, 64)
					denominators = append(denominators, denominator)
					maxDenominator = maxDenominator * denominator
					i = j + 1
					break
				}
			}
		}
	}

	for i := 0; i < l; i++ {

		if expression[i] == '/' && i-2 >= 0 && (expression[i-2] == '+' || expression[i-2] == '-') {
			molecular := expression[i-2 : i]
			moleculars = append(moleculars, molecular)
		}
		if expression[i] == '/' && i-3 >= 0 && (expression[i-3] == '+' || expression[i-3] == '-') {
			molecular := expression[i-3 : i]
			moleculars = append(moleculars, molecular)
		}

	}

	temp := make([]int64, 0, len(moleculars))
	for i := 0; i < len(denominators); i++ {
		k := maxDenominator / denominators[i]
		temp = append(temp, k)
	}

	realMolecular := int64(0)

	for i := 0; i < len(moleculars); i++ {
		sign := int64(1)
		if moleculars[i][0] == '-' {
			sign = int64(-1)
		}
		k, _ := strconv.ParseInt(moleculars[i][1:len(moleculars[i])], 10, 64)
		mm := k * temp[i]
		realMolecular = realMolecular + mm*sign
	}

	if realMolecular == 0 {
		return "0/1"
	}
	if maxDenominator == realMolecular && realMolecular < 0 {
		return "-1/1"
	}

	if maxDenominator == realMolecular && realMolecular > 0 {
		return "1/1"
	}

	r := getMaxCommonDivisor(abs(maxDenominator), abs(realMolecular))

	if maxDenominator/r == 1 {
		return fmt.Sprintf("%d/1", realMolecular/r)
	}

	return fmt.Sprintf("%d/%d", realMolecular/r, maxDenominator/r)
}

func getMaxCommonDivisor(maxDenominator, realMolecular int64) int64 {
	big := int64(0)
	small := int64(0)
	if maxDenominator > realMolecular {
		big = maxDenominator
		small = realMolecular
	}
	if maxDenominator < realMolecular {
		big = realMolecular
		small = maxDenominator
	}
	if small == 0 {
		return 1
	}
	rest := big % small
	if rest == 0 {
		return small
	}
	return getMaxCommonDivisor(rest, small)
}

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}
