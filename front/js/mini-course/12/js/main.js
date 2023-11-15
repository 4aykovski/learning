document.getElementById('main_form').addEventListener('submit', checkForm)

function checkForm(event) {
    event.preventDefault();

    let form = document.getElementById('main_form');

    let name = form.name.value;
    let pass = form.pass.value;
    let repass = form.repass.value;
    let sex = form.state.value;

    let error = '';
    if (name === '' || pass === '' || sex === '') {
        error = 'Fill all fields';
    } else if (name.length <= 1 || name.length > 50) {
        error = 'Fill correct name';
    } else if (pass !== repass) {
        error = 'Passwords must be similar';
    } else if (pass.split('&').length > 1) {
        error = 'Incorrect password';
    }

    if (error !== '') {
        document.getElementById('error').innerHTML = error;
    } else {
        alert('All data is correct');
    }

}