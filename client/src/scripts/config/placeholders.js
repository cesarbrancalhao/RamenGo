/* These are to be used if the server is down for some reason */

const proteins = [
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

const broths = [
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

const recipes = [
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Salt and Chasu Ramen",
        "image": "../../assets/dishes/chasu.png"
    },
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Salt and Yasai Vegetarian Ramen",
        "image": "../../assets/dishes/yasai.png"
    },
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Salt and Karaague Ramen",
        "image": "../../assets/dishes/karaague.png"
    },
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Shoyu and Chasu Ramen",
        "image": "../../assets/dishes/chasu.png"
    },
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Shoyu and Yasai Vegetarian Ramen",
        "image": "../../assets/dishes/yasai.png"
    },
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Shoyu and Karaague Ramen",
        "image": "../../assets/dishes/karaague.png"
    },
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Miso and Chasu Ramen",
        "image": "../../assets/dishes/chasu.png"
    },
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Miso and Yasai Vegetarian Ramen",
        "image": "../../assets/dishes/yasai.png"
    },
    {
        "id": Math.floor(Math.random() * 100000),
        "description": "Miso and Karaague Ramen",
        "image": "../../assets/dishes/karaague.png"
    },
    {
        "id": Math.floor(Math.random() * 100000)
    }
];
const placeholders = {
    proteins: { items: proteins },
    broths: { items: broths },
    recipes: { items: recipes }
};

const findOrderRecipeId = (b, p) => {
    if (b === "0" || p === "0") return recipes[9];
    const brothMap = {"1": 0, "2": 3, "3": 6};
    const proteinMap = {"1": 1, "2": 2, "3": 3};
    
    let responseId = String(parseInt(brothMap[b]) + proteinMap[p]);

    return recipes[parseInt(responseId) - 1];
};

export const getPlaceHolderByName = (name, b = 0, p = 0) => {
    switch (name) {
        case "proteins":
            return placeholders.proteins.items;
        case "broths":
            return placeholders.broths.items;
        case "recipes":
            return placeholders.recipes.items;
        case "orders":
            return findOrderRecipeId(b, p);
    }
}