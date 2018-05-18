require(['jquery','jquery.treetable'], function () {

    var $ = require('jquery');
    require('jquery.treetable');

    $(".treetable").treetable({ expandable: true, initialState: 'collapsed', indent: 10 });

});