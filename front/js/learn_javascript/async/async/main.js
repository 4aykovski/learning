// 'use strict'
//
// class HttpError extends Error {
//     constructor(response) {
//         super(`${response.status} for ${response.url}`);
//         this.name = 'HttpError';
//         this.response = response;
//     }
// }
//
// function loadJson(url) {
//     return fetch(url)
//         .then(response => {
//             if (response.status === 200) {
//                 return response.json();
//             } else {
//                 throw new HttpError(response);
//             }
//         })
// }
//
// // Запрашивать логин, пока github не вернёт существующего пользователя.
// function demoGithubUser() {
//     let name = prompt("Введите логин?", "iliakan");
//
//     return loadJson2(`https://api.github.com/users/${name}`)
//         .then(user => {
//             alert(`Полное имя: ${user.name}.`);
//             return user;
//         })
//         .catch(err => {
//             if (err instanceof HttpError && err.response.status === 404) {
//                 alert("Такого пользователя не существует, пожалуйста, повторите ввод.");
//                 return demoGithubUser();
//             } else {
//                 throw err;
//             }
//         });
// }
//
// async function loadJson2(url) {
//     let response = await fetch(url);
//     if (response.status === 200) {
//         return response.json()
//     }
//     throw new HttpError(response);
// }
//
// async function demoGithubUser2() {
//     let name = prompt("Введите логин?", "iliakan");
//
//     while (true) {
//         let user = await loadJson2(`https://api.github.com/users/${name}`);
//         try {
//             user = await loadJson(`https://api.github.com/users/${name}`);
//             break; // ошибок не было, выходим из цикла
//         } catch (err) {
//             if (err instanceof HttpError && err.response.status === 404) {
//                 // после alert начнётся новая итерация цикла
//                 alert("Такого пользователя не существует, пожалуйста, повторите ввод.");
//             } else {
//                 // неизвестная ошибка, пробрасываем её
//                 throw err;
//             }
//         }
//     }
//
//     alert(`Полное имя: ${user.name}.`);
//     return user;
//
// }
//
// demoGithubUser2();

async function wait() {
  await new Promise(resolve => setTimeout(resolve, 1000));

  return 10;
}

function f() {
    wait().then(r => alert(r))
}

f();