var $ = require('jquery');
var Ractive = require('ractive');
var ractive = new Ractive({
    el:$('.main-container'),
    template:require('/partials/index.html'),
    data:{
        name:'hi',
    }
})
alert('hi');
