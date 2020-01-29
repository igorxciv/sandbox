function bubble_sort(vector) {
  for (let i = 0, l = vector.length; i < l; i+=1) {
    let sorted = true;
    for (let j = 1; j < l; j+=1) {
      if (vector[j] < vector[j - 1]) {
        let temp = vector[j - 1];
        vector[j - 1] = vector[j];
        vector[j] = temp;
        sorted = false;
      }
    }
    /**
     * on the first iteration array already sorted
     */
    if (sorted) break;
  }
  return vector;
}

console.log([1,2,3,4,5,6]);
console.log(bubble_sort([8, 4, 7, 1, 3, 9]));
