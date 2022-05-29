package main

// sum of multiples of 3 and 5
func Problem1(n int){
	sum := 0
	for i:=1;i<n;i++{
		if i%3==0 || i%5==0{
			sum+=i
		}else{
			continue
		}
		
		
	}
	println(sum)
}
func hello(){
	println("hllp balle")
}