function postData(url, data) {
    // Default options are marked with *
    return fetch(url, {
        body: JSON.stringify(data),
        method: 'POST'
    })
}

$('#item-add-form').on('submit', (e) => {
    const json = {};
    json['id'] = $('[name="id"]', e.target).val();
    json['name'] = $('[name="name"]', e.target).val();
    json['brand'] = $('[name="brand"]', e.target).val();
    json['price'] = parseFloat($('[name="price"]', e.target).val());
    postData('/addItem', json);
    e.preventDefault();
});

$('#item-delete-form').on('submit', (e) => {
    const json = {};
    json['id'] = $('[name="id"]', e.target).val();
    postData('/deleteItem', json);
    e.preventDefault();
});

function buildItemTable(json) {
    const table = $('<table>').append(
        $('<tr>').append(
            $('<th>').html('ID'),
            $('<th>').html('Name'),
            $('<th>').html('Brand'),
            $('<th>').html('Price')
        ),
        ...json.map((item) => {
            return $('<tr>').append(
                $('<td>').html(item.id),
                $('<td>').html(item.name),
                $('<td>').html(item.brand),
                $('<td>').html(item.price)
            )
        })
    );

    $('#item-list-table').html(table)
}

$('#item-list-form').on('submit', (e) => {
    const json = {};
    postData('/allItems', json).then((response) => {
        response.json().then(function(data) {
            buildItemTable(data);
        });
    });
    e.preventDefault();
});