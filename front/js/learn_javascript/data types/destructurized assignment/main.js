'use strict'

// let user = {
//   name: "John",
//   years: 30
// };
//
// let {name, years : age, isAdmin = false } = user;
//
// console.log(name);
// console.log(age);
// console.log(isAdmin);


function topSalary(salaries) {
    let max = {salary: 0, name: ''};
    for (let [name, salary] of Object.entries(salaries)) {
        if (max.salary < salary) {
            max.salary = salary;
            max.name = name;
        }
    }
    return max.name;
}

let salaries = {
    "John": 100,
    "Pete": 300,
    "Mary": 250
};

console.log(topSalary(salaries));