'use strict'

// function readNumber() {
//     let num;
//     do {
//         num = prompt('Введите число', 0);
//     } while (!isFinite(num));
//
//     if (num === null || num === '') return null;
//
//     return +num;
// }
//
// alert(`Число: ${readNumber()}`)

// function random(min, max) {
//     return min + Math.random() * (max - min);
// }


function randomInteger(min, max) {
    return Math.floor(min + Math.random() * (max + 1 - min));
}

alert(randomInteger(1, 5));
alert(randomInteger(1, 5));
alert(randomInteger(1, 5));
alert(randomInteger(1, 5));
alert(randomInteger(1, 5));