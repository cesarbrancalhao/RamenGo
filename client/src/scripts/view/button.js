import { generateOrder } from "../controller/apiController";
import { openModal } from "./success";

const getSelected = () => document.querySelectorAll('.selected');

const turnActive = (btn, action) => { btn.classList.remove('inactive'); btn.onclick = action };

const turnInactive = (btn) => { btn.classList.add('inactive'); btn.onclick = null };

export const turnBtnActive = (button, action) => {
    getSelected().length === 2 ? 
    turnActive(button, action) : turnInactive(button);
};

export const requestOrder = () => {
    const brothId = parseInt(document.querySelector('#broth .selected').children[0].innerHTML);
    const proteinId = parseInt(document.querySelector('#protein .selected').children[0].innerHTML);
    const { name, image } = generateOrder(brothId, proteinId);
    openModal(name, image);
};