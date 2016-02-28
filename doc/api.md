# api 接口说明

目前接口都还比较粗糙，待完善。

## 首页展示主题内容

|uri|method|
|:-:|:-:|
|`/topic`|GET|
|`/job/:id`|GET|

## 登录

|uri|method|type|
|:-:|:-:|:-:|
|`/user/login`|POST|`x-www-form-urlencoded`|

|key|description|
|:-:|:-:|
|username|用户名|
|password|密码|

## 注册

|uri|method|type|
|:-:|:-:|:-:|
|`/user/register`|POST|`x-www-form-urlencoded`|

|key|description|
|:-:|:-:|
|username|用户名|
|password|密码|

## 图片验证码

|uri|method|
|:-:|:-:|
|`/captcha`|GET|
