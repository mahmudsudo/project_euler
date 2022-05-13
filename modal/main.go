package main



func main(){
	// multiples_of_n(500)
	// fizzBuzz(40)
	// x := []int{
	// 	48,96,86,68,
	// 	57,82,63,70,
	// 	37,34,83,27,
	// 	19,97, 9,17,
	// 	}
	// getSmallestValue(x)
	// getLargestValue(x...)

		// problem1(1000)
		problem2(400)
		
}
//sum of even fibonacci terms from 1,2,3,5,8........
func problem2(n int){
	var re = getFibonacci(n)
	sum:=0
	for _,r := range re{
		if r % 2 == 0{
			sum+=r
		}
	}
	println(sum)
}
func getFibonacci(n int) ([]int){
	var res  []int
	for i:=1;i<=n;i++{
		res = append(res, f(i))
		
	}
	return res
}
func f(n int) (int){
	if n <= 2{
		return n
	}
	return f(n-1) + f(n-2)
}
// sum of multiples of 3 and 5
// func problem1(n int){
// 	sum := 0
// 	for i:=1;i<n;i++{
// 		if i%3==0 || i%5==0{
// 			sum+=i
// 		}else{
// 			continue
// 		}
		
		
// 	}
// 	println(sum)
// }
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