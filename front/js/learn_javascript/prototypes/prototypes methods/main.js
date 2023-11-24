'use strict'

let dict = Object.create(null, {
    toString: {
        value() {
            return Object.keys(this).join()
        }
    }
});

dict.apple = 'apple';
dict.__proto__ = '123';

for (let key in dict) {
    console.log(key);
}

alert(dict);