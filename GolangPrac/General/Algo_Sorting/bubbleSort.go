func bubbleSort(arr [int]) []int {
  swapped := true
  for swapped {
    swapped = false
    for ind := 0; i < len(arr) - 1 - ind; ind++ {
      if arr[ind] > arr[ind + 1] {
        arr[ind], arr[ind + 1] = arr[ind + 1], arr[ind]
        swapped = true
      }
    }
  }
  return arr
}
