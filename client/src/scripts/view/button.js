const turnActive = (btn, action) => { btn.classList.remove('inactive'); btn.onclick = action };

const turnInactive = (btn) => { btn.classList.add('inactive'); btn.onclick = null };

export const turnBtnActive = (button, action) => {
    document.querySelectorAll('.selected').length === 2 ? 
    turnActive(button, action) : turnInactive(button);
};

export const requestOrder = (/*brothid, proteinid*/) => {
    return {
    name: "Dish",
    image: "https://tech.redventures.com.br/icons/ramen/ramenYasai%20Vegetarian.png"
}};