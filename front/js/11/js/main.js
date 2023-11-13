let counter = 0;

function onClickButton(element) {
    counter++;
    element.innerHTML = 'Clicked ' + counter;


    element.style.cssText = "border-radius: 5px; border: 0; font-size:20px;"
}

function onInput(element) {
    if (element.value === 'Hello')
        alert('Hello!')
    console.log(element.value);
}