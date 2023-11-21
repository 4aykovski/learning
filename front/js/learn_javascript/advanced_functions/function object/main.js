'use strict'

// function makeCounter() {
//     let count = 0;
//
//     function counter() {
//         return count++;
//     }
//
//     counter.set = function set(a) {
//         return count=a;
//     }
//
//     counter.decrease = function decrease() {
//         return count--;
//     }
//
//     return counter;
// }
//
// let counter = makeCounter();
//
// console.log(counter()); // 0
// console.log(counter()); // 1
//
// counter.set(10); // установить новое значение счётчика
//
// console.log(counter()); // 10
//
// counter.decrease(); // уменьшить значение счётчика на 1
//
// console.log(counter()); // 10 (вместо 11)

function sum(a) {
    let currentSum = a;

    function func(b) {
        currentSum += b;
        return func;
    }

    func.toString = function () {
        return currentSum;
    }

    return func;
}

console.log(sum(1)(3)(5));


