function insertion_sort(vector) {
  for (let i = 0; i < vector.length; i++) {
    for (let j = i + 1; j > 0; j--) {
      if (vector[j] < vector[j - 1]) {
        let temp = vector[j - 1];
        vector[j - 1] = vector[j];
        vector[j] = temp;
      } else {
        break;
      }
    }
  }
  return vector;
}

console.log(insertion_sort([8, 4, 1, 3, 9, 5, 2, 7, 6]));