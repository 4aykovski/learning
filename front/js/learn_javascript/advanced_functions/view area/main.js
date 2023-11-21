'use strict'
// function Counter() {
//   this.count = 0;
//
//   this.up = function() {
//     return ++this.count;
//   };
//   this.down = function() {
//     return --this.count;
//   };
// }
//
// let counter = new Counter();
//
// console.log(counter.up());
// console.log(counter.up());
// console.log(counter.up());
// console.log(counter.count);

// function sum(a) {
//     function sum2(b) {
//         return a+b;
//     }
//
//     return sum2(2);
// }
//
// console.log(sum(5));

// function inBetween(a, b) {
//     return function (value) {
//         return a <= value && value <= b;
//     };
// }
//
// function inArray(arr) {
//     return function (value) {
//         return arr.includes(value);
//     };
// }
//
//
// let array = [1, 2, 3, 4, 5, 6];
//
// console.log(array.filter(inBetween(3, 6)));
// console.log(array.filter(inArray([1, 2, 3])));

// function byField(field) {
//     return (a,b) => a[field] > b[field] ? 1 : -1;
// }
//
// let users = [
//   { name: "John", age: 20, surname: "Johnson" },
//   { name: "Pete", age: 18, surname: "Peterson" },
//   { name: "Ann", age: 19, surname: "Hathaway" }
// ];
//
// console.log(users.sort(byField('name')));
// console.log(users.sort(byField('age')));

function makeArmy() {
    let shooters = [];

    for (let i = 0; i < 10; i++) {
        let shooter = function () { // функция shooter
            console.log(i); // должна выводить порядковый номер
        };
        shooters.push(shooter); // и добавлять стрелка в массив
    }

    // ...а в конце вернуть массив из всех стрелков
    return shooters;
}

let army = makeArmy();

// все стрелки выводят 10 вместо их порядковых номеров (0, 1, 2, 3...)
army[0](); // 10 от стрелка с порядковым номером 0
army[1](); // 10 от стрелка с порядковым номером 1
army[2](); // 10 ...и т.д.