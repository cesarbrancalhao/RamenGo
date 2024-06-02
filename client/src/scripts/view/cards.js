import { getBroths, getProteins } from "../controller/apiController";

const brothCardContainer = document.querySelector('#broth');
const proteinCardContainer = document.querySelector('#protein');
const mainText = document.querySelectorAll('.main-text');
const btnCenter = document.querySelector('.btn-center');

const fadeIn = (element) => element.animate([{opacity: 0}, {opacity: 1}], {duration: 900, fill: 'forwards'});


const selectCard = (card, container) => {
    if (card.classList.contains('selected')) card.classList.remove('selected');
    else {
        Array.from(container.children).forEach(child => child.classList.remove('selected'));
        card.classList.add('selected');
    }
};

const createList = (list, container) => {
    for (const x of Object.entries(list)) {
        const card = document.createElement('div');
        card.classList.add('card');
        card.innerHTML = `
            <img src="${x[1].imageInactive}">
            <p class="card-title">${x[1].name}</p>
            <p class="card-description">${x[1].description}</p>
            <p class="card-price">US$ ${x[1].price}</p>`;
        card.onclick = () => selectCard(card, container);
        container.appendChild(card);

        fadeIn(card);
    }
}

export async function fillCards() {
    const brothList = await getBroths();
    const proteinList = await getProteins();

    fadeIn(mainText[0]);
    createList(brothList, brothCardContainer);
    createList(proteinList, proteinCardContainer);
    fadeIn(mainText[1]);
    fadeIn(btnCenter);
}