let array = [1, 2, 3, 4];

console.log(array[1]);
console.log(array);

for (let i = 0; i < array.length; i++) {
    console.log(array[i]);
}

let matrix = [
    [1, 2],
    [3, 4],
    [5, 6]
];

console.log(matrix);

for (let i = 0; i < matrix.length; i++) {
    let res = '';
    for (let j = 0; j < matrix[i].length; j++) {
        res += String(matrix[i][j]) + ' ';
    }
    console.log(res);
}