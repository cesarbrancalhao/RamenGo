import { API_KEY, API_URL } from '../config/config';

const twomin = 2 * 60 * 1000;
const fivesec = 5 * 1000;
const threemin = 3 * 60 * 1000;

const sleepy = (response) => {
    throw new Error(response.status);
}

export function wakeUp() {
    setInterval(() => {
        fetch(`${API_URL}/`, {
            method: 'GET',
            headers: {
                'x-api-key': API_KEY,
                'Content-Type': 'application/json',
            },
        })
        .then(response => response.ok ? response.json() : sleepy(response))
        .catch(error => console.error('Erro:', error));
    }, threemin);
};

/*
* This is an implementation of a wakeup server feature. 
* Since Render spin down APIs with inactivity delaying requests up to 1 minute (or more)
* Since this is a showcase project, theres no way I'm paying, so i came up with this.
*/