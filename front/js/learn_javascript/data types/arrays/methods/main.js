'use strict'

// function camelize(str) {
//     let array = str.split('-');
//     let result = [array[0]];
//     array = array.slice(1).map(item => item[0].toUpperCase() + item.slice(1));
//     return result.concat(array).join('');
// }
//
// let array = camelize("-list-style-image");
// console.log(array);


// function filterRange(arr, a, b) {
//     return arr.filter(num => a <= num && num <= b);
// }
//
// let arr = [5, 3, 8, 1];
// let filtered = filterRange(arr, 1, 4);
// console.log(arr);
// console.log(filtered);

// function filterRangeInPlace(arr, a, b) {
//     for (let i = 0; i < arr.length; i++) {
//         let num = arr[i]
//         if (a <= num && num <= b) {
//             arr.splice(i, 1);
//             i--;
//         }
//     }
// }
//
// let arr = [5, 3, 8, 1];
//
// filterRangeInPlace(arr, 1, 4);
//
// console.log(arr);

// let arr = [5, 2, 1, -10, 8];
//
// arr = arr.sort((a, b) => b - a);
//
// console.log(arr);

// function copySorted(arr) {
//
//     return arr.slice().sort();
// }
//
// let arr = ['HTML', 'JavaScript', 'CSS'];
//
// let sorted = copySorted(arr);
//
// console.log(arr);
// console.log(sorted);

// function Calculator() {
//
//     this.methods = {
//         '+': (a, b) => a + b,
//         '-': (a, b) => a - b,
//     }
//
//     this.calculate = function (str) {
//
//         let split = str.split(' '),
//             a = +split[0],
//             op = split[1],
//             b = +split[2];
//
//         if (!this.methods[op] || isNaN(a) || isNaN(b)) {
//             return NaN;
//         }
//
//         return this.methods[op](a, b);
//     }
//
//     this.addMethod = function (name, func) {
//         this.methods[name] = func;
//     }
// }
//
// let calc = new Calculator();
//
// console.log(calc.calculate('5 - 2'));

// let vasya = { name: "Вася", age: 25 };
// let petya = { name: "Петя", age: 30 };
// let masha = { name: "Маша", age: 28 };
//
// let users = [ vasya, petya, masha ];
//
// let names = users.map(user => user.name);
//
// console.log(names)

// let vasya = {name: "Вася", surname: "Пупкин", id: 1};
// let petya = {name: "Петя", surname: "Иванов", id: 2};
// let masha = {name: "Маша", surname: "Петрова", id: 3};
//
// let users = [vasya, petya, masha];
//
// let usersMapped = users.map(user => ({id: user.id, fullName: user.name + ' ' +user.surname}));
//
// console.log(usersMapped[0].id);
// console.log(usersMapped[0].fullName);

// function sortByAge(users) {
//     users.sort((user1, user2) => user1.age - user2.age);
// }
//
// let vasya = { name: "Вася", age: 25 };
// let petya = { name: "Петя", age: 30 };
// let masha = { name: "Маша", age: 28 };
//
// let arr = [ vasya, petya, masha ];
//
// sortByAge(arr);
//
// console.log(arr[1]);

// function getAverageAge(users) {
//     return users.reduce((sum, user) => sum + user.age, 0) / users.length;
// }
//
// let vasya = { name: "Вася", age: 25 };
// let petya = { name: "Петя", age: 30 };
// let masha = { name: "Маша", age: 29 };
//
// let arr = [ vasya, petya, masha ];
//
// console.log(getAverageAge(arr));

// function unique(arr) {
//     let result = []
//     for (let word of arr) {
//         if (!result.includes(word)) {
//             result.push(word);
//         }
//     }
//     return result;
// }
//
// let strings = ["кришна", "кришна", "харе", "харе",
//     "харе", "харе", "кришна", "кришна", ":-O"
// ];
//
// console.log(unique(strings));

function groupById(users) {
    return users.reduce((result, user) => {
        result[user.id] = user;
        return result;
    }, {});
}

let users = [
    {id: 'john', name: "John Smith", age: 20},
    {id: 'ann', name: "Ann Smith", age: 24},
    {id: 'pete', name: "Pete Peterson", age: 31},
];

let usersById = groupById(users);
console.log(usersById);