var app = angular.module("infor-you-mation",[]);
app.controller("indexCtrl", function($scope, $http){
        $http.get('/data')
        .success(function(res){
            $scope.lists = res;
        })
});

