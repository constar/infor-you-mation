# Redis 数据库设计

## 数据结构

### Topic

|key|type|
|:-:|:-:|
|topic:<id>:words|set|
|topic:<id>:joblist|zset|
|topic:<id>:name|string|

### Job

```
job:<id>:content
job:<id>:url
job:<id>:title
job:<id>:urlmd5
```


