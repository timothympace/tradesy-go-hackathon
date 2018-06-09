function postData(url, data) {
    // Default options are marked with *
    return fetch(url, {
        body: JSON.stringify(data),
        method: 'POST'
    })
}

$('#item-form').on('submit', (e) => {
    const json = {};
    json['id'] = $('[name="id"]', e.target).val();
    json['name'] = $('[name="name"]', e.target).val();
    json['brand'] = $('[name="brand"]', e.target).val();
    json['price'] = parseFloat($('[name="price"]', e.target).val());
    postData('/addItem', json);
    e.preventDefault();
});