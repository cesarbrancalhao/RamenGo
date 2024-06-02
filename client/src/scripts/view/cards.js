import { getBroths, getProteins } from "../controller/apiController";

const brothCardContainer = document.querySelector('#broth');
const proteinCardContainer = document.querySelector('#protein');
const firstMainText = document.querySelector('#main-t1');
const secondMainText = document.querySelector('#main-t2');
const btnCenter = document.querySelector('.btn-center');

const fadeIn = (element) => element.animate([{opacity: 0}, {opacity: 1}], {duration: 900, fill: 'forwards'});

const createList = (list, container) => {
    for (const x of Object.entries(list)) {
        const card = document.createElement('div');
        card.classList.add('card');
        card.innerHTML = `
            <img src="${x[1].imageInactive}">
            <p class="card-title">${x[1].name}</p>
            <p class="card-description">${x[1].description}</p>
            <p class="card-price">US$ ${x[1].price}</p>
        `;
        container.appendChild(card);

        fadeIn(card);
    }
}

export async function fillCards() {
    const brothList = await getBroths();
    const proteinList = await getProteins();

    fadeIn(firstMainText);
    createList(brothList, brothCardContainer);
    createList(proteinList, proteinCardContainer);
    fadeIn(secondMainText);
    fadeIn(btnCenter);
}