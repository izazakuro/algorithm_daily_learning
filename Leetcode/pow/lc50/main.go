package main

// 50. pow(x,n)

func myPow(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n > 0 {
		if n&1 > 0 {
			res *= x
		}
		x *= x
		n >>= 1
	}

	return res

}

// n为次数，将n分解为二进制，例如13 = 1101 -> 通过位运算向右shift查看n的最低位是否为1
// 为1： 结果 *= x
// 没一轮 x *= x
