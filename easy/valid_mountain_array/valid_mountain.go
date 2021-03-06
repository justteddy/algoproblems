package valid_mountain

/*
	Given an array of integers arr, return true if and only if it is a valid mountain array.
	Recall that arr is a mountain array if and only if:
    arr.length >= 3
    There exists some i with 0 < i < arr.length - 1 such that:
        arr[0] < arr[1] < ... < arr[i - 1] < A[i]
        arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

	Example 1:
	Input: arr = [2,1]
	Output: false

	Example 2:
	Input: arr = [3,5,5]
	Output: false

	Example 3:
	Input: arr = [0,3,2,1]
	Output: true
*/

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i, j := 0, len(arr)-1
	for i != j {
		if arr[j-1] > arr[j] {
			j--
			continue
		}
		if arr[i+1] > arr[i] {
			i++
			continue
		}
		return false
	}
	if i == 0 || j == len(arr)-1 {
		return false
	}
	return true
}

func validMountainArray2(arr []int) bool {
	length := len(arr)
	if length < 3 {
		return false
	}

	i := 0
	for i+1 < length && arr[i] < arr[i+1] {
		i++
	}

	if i == length-1 || i == 0 {
		return false
	}

	for i+1 < length && arr[i] > arr[i+1] {
		i++
	}

	return i == length-1
}
