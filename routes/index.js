var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { powerBy: 'Express' });
});

router.get('/register', function(req, res, next) {
  res.render('register', { powerBy: 'Express' });
});

router.get('/login', function(req, res, next) {
    res.render('login', { powerBy: 'Express'});
})
module.exports = router;
