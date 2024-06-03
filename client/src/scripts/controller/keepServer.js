import { API_KEY, API_URL } from '../config/config';

const fivesec = 5 * 1000; // You might adjust intervals for testing
const fivemin = 5 * 60 * 1000;
const tenmin = 10 * 60 * 1000;

export function wakeUp() {
    setInterval(() => {
        fetch(`${API_URL}/`, {
            method: 'OPTIONS'
        })
        .catch(error => console.error('Erro (wakeup):', error));
    }, tenmin);
};

/*
* This is an implementation of a wakeup server feature. 
* Since Render spin down APIs with inactivity delaying requests up to 1 minute (or more)
* Since this is a showcase project, theres no way I'm paying, so i came up with this.
*/