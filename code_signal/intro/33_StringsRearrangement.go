package intro

var ans [][]string

func Permute(nums []string) [][]string {
	ans = make([][]string, 0)
	cp := make([]string, len(nums))
	copy(cp, nums)
	ans = append(ans, cp)

	if len(nums) < 2 {
		return ans
	}

	heap(nums, len(nums), len(nums))
	return ans
}

func heap(nums []string, n, l int) {

	for i := 0; ;i++ {

		if n > 2 {
			heap(nums, n - 1, l)
		}

		if i == n-1 {
			break
		}

		if n % 2 == 0 {
			swap(nums, i, n-1)
		} else {
			swap(nums, 0, n-1)
		}

		cp := make([]string, len(nums))
		copy(cp, nums)
		ans = append(ans, cp)
	}
}

func swap(strs []string, i, j int) {
	strs[i], strs[j] = strs[j], strs[i]
}



//
// I like this code, but it didn't work
//
func popElement(input []string, index int) (string, []string) {
	el := input[index]
	front := input[:index]
	back := input[index+1:]
	return el, append(front, back...)
}


func pushElement(input []string, element string, index int) []string {
	var all []string
	all = append(all, input[:index]...)
	all = append(all, element)
	all = append(all, input[index:]...)
	return all
}