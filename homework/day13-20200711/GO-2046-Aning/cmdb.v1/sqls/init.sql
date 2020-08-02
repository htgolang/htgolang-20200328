create database if not exists cmdb default charset utf8mb4;

create table if not exists user(
  id bigint primary key auto_increment,
  staff_id varchar(32) not null default '',
  name varchar(64) not null default '',
  nickname varchar(64) not null default '',
  password varchar(1024) not null default '',
  gender int not null default 0,
  tel varchar(32) not null default '',
  email varchar(64) not null default '',
  addr varchar(128) not null default '',
  department varchar(128) not null default '',
  status int null default 0,
  create_at datetime not null,
  updated_at datetime not null,
  deleted_at datetime
  )engine=innodb default charset utf8mb4;

insert into user(staff_id,name,nickname,password,gender,tel,email,addr,department,status,create_at,updated_at) values
("k001","kk","kk",md5("123"),1,"110","a@bc.com","beijingxi","CEO",0,now(),now());
insert into user(staff_id,name,nickname,password,gender,tel,email,addr,department,status,create_at,updated_at) values
("k002","kk2","kk2",md5("123"),1,"120","a2@bc.com","beijingxi2","CEO2",0,now(),now());
insert into user(staff_id,name,nickname,password,gender,tel,email,addr,department,status,create_at,updated_at) values
("k003","kk3","kk3",md5("123"),1,"130","a3@bc.com","beijingxi3","CEO3",0,now(),now());
insert into user(staff_id,name,nickname,password,gender,tel,email,addr,department,status,create_at,updated_at) values
("k004","kk4","kk4",md5("123"),1,"140","a4@bc.com","beijingxi4","CEO4",0,now(),now());
insert into user(staff_id,name,nickname,password,gender,tel,email,addr,department,status,create_at,updated_at) values
("k005","kk5","kk5",md5("123"),1,"150","a5@bc.com","beijingxi5","CEO5",0,now(),now());
insert into user(staff_id,name,nickname,password,gender,tel,email,addr,department,status,create_at,updated_at) values
("k006","kk6","kk6",md5("123"),1,"160","a6@bc.com","beijingxi6","CEO6",0,now(),now());