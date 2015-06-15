var express = require('express');
var nodejieba = require('nodejieba');

var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { powerBy: 'Express' });
});

router.get('/cut/:sentence', function(req, res, next) {
    var words = nodejieba.cut(req.params.sentence);
    console.log(words);
    res.send(words);
});
module.exports = router;

