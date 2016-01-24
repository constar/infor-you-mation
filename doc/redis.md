# Redis 数据库设计

## 数据结构

### Topic

信息的展示是通过主题聚合，也可以说是订阅。

|key|type|description|
|:-:|:-:|:-:|
|`topic:<id>:name`|string|主题名|
|`topic:<id>:words`|set|主题相关的word列表|
|`topic:<id>:joblist`|zset|主题相关的Job的id列表|

|key|type|description|
|:-:|:-:|:-:|
|`topic:<name>:id`|string|主题名到id的反向映射|

|key|type|description|
|:-:|:-:|:-:|
|`topic:nextid`|string|topic表的自增ID|

### Job

目前的信息都是一切招聘信息，所以目前取名叫 job 作为表名。

|key|type|description|
|:-:|:-:|:-:|
|`job:<id>:title`|string|标题|
|`job:<id>:content`|string|内容|
|`job:<id>:url`|string|url|
|`job:<id>:urlmd5`|string|url的MD5值，爬虫去重时需要使用|

|key|type|description|
|:-:|:-:|:-:|
|`job:<urlmd5>:id`|string|urlmd5到id的反向映射，对url的MD5值进行查询和去重|

|key|type|description|
|:-:|:-:|:-:|
|`job:nextid`|string|job表的自增ID|

### User

|key|type|description|
|:-:|:-:|:-:|
|`user:<id>:username`|string|用户名|
|`user:<id>:password`|string|密码|

|key|type|description|
|:-:|:-:|:-:|
|`user:<username>:id`|string|username到id的反向映射，登录和注册时作为唯一索引来用|

|key|type|description|
|:-:|:-:|:-:|
|`user:nextid`|string|user表的自增ID|
