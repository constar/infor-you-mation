var router = require('express').Router();
var redis = require('redis');
var captchapng = require('captchapng');
client = redis.createClient();

router.get('/', function(req, res) {
    var v = parseInt(Math.random()*9000+1000);
    var p = new captchapng(80,30, v);
    // First color: background (red, green, blue, alpha) 
    p.color(0, 0, 0, 0);  
    // Second color: paint (red, green, blue, alpha) 
    p.color(80, 80, 80, 255); 

    var img = p.getBase64();
    var imgbase64 = new Buffer(img,'base64');
    res.writeHead(200, {
        'Content-Type': 'image/png'
    });
    res.end(imgbase64);
});
module.exports = router;
