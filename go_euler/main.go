package main




func main(){
Problem2(50)
		
}
func Problem2(n int){
	var re = getEvenFibonacci(n)
	sum:=0
	for _,r := range re{
			sum+=r
	}
	println(sum)
}
func getEvenFibonacci(n int) ([]int){
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

// func multiples_of_n(n int){
// 	for i:=1;i<=n;i++{
// 		if i % 3 !=0{
// 			continue
// 		}
// 	println(i)
// 	}
// }
// func fizzBuzz(n int) {
	
// 	for i:=1;i<=n;i++{
// 		var result = ""
// 		if i % 3 == 0 {
// 			result += "fizz"
// 		}
// 		 if i % 5 == 0{
// 			result += "buzz"
// 		}
// 		println(i,result)
// 	}
// }
// func getSmallestValue(x []int){
// 	var smallest = x[0]
// 	for _,val := range x{
// 		if smallest <= val{
// 			continue
// 		}
// 		smallest = val
// 	}
// 	println(smallest)
// }
// func getLargestValue(x ...int){
// 	var largest = x[0]
// 	for _,val := range x{
// 		if largest >= val{
// 			continue
// 		}
// 		largest = val
// 	}
// 	println(largest)
// }
// // func recFib(n int) (r int){
// // 	if n == 0{
// // 		r = 0
// // 	}else if n == 1 {
		
// // 	}
// // 	return
// // }
// func swap (x *int,y *int) (*int ,*int){
// 	v:=*x
// 	*x = *y
// 	*y = v
// 	return x, y
//}