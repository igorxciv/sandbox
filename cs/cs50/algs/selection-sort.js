function selection_sort(vector) {
  for (let i = 0, l = vector.length; i < l; i++) {
    let minIdx = i;
    for (let j = i + 1; j < l; j++) {
      if (vector[minIdx] > vector[j]) {
        minIdx = j;
      }
    }
    if (minIdx !== i) {
      let temp = vector[i];
      vector[i] = vector[minIdx];
      vector[minIdx] = temp;
    }
  }
  return vector;
}

console.log(selection_sort([8, 4, 1, 3, 9, 5, 2, 7, 6]));