var express = require('express');
var router = express.Router();
var userDB = require('../model/user');

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});

router.get('/register', function(req, res, next) {
  res.render('register', { powerBy: 'Express', display:'none'});
});

router.get('/login', function(req, res, next) {
    res.render('login', { powerBy: 'Express', display: 'none'});
});

router.post('/register', function(req, res, next) {
  if (req.body.password != req.body.checkpassword) {
      res.render('register', { powerBy: 'Express',display:'block' });
      return;
  }
  var u = {
    username: req.body.username,
    password: req.body.password,
    email: req.body.email,
  };
  userDB.register(u, function(err, u) {
    if (err) {
      res.render('register', { powerBy: 'Express',display:'block' });
      return;
    }
    res.redirect('/');
  });
});

router.post('/login', function(req, res, next) {
  var u = {
    username: req.body.username,
    password: req.body.password,
  };
  userDB.login(u, function(err) {
    if (err) {
      res.render('login', { powerBy: 'Express',display:'block' });
      return;
    }
    res.redirect('/');
  })
});

module.exports = router;
