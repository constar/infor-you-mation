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
    });
}]);
app.controller("indexCtrl", ['$scope', '$http', '$cookies', '$rootScope', function($scope, $http, $cookies, $rootScope){
    $scope.isHide = true;
    $rootScope.isLogin = $cookies.SESSIONID;
    $http.get('/topic')
    .success(function(res){
        $scope.lists = res;
        $scope.showJobDetail = function(job){
            $http.get('/job/' + job.id).success(function(res) {
                $scope.title = job.title;
                $scope.source = job.source;
                $scope.url = job.url;
                $scope.content = res.content; //TODO
                $scope.isHide = !$scope.isHide;
            });
        }
        $scope.close = function() {
            $scope.isHide = true;
        }
    })
}]);
app.controller("registerCtrl", function($scope, $http){
    $scope.err = false;
    $scope.submit = function() {
        $http.post('/user/register', $scope.user)
        .success(function(res){
            if(res.success){
                $scope.err = false;
                alert('注册成功！欢迎加入邮订阅！');
            } else {
                //$scope.err = res.error;
                $scope.err = '用户名已存在！';
            }
        })
    }
});
app.controller("loginCtrl", function($scope, $http){
    $scope.err = false;
    $scope.submit = function() {
        $http.post('user/login', $scope.user)
        .success(function(res){
            if (res.success) {
                $scope.err = false;
                alert('登录成功！');
            } else {
                $scope.err = "用户名或密码错误！"
            }
        })
    }
});
app.controller("homeCtrl", function($scope){
});
