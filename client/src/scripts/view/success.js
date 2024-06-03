const modal = document.querySelector('.modal');

const generateModalHTML = (name, image) => `
    <div class="modal-container">

        <div class="modal-header">

            <div class="logo-container">
                <a href="./index.html"><img src="./src/assets/icons/logo-red.svg"  alt="Logo"></a>
            </div>

            <div class="center">
                <img src="${image}" id="dish-img" alt="">
            </div>

            <div>
                <p>Your Order:</p>
                <h2>${name ? name : "We couldn't find this dish name :("}</h2>
            </div>

        </div>

        <div class="modal-main center">

            <div class="message">
                <div class="center"><img class="success-bowing" src="./src/assets/icons/bowing.svg" alt=""></div>
                <div class="center"><img class="success-text" src="./src/assets/icons/bowing-text.png" alt=""></div>
                <div class="center"><h2>Your order is being prepared</h2></div>
                <div class="center"><p>Hold on, when you least expect you will be eating your rámen.</p></div>
                <div class="center"><a class="btn arrow" id="order">Place new order</a></div>
            </div>

        </div>

    </div>`;


export const openModal = (name, image) => {
    modal.innerHTML = generateModalHTML(name, image);
    modal.classList.remove('hidden');
}

export const closeModal = () => {
    modal.classList.add('hidden');
    modal.innerHTML = '';
}