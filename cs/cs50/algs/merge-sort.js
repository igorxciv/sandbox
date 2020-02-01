function merge_sort(arr) {
  if (arr.length < 2) {
    return arr
  }
  const mid = Math.floor(arr.length / 2);
  const left = arr.slice(0, mid);
  const right = arr.slice(mid);
  return merge(merge_sort(left), merge_sort(right));
}

function merge(left, right) {
  const result = [];
  let leftIdx = 0, rightIdx = 0;

  while(leftIdx < left.length && rightIdx < right.length) {
    if (left[leftIdx] < right[rightIdx]) {
      result.push(left[leftIdx]);
      leftIdx++;
    } else {
      result.push(right[rightIdx]);
      rightIdx++;
    }
  }
  
  return result
    .concat(left.slice(leftIdx))
    .concat(right.slice(rightIdx));
}

console.log(merge_sort([8, 4, 7, 1, 9, 3, 5, 2]))