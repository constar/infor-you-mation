var db =  require('mongoose');

db.connect("mongodb://localhost:27017/test");

var userSchema = new db.Schema({
    username : { type: String, default: '' ,trim : true, index: {unique: true} },
    password : { type: String, default: '' ,trim : true},
    email: { type: String, default: '' }
}, {collection: 'user'});

var User = db.model('User', userSchema);

/*
 * user: {
 *   username:
 *   password:
 *   email:
 * }
 * callback: function(err, user)
 */
function register(user, callback) {
    var u = new User({
        username: user.username,
        password: user.password,
        email: user.email,
    });
    
    u.save(callback);
}

/*
 * user: {
 *   username:
 *   password:
 * }
 * callback: function(err);
 */
function login(user, callback) {
    User.find({username: user.username}, function(err, docs) {
        if (err) {
            callback(err);
            return;
        }
        if (docs.length != 1) {
            callback("db data illegal");
            return;
        }
        if (docs[0].password !== user.password) {
            callback("password error");
            return;
        }
        callback(null);
    });
}

module.exports = {
    register: register,
    login: login,
};
