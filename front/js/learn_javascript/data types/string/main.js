'use strict'
//
// function ucFirst(str) {
//     return str ? str[0].toLowerCase() + str.slice(1) : str;
// }
//
// alert(ucFirst('Вася'));


//
// function checkSpam(str) {
//     str = str.toLowerCase();
//     return str.includes('viagra') || str.includes('xxx');
// }
//
// alert(checkSpam('buy ViAgRA now'));
// alert(checkSpam('free xxxxx'));
// alert(checkSpam("innocent rabbit"));

//
// function truncate(str, maxLength) {
//     if (maxLength === undefined) return str;
//     return str.length <= maxLength ? str : str.slice(0, maxLength-3) + '...';
// }
//
// console.log(truncate('', 20));
// console.log(truncate('12345678901234567890', 5));
// console.log(truncate('1234567890123456789', 5));
// console.log(truncate('12345678901234567890123', 5));

function extractCurrencyValue(str) {
    return +str.slice(1);
}

console.log(extractCurrencyValue('$120') === 120);