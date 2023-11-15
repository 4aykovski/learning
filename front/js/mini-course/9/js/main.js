// alert('Hello World!');

// let check = confirm('Hello World!');
// console.log(check);

// let age = Number(prompt("Hello World!", "20"));
// console.log(age);

let person = null;
if (confirm('Вы точно уверены?')) {
    person = prompt("Введите ваше имя");
    alert("Привет, " + person);
} else {
    alert('Пока!');
}
