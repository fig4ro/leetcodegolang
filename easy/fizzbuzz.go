package easy

import "fmt"

func fizzBuzz(n int) []string {

	text := make([]string, n)

	for i := 1; i < n+1; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			text[i-1] = "FizzBuzz"
			continue
		}
		if i % 3 == 0 {
			text[i-1] = "Fizz"
			continue
		}
		if i % 5 == 0 {
			text[i-1] = "Buzz"
			continue
		}
		text[i-1] = fmt.Sprintf("%d", i)
	}
	
	return text
}

/* 
func multipleOf(i, k int) bool{
    return i % k == 0;
}

func fizzBuzz(n int) []string {

    result := make( []string, 0);
    
    for i := 1 ; i <= n ; i++{
        
        if( multipleOf(i, 3) && multipleOf(i, 5) ){
            result = append(result, "FizzBuzz")
            
        }else if( multipleOf(i, 3) ){
            result = append(result, "Fizz")
            
        }else if( multipleOf(i, 5) ){
            result = append(result, "Buzz")
            
        }else{
            result = append(result, strconv.Itoa(i) )
            
        }
    }
    
    return result
}
*/