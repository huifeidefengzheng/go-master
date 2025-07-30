package dreamtask03

import "fmt"

func main() {

	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	str := fmt.Sprintf("%d", x)
	fmt.Println(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
