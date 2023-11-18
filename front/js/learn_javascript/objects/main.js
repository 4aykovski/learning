'use strict'

// task 1

// let user = {};
// console.log(user);
// user['name'] = 'John';
// console.log(user['name']);
// user['surname'] = 'Smith';
// console.log(user['surname']);
// user['name'] = 'Pete';
// console.log(user['name']);
// delete user['name'];
// console.log(user['name']);

// task 2
// function isEmpty(obj) {
//     for (let key in obj) {
//         if (key) return false;
//     }
//     return true;
// }
//
// let schedule = {};
//
// alert( isEmpty(schedule) );
//
// schedule.name = 'John';
//
// alert( isEmpty(schedule) );

// task 3
// function getSalariesSum(salaries) {
//     let summa = 0;
//     for (let key in salaries) {
//         summa += salaries[key];
//     }
//     return summa || 0;
// }
//
// let salaries = {
//     John: 100,
//     Ann: 160,
//     Pete: 130
// };
//
// let otherSalaries = {};
//
// let summaOne = getSalariesSum(salaries);
// let summaTwo = getSalariesSum(otherSalaries);
//
// alert(summaOne);
// alert(summaTwo);

// task 4
function multiplyNumeric(obj) {
    for (let key in obj) {
        if (typeof obj[key] === "number") {
            obj[key] *= 2;
        }
    }
}

let menu = {
    width: 200,
    height: 200,
    title: 'My menu',
};

console.log(menu)

multiplyNumeric(menu);

console.log(menu)
