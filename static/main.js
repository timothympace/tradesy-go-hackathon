function postData(url, data) {
    // Default options are marked with *
    return fetch(url, {
        body: JSON.stringify(data),
        method: 'POST'
    })
}


// Items ------------

$('#item-add-form').on('submit', (e) => {
    const json = {};
    json['id'] = $('[name="id"]', e.target).val();
    json['name'] = $('[name="name"]', e.target).val();
    json['brand'] = $('[name="brand"]', e.target).val();
    json['price'] = parseFloat($('[name="price"]', e.target).val());
    json['userId'] = $('[name="userId"]', e.target).val();
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
            $('<th>').html('Price'),
            $('<th>').html('User ID')
        ),
        ...(json || []).map((item) => {
            return $('<tr>').append(
                $('<td>').html(item.id),
                $('<td>').html(item.name),
                $('<td>').html(item.brand),
                $('<td>').html(item.price),
                $('<td>').html(item.userId)
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

// Users ------------
$('#user-add-form').on('submit', (e) => {
    const json = {};
    json['id'] = $('[name="id"]', e.target).val();
    json['name'] = $('[name="name"]', e.target).val();
    json['email'] = $('[name="email"]', e.target).val();
    json['phone'] = $('[name="phone"]', e.target).val();
    postData('/addUser', json);
    e.preventDefault();
});

$('#user-delete-form').on('submit', (e) => {
    const json = {};
    json['id'] = $('[name="id"]', e.target).val();
    postData('/deleteUser', json);
    e.preventDefault();
});

function buildUserTable(json) {
    const table = $('<table>').append(
        $('<tr>').append(
            $('<th>').html('ID'),
            $('<th>').html('Name'),
            $('<th>').html('Email'),
            $('<th>').html('Phone')
        ),
        ...(json || []).map((item) => {
            return $('<tr>').append(
                $('<td>').html(item.id),
                $('<td>').html(item.name),
                $('<td>').html(item.email),
                $('<td>').html(item.phone)
            )
        })
    );

    $('#user-list-table').html(table)
}

$('#user-list-form').on('submit', (e) => {
    const json = {};
    postData('/allUsers', json).then((response) => {
        response.json().then(function(data) {
            buildUserTable(data);
        });
    });
    e.preventDefault();
});

// Searches -----------

function buildSearchResults(json) {
    const table = $('<table>').append(
        $('<tr>').append(
            $('<th>').html('ID'),
            $('<th>').html('Name'),
            $('<th>').html('Brand'),
            $('<th>').html('Price'),
            $('<th>').html('User ID'),
            $('<th>').html('User Name')
        ),
        ...(json || []).map((result) => {
            return $('<tr>').append(
                $('<td>').html(result.id),
                $('<td>').html(result.name),
                $('<td>').html(result.brand),
                $('<td>').html(result.price),
                $('<td>').html(result.userId),
                $('<td>').html(result.userName)
            )
        })
    );

    $('#search-results').html(table)
}

$('#searches-form').on('submit', (e) => {
    const json = {};
    json['name'] = $('[name="name"]', e.target).val();
    postData('/getItemByName', json).then((response) => {
        response.json().then(function(data) {
            buildSearchResults(data);
        });
    });
    e.preventDefault();
});