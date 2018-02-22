/* eslint-disable */
export function querystring (search) {
    if( search.length === 0 ) return search;
    var query = {};
    var slice = search.replace('?', '');
    var arr = slice.split('&')
    arr.forEach(function (el) {
        var split = el.split('=');
        query[split[0]] = decodeURIComponent(split[1]);
    });
    return query;
}

export function fromNow (date) {
    return moment(date).fromNow();
}

