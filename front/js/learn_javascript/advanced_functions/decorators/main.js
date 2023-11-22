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

function spy(func) {
    return function (a, b) {
        let calls = [];
        func(a, b);
        calls.push(a,b);
    }
}

function work(a, b) {
    console.log(a + b);
}

work = spy(work);

work(1, 2);
work(3, 4);

for (let args of work.calls) {
    console.log(`call ${args.join(', ')}`);
}