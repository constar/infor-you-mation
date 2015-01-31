# InforYouMation

The information for you .

## Dependency

+ [beego]

## MysqlDB

```
CREATE DATABASE `inforyoumation` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
```

```
CREATE TABLE user (
id      INT,
username   VARCHAR(32),
pwd VARCHAR(32),
PRIMARY KEY(id) 
);
```

## Start

```
go build
./inforyoumation
```



[beego]:https://github.com/astaxie/beego.git
