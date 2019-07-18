# 总结

### handler设计
handler -> validation(1.request, 2.user)-> business logic -> response                           
1、 data model                   
2、 error handler

```go
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}
```
这种是go原生处理json 的一种方法， 标记

### 数据库的设计
```mysql
create database if not exists video_server;
use video_server;

create table if not exists user (
  id         int(10)     not null primary key auto_increment,
  login_name varchar(64) not null,
  pwd        text        not null
)
  ENGINE = InnoDB
  default charset = utf8;

create table if not exists video_info (
  id            varchar(64) not null primary key,
  author_id     int(10),
  name          text,
  display_ctime text,
  create_time   datetime default current_timestamp()
)
  ENGINE = InnoDB
  default charset = utf8;

create table if not exists comments (
  id       varchar(64) not null primary key,
  video_id varchar(64),
  auto_id  int(10),
  content  text
)
  ENGINE = InnoDB
  default charset = utf8;

create table if not exists sessions (
  session_id varchar(255) not null primary key,
  TTL        varchar(255),
  login_name text
)
  ENGINE = InnoDB
  default charset = utf8;
```