/* These are to be used if the server is down for some reason */

export const proteins = [
    {
        "id": 1,
        "imageInactive": "https://tech.redventures.com.br/icons/pork/inactive.svg",
        "imageActive": "https://tech.redventures.com.br/icons/pork/active.svg",
        "name": "Chasu",
        "description": "A sliced flavourful pork meat with a selection of season vegetables.",
        "price": 10
    },
    {
        "id": 2,
        "imageInactive": "https://tech.redventures.com.br/icons/yasai/inactive.svg",
        "imageActive": "https://tech.redventures.com.br/icons/yasai/active.svg",
        "name": "Yasai Vegetarian",
        "description": "A delicious vegetarian lamen with a selection of season vegetables.",
        "price": 10
    },
    {
        "id": 3,
        "imageInactive": "https://tech.redventures.com.br/icons/chicken/inactive.svg",
        "imageActive": "https://tech.redventures.com.br/icons/chicken/active.svg",
        "name": "Karaague",
        "description": "Three units of fried chicken, moyashi, ajitama egg and other vegetables.",
        "price": 0
    }
];

export const broths = [
    {
        "id": 1,
        "imageInactive": "https://tech.redventures.com.br/icons/salt/inactive.svg",
        "imageActive": "https://tech.redventures.com.br/icons/salt/active.svg",
        "name": "Salt",
        "description": "Simple like the seawater, nothing more",
        "price": 10
    },
    {
        "id": 2,
        "imageInactive": "https://tech.redventures.com.br/icons/miso/inactive.svg",
        "imageActive": "https://tech.redventures.com.br/icons/miso/active.svg",
        "name": "Miso",
        "description": "Paste made of fermented soybeans",
        "price": 12
    },
    {
        "id": 3,
        "imageInactive": "https://tech.redventures.com.br/icons/shoyu/inactive.svg",
        "imageActive": "https://tech.redventures.com.br/icons/shoyu/active.svg",
        "name": "Shoyu",
        "description": "The good old and traditional soy sauce",
        "price": 10
    }
];

export const recipes = [
    {
        "id": 1,
        "name": "Salt and Chasu Ramen",
        "description": "../../assets/dishes/chasu.png"
    },
    {
        "id": 2,
        "description": "Salt and Yasai Vegetarian Ramen",
        "image": "../../assets/dishes/yasai.png"
    },
    {
        "id": 3,
        "description": "Salt and Karaague Ramen",
        "image": "../../assets/dishes/karaague.png"
    },
    {
        "id": 4,
        "description": "Shoyu and Chasu Ramen",
        "image": "../../assets/dishes/chasu.png"
    },
    {
        "id": 5,
        "description": "Shoyu and Yasai Vegetarian Ramen",
        "image": "../../assets/dishes/yasai.png"
    },
    {
        "id": 6,
        "description": "Shoyu and Karaague Ramen",
        "image": "../../assets/dishes/karaague.png"
    },
    {
        "id": 7,
        "description": "Miso and Chasu Ramen",
        "image": "../../assets/dishes/chasu.png"
    },
    {
        "id": 8,
        "description": "Miso and Yasai Vegetarian Ramen",
        "image": "../../assets/dishes/yasai.png"
    },
    {
        "id": 9,
        "description": "Miso and Karaague Ramen",
        "image": "../../assets/dishes/karaague.png"
    }
];

export const findOrderRecipeId = (b, p) => {
    const brothMap = {"1": 0, "2": 3, "3": 6};
    const proteinMap = {"1": 1, "2": 2, "3": 3};
    
    let responseId = String(parseInt(brothMap[b]) + proteinMap[p]);

    return responseId;
};