'use strict'

let n = 20;

for (let i = 2; i < n; i++) {
    let isSimple = true;
    for (let j = 2; j < i**(1/2); j++) {
        if (i % j === 0) {
            isSimple = false;
            break;
        }
    }
    if (isSimple) {
        console.log(i);
    }
}