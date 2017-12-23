// http://adventofcode.com/2017/day/3

const n = 289326;

const sqrt = Math.ceil(Math.sqrt(n));
const max = sqrt % 2 ? sqrt : sqrt + 1;
const min = max - 2;

const cells = max * max - min * min;
const sideLen = cells / 4;
const position = n - min * min;
const middle = sideLen / 2;
const offset = cells - position;
const depth = Math.floor(max / 2);
const steps = position % sideLen - middle;

console.log(depth + steps);
