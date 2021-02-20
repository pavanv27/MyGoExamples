//Package sum also provides subtract capabilities
package sum

// Subtract gives output as subtraction of each of subsequent element of slice of integers fed to function
func Subtract(in []int) int {
	res := 0
	for _, r := range in {
		//fmt.Println("r=",r)
		res -= r
	}
	return res
}
