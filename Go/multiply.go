/*
字符串相乘

给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:
输入: num1 = "2", num2 = "3"
输出: "6"

示例 2:
输入: num1 = "123", num2 = "456"
输出: "56088"

说明：
num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

package leetcode

import (
	"bytes"
	"strconv"
)

/**
num1的第i位(高位从0开始)和num2的第j位相乘的结果在乘积中的位置是[i+j, i+j+1]
例: 123 * 45,  123的第1位 2 和45的第0位 4 乘积 08 存放在结果的第[1, 2]位中
  index:    0 1 2 3 4

                1 2 3
            *     4 5
            ---------
                  1 5
                1 0
              0 5
            ---------
              0 6 1 5
                1 2
              0 8
            0 4
            ---------
            0 5 5 3 5
这样我们就可以单独都对每一位进行相乘计算把结果存入相应的index中
**/
func multiply(num1 string, num2 string) string {
	lenNum1 := len(num1)
	lenNum2 := len(num2)
	mul := make([]int, lenNum1+lenNum2)

	var bitMul int
	for i := lenNum1 - 1; i >= 0; i-- {
		for j := lenNum2 - 1; j >= 0; j-- {
			bitMul = int(num1[i]-'0') * int(num2[j]-'0')
			bitMul += mul[i+j+1]
			mul[i+j+1] = bitMul % 10
			mul[i+j] += bitMul / 10
		}
	}

	var ret bytes.Buffer
	start := 0
	for start = range mul {
		if mul[start] != 0 {
			break
		}
	}
	for _, ch := range mul[start:] {
		ret.WriteString(strconv.Itoa(ch))
	}

	return ret.String()
}
