'use strict'
//
// console.log( new Date('2012-02-20T03:12'))
// console.log( new Date(2012, 1, 20, 3, 12))

// function getWeekDay(date) {
//     let days = new Map([
//             [0, 'ВС'],
//             [1, 'ПН'],
//             [2, 'ВТ'],
//             [3, 'СР'],
//             [4, 'ЧТ'],
//             [5, 'ПТ'],
//             [6, 'СБ'],
//         ]
//     );
//     return days.get(date.getDay());
// }
//
// let date = new Date(2012, 0, 3);
// console.log(getWeekDay(date));

// function getDateAgo(date, days) {
//     let localeDate = new Date(date);
//     return new Date(localeDate.setDate(localeDate.getDate()-days));
// }
//
// let date = new Date(2015, 0, 2);
//
// console.log(date);
// console.log(getDateAgo(date, 1));
// console.log(getDateAgo(date, 2));
// console.log(getDateAgo(date, 365));

// function getLastDayOfMonth(year, month) {
//     let date = new Date(year, month + 1, 0);
//     return date.getDate();
// }
//
// console.log(getLastDayOfMonth(2012, 1))

// function getSecondsToday() {
//     let date = new Date();
//     let start_date = new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0)
//     return (date-start_date)/1000;
// }
//
// console.log(getSecondsToday())


// function getSecondsToTomorrow() {
//     let today = new Date();
//     let tomorrow = new Date(today.getFullYear(), today.getMonth(), today.getDate() + 1);
//     return (tomorrow - today) / 1000;
// }
//
// console.log(getSecondsToTomorrow())

function formatDate(date) {
    let now = new Date()
    if (now - date < 1000) {
        return 'right now';
    } else if (now - date < 60 * 1000) {
        return `${now.getSeconds() - date.getSeconds()} сек. назад`;
    } else if (now - date < 60 * 60 * 1000) {
        return `${now.getMinutes() - date.getMinutes()} мин. назад`;
    } else {
        return `${date.toLocaleDateString()}, ${date.toTimeString()}`.slice(0,20);
    }
}

console.log(formatDate(new Date(new Date - 1)));
console.log(formatDate(new Date(new Date - 30 * 1000)));
console.log(formatDate(new Date(new Date - 5 * 60 * 1000)));
console.log(formatDate(new Date(new Date - 86400 * 1000)));