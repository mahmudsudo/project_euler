package main


//sum of even fibonacci terms from 1,2,3,5,8........
func Problem2(n int){
	var re = getFibonacci(n)
	sum:=0
	for _,r := range re{
		
			sum+=r
		
	}
	println(sum)
}
func getFibonacci(n int) ([]int){
	var res  []int
	for i:=1;i<=n;i++{
		if f(i) % 2 == 0{
		res = append(res, f(i))
		}
	}
	return res
}
func f(n int) (int){
	if n <= 2{
		return n
	}
	return f(n-1) + f(n-2)
}