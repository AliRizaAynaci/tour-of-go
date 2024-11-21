package main

func main() {
	n := 15
	memo := make([]int, n)
	println(Fib1(n))
	println(Fib2(n, memo))
	println(Fib3(n, make([]int, n+1)))
}

func Fib1(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)
}

func Fib2(n int, memo []int) int {
	if n <= 1 {
		return 1
	}
	if memo[n-1] != 0 {
		return memo[n-1]
	}
	res := Fib2(n-1, memo) + Fib2(n-2, memo)
	memo[n-1] = res
	return res
}

func Fib3(n int, k []int) int {
	if n <= 1 {
		return 1
	}
	k[0] = 1
	k[1] = 1

	for i := 2; i <= n; i++ {
		k[i] = k[i-1] + k[i-2]
	}

	return k[n]
}
