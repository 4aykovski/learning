'use strict'

let calculator = {
    read() {
        this.a = +prompt('Введите первое число: ');
        this.b = +prompt('Введите второе число: ');
    },

    sum() {
        return this.a + this.b;
    },

    mul() {
        return this.a * this.b;
    },
}

calculator.read();
let summa = calculator.sum();
let multi = calculator.mul();
alert(summa);
alert(multi);

