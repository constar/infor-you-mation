var express = require('express');
var router = express.Router();
var redis = require('redis');
var Promise = require('bluebird');
client = redis.createClient();

/* GET home page. */
router.get('/', function(req, res, next) {
    res.render('index');
});

router.get('/user', function(req, res, next) {
    getUser(req.session.userid)
    .then(function(userinfo) {
        res.json(userinfo);
    });
});

/* GET home data */
router.get('/topic', function(req, res, next) {
    var limit = 5;
    getTopicMaxID().then(function(maxid) {
        var promises = [];
        for (var i = 1; i <= maxid; i++) {
            promises.push(getTopicInfo(i, limit));
        }
        Promise.all(promises).then(function (data) {
            res.json(data);
        });
    });
});

router.get('/topic/:id', function(req, res, next) {
    getTopicInfo(req.params.id).then(function (info) {
        res.json(info);
    });
});

router.get('/job/:id', function(req, res) {
    getJobInfo(req.params.id, true).then(function (jobinfo) {
        res.json(jobinfo);
    });
});

function getTopicMaxID() {
    return new Promise(function(resolve) {
        client.get("topic:nextid", function(err, reply) {
            resolve(reply);
        });
    });
}
function getTopicInfo(topicid, limit) {
    var topicinfo = {}
    return (new Promise(function(resolve) {
        client.get("topic:" + topicid + ":name", function(err, reply) {
            topicinfo.topic = reply;
            resolve();
        });
    })).then(function() {
        return new Promise(function(resolve) {
            client.zcard('topic:' + topicid + ':joblist', function(err, reply) {
                topicinfo.total = reply;
                resolve();
            });
        });
    }).then(function() {
        return new Promise(function(resolve) {
            var lim = limit || 0;
            client.zrevrange('topic:' + topicid + ':joblist', 0, lim - 1, function(err, jobids) {
                topicinfo.jobids = jobids;
                resolve();
            });
        });
    }).then(function() {
        var promises = [];
        for (var i = 0; i < topicinfo.jobids.length; i++) {
            promises.push(getJobInfo(topicinfo.jobids[i]));
        }
        return Promise.all(promises);
    }).then(function(jobinfos) {
        delete topicinfo.jobids;
        topicinfo.jobs = jobinfos;
        return topicinfo;
    });
}

function getUser(userid) {
    return new Promise(function(resolve) {
        client.mget(['user:' + userid + ':username'], function (err, reply) {
            if (err) {
                resolve({});
                return
            }
            resolve({'userid': userid, 'username': reply[0]});
        });
    });
}

function getJobInfo(jobid, withcontent) {
    var p = new Promise(function(resolve) {
        var titlekey = "job:" + jobid + ":title";
        var urlkey = "job:" + jobid + ":url";
        var sourcekey = "job:" + jobid + ":source";
        var contentkey = "job:" + jobid + ":content";
        if (withcontent) {
            client.mget([titlekey, urlkey, sourcekey, contentkey], function(err, reply) {
                resolve({"id": jobid, "title": reply[0], "url": reply[1], "source": reply[2], "content": reply[3]});
            });
        } else {
            client.mget([titlekey, urlkey, sourcekey], function(err, reply) {
                resolve({"id": jobid, "title": reply[0], "url": reply[1], "source": reply[2]});
            });
        }
    });
    return p;
}

module.exports = router;
