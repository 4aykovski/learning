'use strict'
//
// function f() {
//     console.log('test');
// }
//
// Function.prototype.defer = function defer(delay) {
//     setTimeout(this, delay);
// }
//
// f.defer(1000);


function f(a, b) {
    console.log(a + b);
}

Function.prototype.defer = function (delay) {
    return (...args) => setTimeout(this, delay, ...args)
}

f.defer(1000)(1, 2, 3, 4);