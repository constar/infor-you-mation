var app = angular.module("infor-you-mation",['ngRoute','ngSanitize']);
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
    $scope.isHide = true;
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
});
app.controller("registerCtrl", function($scope){
});
app.controller("loginCtrl", function($scope){
});
app.controller("homeCtrl", function($scope){
});

