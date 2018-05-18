// file: js/require-setup.js
//
// Declare this variable before loading RequireJS JavaScript library
// To config RequireJS after it’s loaded, pass the below object into require.config();
var require = {
    baseUrl: '/assets/js/lib',
    urlArgs: "v=0.0.1",
    paths: {
        'app': '../app',
        'jquery': 'jquery/jquery-2.1.3.min',
        'jquery-form': 'jquery-form/jquery.form.3.51',
        'bootstrap': 'bootstrap/bootstrap.min',
        'jquery.treetable': 'jquery-treetable/jquery.treetable'
    },
    shim : {
        "bootstrap" : { "deps" :['jquery']},
    }
};
