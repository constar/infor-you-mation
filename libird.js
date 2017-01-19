var libird = require('libird');
var redis = require('redis');
var Promise = require('bluebird');
client = redis.createClient();
var router = libird.router;

libird.setDirPath('./app');
//router.get('/user', function(req, res, next) {
    //getUser(req.session.userid)
    //.then(function(userinfo) {
        //res.json(userinfo);
    //});
//});
router.post('/user/register', function(req, res) {
    var username = req.body.username;
    var password = req.body.password;
    if (!username) {
        res.send({'error': 'username not found'}, 'json');
        return;
    }
    if (!password) {
        res.send({'error': 'password not found'}, 'json');
        return;
    }
    
    client.get('user:' + username + ':id', function(err, reply) {
        if (err) {
            res.send({'error': err}, 'json');
            return;
        }
        if (reply) {
            res.send({'error': 'username: ' + username + " already exists", 'success': false}, 'json');
            return;
        } 
        client.incr('user:nextid', function(err, newid) {
            if (err) {
                res.send({'error': err, 'success': false}, 'json');
                return;
            }
            client.mset(['user:' + newid + ':username', 
                username,
                'user:' + newid + ':password',
                password,
                'user:' + username + ':id',
                newid], 
                function(err) {
                    if (err) {
                        res.send({'error': err, 'success': false}, 'json');
                        return;
                    }
                    //res.cookie('SESSIONID', req.sessionID, req.session.cookie);
                    res.send({'msg': 'register ok', 'success': true}, 'json');
            });
        });
    });
});
router.post('/user/login', function(req, res) {
    var username = req.body.username;
    var password = req.body.password;
    client.get('user:' + username + ':id', function(err, id) {
        if (err) {
            res.send({'error': err}, 'json');
            return;
        }
        if (!id) {
            res.send({'error': 'username: ' + username + ' not found'}, 'json');
            return;
        }
        client.get('user:' + id + ':password', function(err, reply) {
            if (password == reply) {
                //req.session.regenerate(function() {
                    //req.session.userid = id;
                    //req.session.save();
                    res.send({'msg': 'login ok', 'success': true}, 'json');
                //});
            } else {
                res.send({'error': 'password error', 'success': false}, 'json');
            }
        });
    });
});
router.post('/user/logout', function(req, res) {
    res.clearCookie('connect.sid');
    req.session.destroy(function() {
        res.send({'success': true}, 'json');
    });
});
router.get('/topic', function(req, res) {
    var limit = 5;
    getTopicMaxID().then(function(maxid) {
        var promises = [];
        for (var i = 1; i <= maxid; i++) {
            promises.push(getTopicInfo(i, limit));
        }
        return Promise.all(promises);
    }).then(function (data) {
        res.send(data, 'json');
    });
})
router.get('/topic/:id', function(req, res) {
    getTopicInfo(req.params.id).then(function (info) {
        res.send(info, 'json');
    });
});
router.get('/job/:id', function(req, res) {
    getJobInfo(req.params.id, true).then(function (jobinfo) {
        res.send(jobinfo, 'json');
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
    var topicinfo = {'topicId': topicid}
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

libird.start(8888);
