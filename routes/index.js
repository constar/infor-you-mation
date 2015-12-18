var express = require('express');
var router = express.Router();
var redis = require('redis');
client = redis.createClient();

//var data = [{
        //'keyword':'PHP',
        //'arrs':[{
            //'text':'标题1',
            //'url':'http://baidu.com'
        //},{
            //'text':'标题2',
            //'url':'http://baidu.com'
        //},{
            //'text':'标题3',
            //'url':'http://baidu.com'
        //}]
    //},{
        //'keyword':'Java',
        //'arrs':[{
            //'text':'标题1',
            //'url':'http://baidu.com'
        //},{
            //'text':'标题2',
            //'url':'http://baidu.com'
        //},{
            //'text':'标题3',
            //'url':'http://baidu.com'
        //}]
    //},{
        //'keyword':'CSS',
        //'arrs':[{
            //'text':'标题1',
            //'url':'http://baidu.com'
        //},{
            //'text':'标题2',
            //'url':'http://baidu.com'
        //},{
            //'text':'标题3',
            //'url':'http://baidu.com'
        //},{
            //'text':'标题4',
            //'url':'http://baidu.com'
        //}]
//}];


/* GET home page. */
router.get('/', function(req, res, next) {
    res.render('index');
});

router.get('/data', function(req, res, next) {
    client.get('data', function(err, reply) {
        res.json(JSON.parse(reply));
        console.log(reply);
    })
});

//router.get('/:path', function(req, res, next) {
  //res.render(req.params.path);
//});
module.exports = router;
