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

const getFromApi = (path, options) => fetch(`${API_URL}/${path}`, options)
    .then(response => response.ok ? response.json() : throwError(response))
    .catch(error => console.error('Erro:', error));

export const getBroths = () => getFromApi('broths', getoptions);
export const getProteins = () => getFromApi('proteins', getoptions);
export const getRecipes = () => getFromApi('recipes', getoptions);
export const generateOrder = (brothId, proteinId) => getFromApi('orders', generateorderoptions(brothId, proteinId));