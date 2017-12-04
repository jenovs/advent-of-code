// http://adventofcode.com/2017/day/4

const fs = require('fs');
const input = fs.readFileSync('./input.txt', 'utf8');

function countPhrases(str, noAnagramms) {
  const phrases = str.split('\n');
  phrases.pop();
  let counter = 0;
  phrases.forEach(phrase => {
    let hasDuplicate = false;
    const obj = phrase.split(' ').reduce((obj, val) => {
      if (noAnagramms) {
        val = val
          .split('')
          .sort()
          .join('');
      }
      if (obj[val]) {
        obj[val] = 2;
      } else {
        obj[val] = 1;
      }
      return obj;
    }, {});
    if (!Object.values(obj).includes(2)) counter++;
  });
  return counter;
}

console.log(countPhrases(input));
console.log(countPhrases(input, true));
