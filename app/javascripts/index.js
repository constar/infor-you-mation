var app = angular.module("infor-you-mation",['ngRoute','ngSanitize','ngCookies']);
app.config(['$routeProvider',
    function($routeProvider) {
    $routeProvider
    .when('/', {
        templateUrl:'/pages/index.html',
        controller:'indexCtrl'
    })
    .when('/register', {
        templateUrl:'/pages/register.html',
        controller:'registerCtrl'
    })
    .when('/login', {
        templateUrl:'/pages/login.html',
        controller:'loginCtrl'
    })
    .when('/home', {
        templateUrl:'/pages/home.html',
        controller:'homeCtrl'
    })
    .when('/detail/:id', {
        templateUrl:'/pages/detail.html',
        controller:'detailCtrl'
    })
    .when('/list/:topicId', {
        templateUrl:'/pages/list.html',
        controller:'listCtrl'
    });
}]);
app.controller("indexCtrl", ['$scope', '$http', '$cookies', '$rootScope', function($scope, $http, $cookies, $rootScope){
    $http.get('/user')
    .success(function(res) {
        $rootScope.isLogin = res.userid;
    });
    $http.get('/topic')
    .success(function(res){
        $scope.lists = res;
        $scope.notEmpty = function(item) {
            return item.total != 0;
        }
    });
    $rootScope.logout = function() {
        $http.post('/user/logout').success(function(res) {
            $rootScope.isLogin = !res.success;
        });
    }
}]);
app.controller("registerCtrl", function($scope, $http){
    $scope.err = false;
    $scope.submit = function() {
        $http.post('/user/register', $scope.user)
        .success(function(res){
            if(res.success){
                $scope.err = false;
                window.location.href = "/";
                //alert('注册成功！欢迎加入邮订阅！');
            } else {
                $scope.err = '用户名已存在！';
            }
        })
    }
});
app.controller("loginCtrl", function($scope, $http){
    $scope.err = false;
    $scope.submit = function() {
        $http.post('/user/login', $scope.user)
        .success(function(res){
            if (res.success) {
                $scope.err = false;
                window.location.href = "/";
            } else {
                $scope.err = "用户名或密码错误！"
            }
        })
    }
});
app.controller("detailCtrl", ['$scope', '$http', '$routeParams', function($scope, $http, $routeParams){
    $http.get('/job/' + $routeParams.id).success(function(res) {
        $scope.title = res.title;
        $scope.source = res.source;
        $scope.url = res.url;
        $scope.content = res.content;
    });
}]);
app.controller("listCtrl", ['$scope', '$http', '$routeParams', function($scope, $http, $routeParams){
    $http.get('/topic/' + $routeParams.topicId).success(function(res) {
        $scope.topic = res.topic;
        $scope.total = res.total;
        $scope.jobs = res.jobs;
    });
}]);
app.controller("homeCtrl", function($scope){
});
