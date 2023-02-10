const userNameInput = document.querySelector( 'input[name="user"]' );
const passwordInput = document.querySelector('input[name="pass"]');
const inputBoxBtn = document.querySelector('.input-box-btn');
const inputBoxError = document.querySelector('.input-box-error');

if(
    userNameInput !== undefined &&
    passwordInput !== undefined &&
    inputBoxBtn !== undefined &&
    inputBoxError !== undefined
) {
    let userName = '';
    let password = '';

    userNameInput.addEventListener('input', ( event ) => {
        userName = event.target.value;
    });

    passwordInput.addEventListener( 'input', ( event ) => {
        password = event.target.value;
    });

    // click login btn event
    inputBoxBtn.addEventListener( 'click', async () => {
        inputBoxError.textContent = '';

        const request = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify( {
                'user': userName,
                'pass': password
            } ),
        }
        const res = await fetch( '/api/login', request );
        if ( res.status === 200 ) {
            // redirect to after login page.
            window.location.pathname = '/after_login.html';
        } else {
            inputBoxError.textContent = 'login error.';
        }
    } );
}
