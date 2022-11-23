let mainArr = [
  { name: 'A', value: ['a', 'b', 'c'] },
  { name: 'B', value: [1, 2, 3, 4] },
  { name: 'C', value: ['x', 'y', 'z'] },
];

// expectedArr = [{info: [{key: A, value:a}, {key:B, value:1}], ...,value: a1x},a1y,a1z,a2x,...,c4z]

const { name: key, value } = mainArr.shift();
// mainArr = mainArr.slice(0); // [1,2,3,...] => [2,3,...]
let mutateArr = [...value]; // [a,b,c]

console.log('mainArr', mutateArr);
for (let mainIndex = 0; mainIndex < mainArr.length; mainIndex++) {
  const newArr: string[] = [];
  for (
    let mutateArrIndex = 0;
    mutateArrIndex < mutateArr.length;
    mutateArrIndex++
  ) {
    for (
      let valueIndex = 0;
      valueIndex < mainArr[mainIndex].value.length;
      valueIndex++
    ) {
      newArr.push(
        `${mutateArr[mutateArrIndex]}-${mainArr[mainIndex].value[valueIndex]}`,
      );
    }
  }
  mutateArr = newArr;
}

console.log(mutateArr);
