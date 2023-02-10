// logout input parts
const logoutBtn = document.querySelector( '.logout-btn' );

// click logout btn event
logoutBtn.addEventListener( 'click', async () => {
    const request = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify( {} ),
    }
    const res = await fetch( '/api/logout', request );
    if ( res.status === 200 ) {
        // redirect to login page.
        window.location.pathname = '/';
    }
} );
