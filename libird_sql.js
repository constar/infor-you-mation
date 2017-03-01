var libird = require('libird');
var Promise = require('bluebird');
var mysql = require('mysql');
var crypto = require('crypto');
var router = libird.router;

var connection = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: '',
    database: 'youdingyue'
})

libird.setDirPath('./app');
router.get('/user', function(req, res) {
    var userid = req.getSession('userid');
    if (!userid) {
        res.send({});
    } else {
        getUser(userid)
        .then(function(userinfo) {
            res.send(userinfo);
        });
    }
});
router.post('/user/register', function(req, res) {
    var username = req.body.username;
    var password = req.body.password;
    var checkUserSql = 'select userid from users where username = ?';
    var checkUserSql_Params = [username];
    connection.query(checkUserSql, checkUserSql_Params, function(err, rows, fields) {
        if (err) {
            console.log(err);
            res.send({'error': err});
        } else if (rows.length) {
            res.send({'error': 'username: ' + username + " already exists", 'success': false});
        } else {
            var str = username + new Date().getTime();
            var userid = crypto.createHash('md5').update(str).digest('hex');
            var pwd = crypto.createHash('md5').update(password).digest('hex');
            var addUserSql = 'insert into users (userid, username, password) values (?, ?, ?)';
            var addUserSql_Params = [userid, username, pwd];
            connection.query(addUserSql, addUserSql_Params, function(err, rows, field) {
                if(err) {
                    console.log(err);
                    res.send({'error': err});
                } else {
                    res.setCookie('userid', userid);
                    res.send({'msg': 'register ok', 'success': true});
                }
            })
        }
    })
});
router.post('/user/login', function(req, res) {
    var username = req.body.username;
    var password = req.body.password;
    var pwd = crypto.createHash('md5').update(password).digest('hex');
    var getUserSql = 'select userid, password from users where username = ?';
    var getUserSql_Params = [username];
    connection.query(getUserSql, getUserSql_Params, function(err, rows, fields) {
        if (err) {
            console.log(err);
            res.send({'error': err});
        } else if(!rows.length) {
            res.send({'error': 'username ' + username + ' not found'});
        } else if (pwd != rows[0].password) {
            res.send({'error': 'password error', 'success': false});
        } else {
            res.setCookie('userid', rows[0].userid);
            res.send({'msg': 'login ok', 'success': true});
        }
    })
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
        connection.query(getTopicIdSql, function(err, rows, fields) {
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
    return new Promise (function(resolve) {
        var getNameSql = 'select name from topics where id = ?';
        var getNameSql_Params = [topicid];
        connection.query(getNameSql, getNameSql_Params, function(err, rows, fields) {
            if (err) {
                console.log(err);
            } else if (rows.length) {
                topicinfo.topic = rows[0].name;
                resolve(rows[0].name);
            } else {
                resolve();
            }
        })
    }).then(function(topic) {
        if(topic) {
            return new Promise(function(resolve) {
                var max = 500;
                var getJobIdSql = 'select jobid from jobs_topic where topic = ? order by modify_time limit ?';
                var getJobIdSql_Params = [topic, max];
                connection.query(getJobIdSql, getJobIdSql_Params, function(err, rows, fields) {
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
        } else {
            return 'topic not found';
        }
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
        connection.query(getJobInfoSql, getJobInfoSql_Params, function(err, rows, fields) {
            if (err) {
                console.log(err);
            } else if (rows.length){
                if (withcontent) {
                    resolve({"id": jobid, "title": rows[0].title, "url": rows[0].url, "source": rows[0].source, "content": rows[0].content});
                } else {
                    resolve({"id": jobid, "title": rows[0].title, "url": rows[0].url, "source": rows[0].source});
                }
            } else {
                resolve('jobid not found')
            }
        })
    });
}
function getUser(userid) {
    if(userid)
    return new Promise(function(resolve) {
        var getUserSql = 'select username from users where userid = ?';
        var getUserSql_Params = [userid];
        connection.query(getUserSql, getUserSql_Params, function(err, rows, fields) {
            if (err) {
                console.log(err);
            } else if (rows.length){
                resolve({'userid': userid, 'username': rows[0].username});
            } else {
                resolve({});
            }
        })
    });
}

libird.start(8888);
