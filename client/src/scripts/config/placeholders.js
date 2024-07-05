/* These are to be used if the server is down for some reason */

const proteins = {
    "1": {
        "id": 1,
        "imageInactive": "./src/assets/icons/inactive/inactive-pork.svg",
        "imageActive": "./src/assets/icons/active/active-pork.svg",
        "name": "Chasu",
        "description": "A sliced flavourful pork meat with a selection of season vegetables.",
        "price": 10
    },
    "2": {
        "id": 2,
        "imageInactive": "./src/assets/icons/inactive/inactive-yasai.svg",
        "imageActive": "./src/assets/icons/active/active-yasai.svg",
        "name": "Yasai Vegetarian",
        "description": "A delicious vegetarian lamen with a selection of season vegetables.",
        "price": 10
    },
    "3": {
        "id": 3,
        "imageInactive": "./src/assets/icons/inactive/inactive-chicken.svg",
        "imageActive": "./src/assets/icons/active/active-chicken.svg",
        "name": "Karaague",
        "description": "Three units of fried chicken, moyashi, ajitama egg and other vegetables.",
        "price": 12
    }
};

const broths = {
    "1": {
        "id": 1,
        "imageInactive": "./src/assets/icons/inactive/inactive-salt.svg",
        "imageActive": "./src/assets/icons/active/active-salt.svg",
        "name": "Salt",
        "description": "Simple like the seawater, nothing more",
        "price": 10
    },
    "2": {
        "id": 2,
        "imageInactive": "./src/assets/icons/inactive/inactive-miso.svg",
        "imageActive": "./src/assets/icons/active/active-miso.svg",
        "name": "Miso",
        "description": "Paste made of fermented soybeans",
        "price": 12
    },
    "3": {
        "id": 3,
        "imageInactive": "./src/assets/icons/inactive/inactive-shoyu.svg",
        "imageActive": "./src/assets/icons/active/active-shoyu.svg",
        "name": "Shoyu",
        "description": "The good old and traditional soy sauce",
        "price": 10
    }
};

const recipes = {
    "0": {
        "id": Math.floor(Math.random() * 100000),
        "description": "Salt and Chasu Ramen",
        "image": "./src/assets/dishes/chasu.png"
    },
    "1": {
        "id": Math.floor(Math.random() * 100000),
        "description": "Salt and Yasai Vegetarian Ramen",
        "image": "./src/assets/dishes/yasai.png"
    },
    "2": {
        "id": Math.floor(Math.random() * 100000),
        "description": "Salt and Karaague Ramen",
        "image": "./src/assets/dishes/karaague.png"
    },
    "3": {
        "id": Math.floor(Math.random() * 100000),
        "description": "Miso and Chasu Ramen",
        "image": "./src/assets/dishes/chasu.png"
    },
    "4": {
        "id": Math.floor(Math.random() * 100000),
        "description": "Miso and Yasai Vegetarian Ramen",
        "image": "./src/assets/dishes/yasai.png"
    },
    "5": {
        "id": Math.floor(Math.random() * 100000),
        "description": "Miso and Karaague Ramen",
        "image": "./src/assets/dishes/karaague.png"
    },
    "6": {
        "id": Math.floor(Math.random() * 100000),
        "description": "Shoyu and Chasu Ramen",
        "image": "./src/assets/dishes/chasu.png"
    },
    "7": {
        "id": Math.floor(Math.random) * 100000,
        "description": "Shoyu and Yasai Vegetarian Ramen",
        "image": "./src/assets/dishes/yasai.png"
    },
    "8": {
        "id": Math.floor(Math.random()) * 100000,
        "description": "Shoyu and Karaague Ramen",
        "image": "./src/assets/dishes/karaague.png"
    }
};

const findOrderRecipeId = (b, p) => {
    if (b === "0" || p === "0") return recipes[9];
    const brothMap = {"1": 0, "2": 3, "3": 6};
    const proteinMap = {"1": 1, "2": 2, "3": 3};
    
    let responseId = parseInt(brothMap[b]) + parseInt(proteinMap[p]);
    
    return recipes[responseId - 1];
};

export const getPlaceHolderByName = (name, b = 0, p = 0) => {
    switch (name) {
        case "proteins":
            return proteins;
        case "broths":
            return broths;
        case "recipes":
            return recipes;
        case "orders":
            return findOrderRecipeId(b, p);
    }
}