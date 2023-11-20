'use strict'
//
// function sumToRec(n) {
//     if (n === 1) {
//         return 1
//     }
//     return n + sumToRec(n - 1)
// }
//
// function sumToIter(n) {
//     let result = 0;
//     while (n > 0) {
//         result += n--;
//     }
//     return result;
// }
//
// function sumToProg(n) {
//     return ((1 + n) * n) / 2
// }
//
// console.log(sumToRec(100));
// console.log(sumToIter(100));
// console.log(sumToProg(100));
//
// function factorial(n) {
//     if (n===1) return 1;
//     return n*factorial(n-1);
// }
//
// console.log(factorial(4));


// function fib(n) {
//     let a = 1;
//     let b = 1;
//     for (let i = 3; i <= n; i++) {
//         let c = a + b;
//         a = b;
//         b = c;
//     }
//     return b;
// }
//
//
// console.log(fib(77));
//
// function printListRec(list) {
//     console.log(list.value);
//     if (list.next) printListRec(list.next);
// }
//
// function printListIter(list) {
//     let tmp = list;
//     while (tmp) {
//         console.log(tmp.value);
//         tmp = tmp.next;
//     }
// }
//
//
// let list = {
//     value: 1,
//     next: {
//         value: 2,
//         next: {
//             value: 3,
//             next: {
//                 value: 4,
//                 next: null
//             }
//         }
//     }
// };
//
// printListRec(list);
// printListIter(list);

function printListReverseRec(list) {
    if (list.next) printListReverseRec(list.next);
    console.log(list.value);
}

function printListIter(list) {
    let tmp = list;
    while (tmp) {
        console.log(tmp.value);
        tmp = tmp.next;
    }
}


let list = {
    value: 1,
    next: {
        value: 2,
        next: {
            value: 3,
            next: {
                value: 4,
                next: null
            }
        }
    }
};

printListReverseRec(list);