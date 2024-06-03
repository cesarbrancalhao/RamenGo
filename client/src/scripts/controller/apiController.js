import { API_KEY, API_URL } from '../config/config';
import { getPlaceHolderByName } from '../config/placeholders';

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

const getPlaceholders = (name, b, p) => {
    return getPlaceHolderByName(name, b, p);
}

const throwError = (res) => {
    throw new Error(res.status);
}

const getFromApi = (path, options, b = 0, p = 0) => fetch(`${API_URL}/${path}`, options)
    .then(res => res.ok ? res.json() : throwError(res))
    .catch(() => getPlaceholders(path, b, p));

export const getBroths = () => getFromApi('broths', getoptions);
export const getProteins = () => getFromApi('proteins', getoptions);
export const getRecipes = () => getFromApi('recipes', getoptions);
export const generateOrder = (brothId, proteinId) => getFromApi('orders', generateorderoptions(brothId, proteinId), brothId, proteinId);