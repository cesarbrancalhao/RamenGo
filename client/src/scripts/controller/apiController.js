import { API_KEY, API_URL } from '../config/getEnv';

export function getBroths() {
    fetch(`${API_URL}/broths`, {
        method: 'GET',
        headers: {
            'x-api-key': API_KEY,
            'Content-Type': 'application/json',
        },
    })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
}