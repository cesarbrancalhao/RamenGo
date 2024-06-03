window.onload = function() {
    const headerImage = document.querySelector('.header-image');
    const headerText = document.querySelector('.header-text');

    if (headerImage.complete)
        headerText.animate([{opacity: 0}, {opacity: 1}], {duration: 1800, fill: 'forwards'});
    else
        headerImage.onload(headerText.animate([{opacity: 0}, {opacity: 1}], {duration: 1800, fill: 'forwards'}));
} // Didn't want to do this but css animations were messing up div positions (only in Edge)