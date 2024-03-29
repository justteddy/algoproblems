package replace_with_greatest

/*
   Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.
   After doing so, return the array.

   Example 1:
   Input: arr = [17,18,5,4,6,1]
   Output: [18,6,6,6,1,-1]
*/

func replaceElements(arr []int) []int {
	max := 0
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = max
		if tmp > max {
			max = tmp
		}
	}

	arr[len(arr)-1] = -1
	return arr
}
