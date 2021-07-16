package memoization

func SumToHundred(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if arr[0]+SumToHundred(arr[1:len(arr)-1]) > 100 {
		return SumToHundred(arr[1 : len(arr)-1])
	} else {
		return arr[0] + SumToHundred(arr[1:len(arr)-1])
	}
}
