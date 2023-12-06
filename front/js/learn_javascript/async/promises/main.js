'use strict'

function delay(ms) {
    return new Promise(function(resolve, reject) {
        return setTimeout(resolve, ms);
    })
}

delay(3000).then(() => console.log('123'));