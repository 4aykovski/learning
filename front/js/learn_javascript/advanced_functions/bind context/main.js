'use strict'
//
// function sayHi() {
//     console.log(this.name);
// }
// sayHi.text = 5;
//
// let bound = sayHi.bind({ name: 'Вася' });
//
// console.log(bound);

function askPassword(ok, fail) {
  let password = prompt("Password?", '');
  if (password === "rockstar") ok();
  else fail();
}

let user = {
  name: 'Вася',

  loginOk() {
    alert(`${this.name} logged in`);
  },

  loginFail() {
    alert(`${this.name} failed to log in`);
  },

};

askPassword(user.loginOk.bind(user), user.loginFail.bind(user));