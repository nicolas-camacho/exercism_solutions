//
// This is only a SKELETON file for the 'Binary Search' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const find = (arr, n) => {
  let lower = 0
  let higher = arr.length
  let middle = Math.floor((lower + higher)/2)
  let element = arr[middle]

  if (element === n) {
    return middle
  }

  
  while (element !== n) {
    if(middle === higher) throw new Error('Value not in array')
    if (element > n) {
      higher = middle - 1
      lower = lower
      middle = Math.floor((lower + higher)/2)
      element = arr[middle]
    } else if(element < n) {
      lower = middle + 1
      higher = higher
      middle = Math.floor((lower + higher)/2)
      element = arr[middle]
    }
  }

  return middle
};
