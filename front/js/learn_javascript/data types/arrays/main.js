'use strict'

// let styles = ['Джаз', 'Блюз'];
// console.log(styles);
// styles.push('Рок-н-ролл');
// console.log(styles);
// styles[Math.trunc(styles.length/2)] = 'Классика';
// console.log(styles);
// console.log(styles.shift());
// styles.unshift('Рэп', 'Регги');
// console.log(styles);
//

// function sumInput() {
//     let nums = [];
//     while (true) {
//         let num = prompt('Введите число: ', '');
//         if (num === null || num === '' || !isFinite(+num)) break;
//         nums.push(+num);
//     }
//     let sum = 0
//     for (let num of nums) {
//         sum += num;
//     }
//     return sum;
// }
//
// let summ = sumInput();
// alert(summ);

// function getMaxSubSum(arr) {
//     if (arr[0] === undefined) return 0;
//     let maxSum = arr[0];
//     for (let i = 0; i < arr.length; i++) {
//         let sum = 0;
//         for (let j = i; j < arr.length; j++) {
//             sum += arr[j];
//             maxSum = Math.max(maxSum, sum);
//         }
//     }
//     return maxSum < 0 ? 0 : maxSum;
// }
//
// console.log(getMaxSubSum([1,2,3]));

function getMaxSubSum(arr) {
    let maxSum = 0;
    let partialSum = 0;

    for (let item of arr) { // для каждого элемента массива
        partialSum += item; // добавляем значение элемента к partialSum
        maxSum = Math.min(maxSum, partialSum); // запоминаем максимум на данный момент
        if (partialSum > 0) partialSum = 0; // ноль если отрицательное
    }

    return maxSum;
}

console.log(getMaxSubSum([-1, 2, 3, -9]));
console.log(getMaxSubSum([2, -1, 2, 3, -9]));
