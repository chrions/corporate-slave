package fyf

import (
	"fmt"
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	ss := strings.Split(equation, "=")
	left := "+" + ss[0]
	right := "+" + ss[1]

	leftNum, leftX := getNumValue(left)
	rightNum, rightX := getNumValue(right)

	restX := leftX - rightX
	restNum := rightNum - leftNum
	if restX == 0 && restNum == 0 {
		return "Infinite solutions"
	}
	if restX == 0 && restNum != 0 {
		return "No solution"
	}
	if restX != 0 && restNum != 0 {
		i := restNum % restX
		if i == 0 {
			x := restNum / restX
			return fmt.Sprintf("x=%d", x)
		} else {
			return "No solution"
		}
	}
	return "x=0"
}

func getNumValue(ss string) (num int, x int) {
	simpleIndex := map[int]struct{}{}
	for i := 0; i < len(ss); i++ {
		if ss[i] == 'x' && (ss[i-1] == '+' || ss[i-1] == '-') {
			simpleIndex[i] = struct{}{}
		}
	}
	news := make([]byte, 0, len(ss)+len(simpleIndex))
	for i := len(ss) - 1; i >= 0; i-- {
		news = append(news, ss[i])
		if _, ok := simpleIndex[i]; ok {
			news = append(news, '1')
		}
	}
	for i, j := 0, len(news)-1; i < j; i, j = i+1, j-1 {
		news[i], news[j] = news[j], news[i]
	}
	ss = string(news)

	for i := 0; i < len(ss); i++ {
		if ss[i] != 'x' && ss[i] != '+' && ss[i] != '-' {
			if i == len(ss)-1 {
				num += stringToInt(ss[i-1:])
				continue
			}

			for j := i + 1; j < len(ss); j++ {
				if ss[j] == 'x' {
					x += stringToInt(ss[i-1 : j])
					i = j
					break
				}
				if ss[j] == '+' || ss[j] == '-' {
					num += stringToInt(ss[i-1 : j])
					i = j
					break
				}
				if j == len(ss)-1 {
					num += stringToInt(ss[i-1:])
					i = j
					break
				}

			}
		}
	}
	return
}

func stringToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
