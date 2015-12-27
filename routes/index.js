var express = require('express');
var router = express.Router();
var redis = require('redis');
var Promise = require('bluebird');
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
    var config = 10;
    var names = [];
    var joblists = [];
    var jobIds =[];
    var data = [];
    for (var i = 1; i <= config; i++) {
        names.push('topic:' + i + ':name');
        joblists.push(getJobIds('topic:' + i + ':joblist'));
    }
    function getJobIds(keyname) {
        return new Promise(function(resolve, reject) {
            client.zrange(keyname, 0, -1, function(err, reply) {
                resolve(reply);
            })
        })
    } 
    Promise.all(joblists).then(function(value){
        jobIds = value;
        Promise.all([getJobTopics(), getJobTitles()]).then(function(){
            res.json(data);
        })
    })
    function getJobTopics() {
        return new Promise(function(resolve, reject) {
            client.mget(names, function(err, reply){
                for (var i = 0; i < reply.length; i++) {
                    data.push({'topic':reply[i]});
                } 
                resolve();
            })
        })
    }
    function getJobTitles() {
        return new Promise(function(resolve, reject) {
            var jobIdKeyNames = [];
            var count = 0;
            for (var i = 0; i < jobIds.length; i++) {
                for (var j = 0; j < jobIds[i].length; j++) {
                    jobIdKeyNames.push('job:' + jobIds[i][j] + ':title');
                }
                client.mget(jobIdKeyNames, function(err, reply) {
                    for (var i = 0; i < reply.length; i++) {
                        reply[i] = {'title':reply[i]};
                    }
                    data[count++].jobs = reply;
                })
                jobIdKeyNames = [];
            }
            resolve();
        })
    }

//router.get('/:path', function(req, res, next) {
  //res.render(req.params.path);
//});
});
module.exports = router;
