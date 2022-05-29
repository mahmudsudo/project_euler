# to get the sum of even numbers ina fibonacci sequence
def f(n):
    if n<=2 :
        return n
    
    return f(n-1) + f(n-2)
def getFibonacci(n):
    # list = [f(i) for i in range(n+1) if f(i) % 2 ==0]
    # # for i in range(n+1):
    # #     if f(i) % 2 ==0 :
    # #         list.append(f(i))
        
    # return list
    prev =1
    new=2
    res=0
    for i in range(n+1):
        then_add= new + prev
        if then_add % 2 == 0 :
            res += then_add
        prev = new
        new=then_add
    return res



print(getFibonacci(4000000))


