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
id      INT auto_increment,
username   VARCHAR(32),
pwd VARCHAR(32),
PRIMARY KEY(id) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8 ;
```

## Start

```
go build
./inforyoumation
```

## Contact

```
i@yanyiwu.com
```


[beego]:https://github.com/astaxie/beego.git
