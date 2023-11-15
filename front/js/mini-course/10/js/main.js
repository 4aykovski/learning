function info(word) {
    console.log(word);
}

function summa(array) {
    let sum = 0;

    for (let i = 0; i < array.length; i++) {
        sum += array[i];
    }

    return sum;
}


array = [5, 3, 5]
let sum = summa(array);
console.log(sum);