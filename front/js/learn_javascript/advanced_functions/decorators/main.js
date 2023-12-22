'use strict'
//
// let workers = {
//     test() {
//         return '123';
//     },
//
//     func(phrase) {
//         console.log(`${phrase}` + this.test())
//     },
// }
//
// function decorator(func) {
//     return function (x) {
//         console.log(`Начало работы функции ${func.name}...`);
//         func.call(this, x);
//         console.log(`Конец работы функции ${func.name}!`);
//     }
// }
//
// workers.func = decorator(workers.func);
//
// workers.func('phrase');

// function spy(func) {
//     function wrapper(...args) {
//         wrapper.calls.push(args);
//         return func.apply(this, args);
//     }
//
//     wrapper.calls = [];
//     return wrapper;
// }
//
// function work(a, b) {
//     console.log(a + b);
// }
//
// work = spy(work);
//
// work(1, 2);
// work(3, 4);
//
// for (let args of work.calls) {
//     console.log(`call ${args.join(', ')}`);
// }

// function delay(func, delay) {
//     function wrapper(...args) {
//         setTimeout(() => func.apply(this, args), delay);
//     }
//
//     return wrapper;
// }
//
// function f(x) {
//     console.log(x);
// }
//
// let f1000 = delay(f, 1000);
// let f1500 = delay(f, 1500);
//
// f1000('text');
// f1500('text');

