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
  let leftIdx = 0,
    rightIdx = 0,
    result = [];

  while(result.length !== left.length + right.length) {
    let elem = left[leftIdx];

    if (!elem || right[rightIdx] < left[leftIdx]) {
      elem = right[rightIdx];
      rightIdx++;
    } else {
      leftIdx++;
    }
    result.push(elem);
  }
  return result;
}

console.log(merge_sort([8, 4, 7, 1, 9, 3, 5, 2, 11]))