import { getBroths, getProteins } from "../controller/apiController";
import { turnBtnActive } from "./order";

const brothCardContainer = document.querySelector('#broth');
const proteinCardContainer = document.querySelector('#protein');
const mainText = document.querySelectorAll('.main-text');
const btnCenter = document.querySelector('.btn-center');

const fadeIn = (element) => element.animate([{opacity: 0}, {opacity: 1}], {duration: 900, fill: 'forwards'});

const createList = (list, container) => {
    for (const x of Object.entries(list)) {
        const card = document.createElement('div');
        card.classList.add('card');
        card.innerHTML = `
            <img class="card-image" id="inactive" src="${x[1].imageInactive}">
            <img class="card-image hidden" id="active" src="${x[1].imageActive}">
            <p class="card-title">${x[1].name}</p>
            <p class="card-description">${x[1].description}</p>
            <p class="card-price">US$ ${x[1].price}</p>`;
        card.onclick = (e) => { 
            e.stopPropagation(); 
            selectCard(card, container); };
        card.addEventListener('click', () => turnBtnActive(document.querySelector('#order'), 
            () => {}));
        container.appendChild(card);

        fadeIn(card);
    }
};

const removeActive = card => {
    card.classList.remove('selected');

    const imgActive = card.querySelector('#active');
    const imgInactive = card.querySelector('#inactive');

    imgActive.classList.add('hidden');
    imgInactive.classList.remove('hidden');
};

const turnActive = (card, container) => {
    Array.from(container.children).forEach(child => {
        if (child !== card) {

            child.classList.remove('selected');

            const imgActive = child.querySelector('#active');
            const imgInactive = child.querySelector('#inactive');

            imgActive.classList.add('hidden');
            imgInactive.classList.remove('hidden');
        }
    });

    card.classList.add('selected');
    const imgActive = card.querySelector('#active');
    const imgInactive = card.querySelector('#inactive');

    imgActive.classList.remove('hidden');
    imgInactive.classList.add('hidden');
};

const selectCard = (card, container) => {
    card.classList.contains('selected') ?
        removeActive(card) : 
        turnActive(card, container);
};

export async function fillCards() {
    const brothList = await getBroths();
    const proteinList = await getProteins();

    fadeIn(mainText[0]);
    createList(brothList, brothCardContainer);
    createList(proteinList, proteinCardContainer);
    fadeIn(mainText[1]);
    fadeIn(btnCenter);
};