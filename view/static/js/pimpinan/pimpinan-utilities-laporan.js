function searchGoogle(index) {
    // Get the <a> element
    var linkElement = document.querySelector('td.img-click-'+index+' a');
    console.log(linkElement);

    // Update the href attribute with the coordinate value
    if (linkElement) {
        linkElement.click();
    }
}

function searchGoogleKetua() {
    // Get the <a> element
    var linkElement = document.querySelector('div#img-ketua a');
    console.log(linkElement);

    // Update the href attribute with the coordinate value
    if (linkElement) {
        linkElement.click();
    }
}