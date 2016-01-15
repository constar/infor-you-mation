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
    });
}]);
app.controller("indexCtrl", function($scope, $http){
    $http.get('/data')
    .success(function(res){
        $scope.lists = res;
    })
});
app.controller("registerCtrl", function($scope){
});

