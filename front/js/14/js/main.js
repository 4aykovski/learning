// let date = new Date();
// console.log(date.getFullYear())
// console.log(date.getMonth() + 1)
//
// let array = [5, 2, 3, 4];
// console.log(array.length);
// console.log(array.join(', '));
// console.log(array.sort());
//
// let string = array.join(', ');
// console.log(string.split(', '));

class Person {
    constructor(name, age, happiness) {
        this.name = name;
        this.age = age;
        this.happiness = happiness;
    }
}

let person = new Person('Egor', 14, 5);
console.log(person);