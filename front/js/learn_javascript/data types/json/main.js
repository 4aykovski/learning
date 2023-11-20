'use strict'
//
// let user = {
//   name: "Василий Иванович",
//   age: 35
// };
//
// console.log(user);
// let userJson = JSON.stringify(user);
// console.log(userJson);
// let userParse = JSON.parse(userJson);
// console.log(userParse);

let room = {
    number: 23
};

let meetup = {
    title: "Совещание",
    occupiedBy: [{name: "Иванов"}, {surname: "Петров"}],
    place: room,
};

room.occupiedBy = meetup;
meetup.self = meetup;

console.log(JSON.stringify(meetup, function replacer(key, value) {
    return (key !== "" && value === meetup) ? undefined : value;
}))

