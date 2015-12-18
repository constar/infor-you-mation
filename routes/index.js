var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index');
});

router.get('/data', function(req, res, next) {
    res.json([{
        'keyword':'PHP',
        'arrs':[{
            'text':'标题1',
            'url':'http://baidu.com'
        },{
            'text':'标题2',
            'url':'http://baidu.com'
        },{
            'text':'标题3',
            'url':'http://baidu.com'
        }]
    },{
        'keyword':'Java',
        'arrs':[{
            'text':'标题1',
            'url':'http://baidu.com'
        },{
            'text':'标题2',
            'url':'http://baidu.com'
        },{
            'text':'标题3',
            'url':'http://baidu.com'
        }]
    },{
        'keyword':'CSS',
        'arrs':[{
            'text':'标题1',
            'url':'http://baidu.com'
        },{
            'text':'标题2',
            'url':'http://baidu.com'
        },{
            'text':'标题3',
            'url':'http://baidu.com'
        },{
            'text':'标题4',
            'url':'http://baidu.com'
        }]
    }]
   );
});

//router.get('/:path', function(req, res, next) {
  //res.render(req.params.path);
//});
module.exports = router;
