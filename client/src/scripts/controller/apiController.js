import { API_KEY, API_URL } from '../config/config';

const getoptions = {
    method: 'GET',
    headers: {
        'x-api-key': API_KEY,
        'Content-Type': 'application/json',
    },
}

const generateorderoptions = (brothId, proteinId) => {
    return {
    method: 'POST',
    headers: {
        'x-api-key': API_KEY,
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        brothId: `${brothId}`, 
        proteinId: `${proteinId}`,
    }),
}};

const throwError = (response) => {
    throw new Error(response.status);
}

export function getBroths() {
    fetch(`${API_URL}/broths`, getoptions)
    .then(response => response.ok ? response.json() : throwError(response))
    .then(data => console.log(data))
    .catch(error => console.error('Erro:', error));
}

export function getProteins() {
    fetch(`${API_URL}/proteins`, getoptions)
    .then(response => response.ok ? response.json() : throwError(response))
    .catch(error => console.error('Erro:', error));
}

export function getRecipes() {
    fetch(`${API_URL}/recipes`, getoptions)
    .then(response => response.ok ? response.json() : throwError(response))
    .catch(error => console.error('Erro:', error));
}

export function generateOrder(brothId, proteinId) {
    fetch(`${API_URL}/orders`, generateorderoptions(brothId, proteinId))
    .then(response => response.ok ? response.json() : throwError(response))
    .then(data => console.log(data))
    .catch(error => console.error('Erro:', error));
}