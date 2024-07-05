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

const getFromApi = async (path, options, b = 0, p = 0) => {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), 10000);

    try {
        
        const res = await fetch(`${API_URL}/${path}`, { ...options, signal: controller.signal });

        clearTimeout(timeoutId);

        if (!res.ok) throwError(res.status);

        return response.json();

    } catch (error) {

        return getPlaceholders(path, b, p);
    }
};

export const getBroths = () => getFromApi('broths', getoptions);
export const getProteins = () => getFromApi('proteins', getoptions);
export const getRecipes = () => getFromApi('recipes', getoptions);
export const generateOrder = (brothId, proteinId) => getFromApi('orders', generateorderoptions(brothId, proteinId), brothId, proteinId);