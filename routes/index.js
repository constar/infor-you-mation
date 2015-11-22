var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index');
});

//router.get('/:path', function(req, res, next) {
  //res.render(req.params.path);
//});
module.exports = router;
