package main

// 1922. 统计好的数目

const mod = 1_000_000_007

func countGoodNumbers(n int64) int {
	return pow(5, (int(n)+1)/2) * pow(4, int(n)/2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

// 长度为n的整数， 偶数项包含0 总共有 (n+1) / 2 项
// 			     奇数项总共有 n / 2 项
// 偶数项要为偶数(0,2,4,6,8)五项 所以有 5^a
// 奇数项要为10以下的质数(2,3,5,7)四项 所以有 4^b
// 总数为两者的乘积 5^a * 4^b
// 因为数量过大所以需要取模计算幂乘
