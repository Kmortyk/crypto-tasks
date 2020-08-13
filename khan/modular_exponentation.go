package khan

import (
	"math"
	"strconv"
)

//Fast Modular Exponentiation (by Cameron, kmortyk - go implementation)

//numToBinary converts a non-negative integer to binary represented by a string of 1s and 0s
func numToBinary(num int) string {
	binText := ""
	bits := math.Floor(math.Log(float64(num))/math.Log(2)) + 1

	curNum := float64(num)
	for i := float64(0); i < bits; i++ {
		binText = strconv.Itoa(int(curNum)%2) + binText
		curNum = math.Floor(curNum / 2)
	}
	return binText
}

//binaryToNum converts a binary number (represented by a string of 1s and 0s) to a non-negative integer
func binaryToNum(binText string) int {
	num, k := 0, 1
	for i := 0; i < len(binText); i++ {
		if binText[len(binText)-i-1] == '1' {
			num += 1 * k
		}
		k *= 2
	}
	return num
}

//mod calculates A mod B (using quotient remainder theorem)
//A=B*Q+R, where 0<=R<B - remainder
//A mod B = R
//R=A-B*Q, Q=floor(A/B)
func mod(a, b int) int {
	q := int(math.Floor(float64(a) / float64(b)))
	return a - b*q
}

//fastModExp calculates A^B mod C using fast modular exponentiation
func fastModExp(a, b, c int) int {
	bin := numToBinary(b)
	remainders := make([]int, len(bin))
	product := 0
	power := 1

	for i := 0; i < len(bin); i++ {
		if i == 0 {
			remainders[0] = mod(a, c)
		} else {
			remainders[i] = mod(remainders[i-1]*remainders[i-1], c)
		}

		if bin[len(bin)-1-i] == '1' {
			if product == 0 {
				product = remainders[i]
			} else {
				product *= remainders[i]
			}
			product = mod(product, c)
		}
		power *= 2
	}
	return mod(product, c)
}
