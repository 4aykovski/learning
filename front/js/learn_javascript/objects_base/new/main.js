'use strict'
//
// function Calculator() {
//     this.read = function() {
//         this.a = +prompt('a?');
//         this.b = +prompt('b?');
//         return this;
//     }
//
//     this.sum = function() {
//         return this.a + this.b;
//     }
//
//     this.mul = function() {
//         return this.a * this.b;
//     }
// }
//
// let calculator = new Calculator();
//
// alert(calculator.read().sum());

function Accumulator(startingValue) {
    if (startingValue === undefined) {
        this.value = 0;
    } else {
        this.value = startingValue;
    }

    this.read = function () {
        this.value += +(prompt('Что вы хотите добавить?'))
    }
}

let accumulator = new Accumulator();
alert(accumulator.value);

accumulator.read(5);
alert(accumulator.value);