var app = angular.module("infor-you-mation",['ngRoute']);
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
app.controller("indexCtrl", function($scope, $http){
    $http.get('/topic')
    .success(function(res){
        $scope.lists = res;
    })
});
app.controller("registerCtrl", function($scope){
});
app.controller("loginCtrl", function($scope){
});
app.controller("homeCtrl", function($scope){
});

