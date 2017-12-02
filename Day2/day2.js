const fs = require('fs');
const input = fs.readFileSync('./input.txt', 'utf8');

function getChecksum(str) {
  let sum = 0;
  str.split('\n').forEach(line => {
    let min, max;
    line.split('\t').forEach(c => {
      if (min === undefined || max === undefined) {
        min = +c;
        max = +c;
        return;
      }
      if (+c > max) max = +c;
      if (+c < min) min = +c;
    });
    sum += max - min;
  });
  return sum;
}

function evenlyDivisible(str) {
  let sum = 0;
  str.split('\n').forEach(line => {
    const arr = line.split('\t');
    for (let i = 0, l = arr.length; i < l; i++) {
      const num = +arr[i];
      for (let j = i + 1; j < l; j++) {
        if (+arr[i] % +arr[j] === 0) {
          sum += +arr[i] / +arr[j];
          break;
        } else if (+arr[j] % +arr[i] === 0) {
          sum += +arr[j] / +arr[i];
          break;
        }
      }
    }
  });
  return sum;
}

console.log(getChecksum(input));
console.log(evenlyDivisible(input));
