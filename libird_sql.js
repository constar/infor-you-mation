var libird = require('libird');
var Promise = require('bluebird');
var mysql = require('mysql');
var redis = require('redis');
client = redis.createClient();
var router = libird.router;

var connection = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: '',
    database: 'youdingyue'
})

libird.setDirPath('./app');
router.get('/user', function(req, res) {
    getUser(req.getSession('userid'))
    .then(function(userinfo) {
        res.send(userinfo);
    });
});
router.post('/user/register', function(req, res) {
    var username = req.body.username;
    var password = req.body.password;
    if (!username) {
        res.send({'error': 'username not found'});
        return;
    }
    if (!password) {
        res.send({'error': 'password not found'});
        return;
    }
    
    client.get('user:' + username + ':id', function(err, reply) {
        if (err) {
            res.send({'error': err});
            return;
        }
        if (reply) {
            res.send({'error': 'username: ' + username + " already exists", 'success': false});
            return;
        } 
        client.incr('user:nextid', function(err, newid) {
            if (err) {
                res.send({'error': err, 'success': false});
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
                        res.send({'error': err, 'success': false});
                        return;
                    }
                    res.setCookie('userid', newid);
                    res.send({'msg': 'register ok', 'success': true});
            });
        });
    });
});
router.post('/user/login', function(req, res) {
    var username = req.body.username;
    var password = req.body.password;
    client.get('user:' + username + ':id', function(err, id) {
        if (err) {
            res.send({'error': err});
            return;
        }
        if (!id) {
            res.send({'error': 'username: ' + username + ' not found'});
            return;
        }
        client.get('user:' + id + ':password', function(err, reply) {
            if (password == reply) {
                    res.setCookie('userid', id);
                    res.send({'msg': 'login ok', 'success': true});
            } else {
                res.send({'error': 'password error', 'success': false});
            }
        });
    });
});
router.post('/user/logout', function(req, res) {
    res.clearCookie('userid');
    res.send({'success': true});
});

router.get('/topic', function(req, res) {
    var limit = 5;
    getTopicId().then(function(topicIds) {
        var promises = [];
        for (var i = 0; i < topicIds.length; i++) {
            promises.push(getTopicInfo(topicIds[i], limit));
        }
        return Promise.all(promises);
    }).then(function (data) {
        res.send(data);
    });
})
router.get('/topic/:id', function(req, res) {
    getTopicInfo(req.params.id).then(function (info) {
        res.send(info);
    });
});
router.get('/job/:id', function(req, res) {
    getJobInfo(req.params.id, true).then(function (jobinfo) {
        res.send(jobinfo);
    });
});
function getTopicId() {
    return new Promise(function(resolve) {
        var getTopicIdSql = 'select id from topics';
        connection.query(getTopicIdSql, function(err, rows, field) {
            if (err) {
                console.log(err);
            } else {
                var topicIds = [];
                for(var i = 0; i < rows.length; i++) {
                    topicIds.push(rows[i].id);
                    resolve(topicIds);
                }
            }
        })
    });
}
function getTopicInfo(topicid, limit) {
    var topicinfo = {'topicId': topicid}
    return (new Promise(function(resolve) {
        var getNameSql = 'select name from topics where id = ?';
        var getNameSql_Params = [topicid];
        connection.query(getNameSql, getNameSql_Params, function(err, rows, field) {
            if (err) {
                console.log(err);
            } else {
                topicinfo.topic = rows[0].name;
                resolve(rows[0].name);
            }
        })
    })).then(function(topic) {
        return new Promise(function(resolve) {
            var max = 500;
            var getJobIdSql = 'select jobid from jobs_topic where topic = ? order by modify_time limit ?';
            var getJobIdSql_Params = [topic, max];
            connection.query(getJobIdSql, getJobIdSql_Params, function(err, rows, field) {
                if (err) {
                    console.log(err);
                } else {
                    var total = rows.length;
                    topicinfo.total = total;
                    if(limit) {
                        total = total < limit ? total : limit; 
                    }
                    var jobIds = [];
                    for (var i = 0; i < total; i++) {
                        jobIds.push(rows[i].jobid);
                    }
                    resolve(jobIds);
                }
            })
        });
    }).then(function(jobIds) {
        var promises = [];
        for (var i = 0; i < jobIds.length; i++) {
            promises.push(getJobInfo(jobIds[i]));
        }
        return Promise.all(promises);
    }).then(function(jobinfos) {
        topicinfo.jobs = jobinfos;
        return topicinfo;
    });
}

function getJobInfo(jobid, withcontent) {
    return  new Promise(function(resolve) {
        if (withcontent) {
            var getJobInfoSql = 'select title,source,url,content from jobs where jobid = ?';
        } else {
            var getJobInfoSql = 'select title,source,url from jobs where jobid = ?';
        }
        var getJobInfoSql_Params = [jobid];
        connection.query(getJobInfoSql, getJobInfoSql_Params, function(err, rows, field) {
            if (err) {
                console.log(err);
            } else {
                if (withcontent) {
                    resolve({"id": jobid, "title": rows[0].title, "url": rows[0].url, "source": rows[0].source, "content": rows[0].content});
                } else {
                    resolve({"id": jobid, "title": rows[0].title, "url": rows[0].url, "source": rows[0].source});
                }
            }
        })
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

libird.start(8888);
