/*
@Time : 2018/10/24 上午11:20
@Author : xiaoxuez
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
537 complex number multiplication
Given two strings representing two complex numbers.

You need to return a string representing their multiplication. Note i2 = -1 according to the definition.

Example 1:
Input: "1+1i", "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
Example 2:
Input: "1+-1i", "1+-1i"
Output: "0+-2i"
Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
Note:

The input strings will not have extra blank.
The input strings will be given in the form of a+bi, where the integer a and b will both belong to the range of [-100, 100]. And the output should be also in this form.

*/
func complexNumberMultiply(a string, b string) string {
	aarray := atoiComplex(a)
	barray := atoiComplex(b)
	x := aarray[0]*barray[0] - aarray[1]*barray[1]
	y := aarray[0]*barray[1] + aarray[1]*barray[0]
	return fmt.Sprintf("%d+%di", x, y)
}

func atoiComplex(a string) (aarray [2]int) {
	ac := strings.Split(a, "+")
	aarray[0], _ = strconv.Atoi(ac[0])
	aarray[1], _ = strconv.Atoi(ac[1][:len(ac[1])-1])
	return
}

func main() {
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}
