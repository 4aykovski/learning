'use strict'

function checkAge1(age) {
    return age > 18 ? true : confirm('Родители разрешили?')
}

function checkAge2(age) {
    return age > 18 || confirm('Родители разрешили?')
}

function min(a, b) {
    return a > b ? b : a
}

function pow(x, n) {
    let res = x;
    while (n > 1) {
        res *= x;
        n--;
    }
    return res;
}

function ask(question, yes, no) {
    if (confirm(question)) yes()
    else no();
}

ask(
    'Вы согласны?',
    () => alert('Вы согласились'),
    () => alert('Вы отменили выполнение')
)
//
// console.log(checkAge1(20))
// console.log(checkAge2(20))
// console.log(min(5, 3))
// console.log(pow(5, 3))