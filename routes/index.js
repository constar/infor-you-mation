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
/* POST */
router.post('/register', function(req, res, next) {
    if (req.body.password != req.body.checkpassword) {
        res.render('register', { powerBy: 'Express',display:'block' });

    } else {
        res.render('register', { powerBy: 'Express', display:'none' });
    }
})
module.exports = router;

