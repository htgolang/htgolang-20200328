create database if not exists cmdb default charset utf8mb4;
create table if not exists user(
    id bigint primary key auto_increment,
    staff_id varchar(32) not null default '',
    name varchar(64) not null default '',
    nickname varchar(64) not null default '',
    password varchar(1024) not null default '',
    gender int not null default 0 comment '0: 女, 1: 男',
    tel varchar(32) not null default '',
    email varchar(64) not null default '',
    addr varchar(128) not null default '',
    department varchar(128) not null default '',
    status int not null default 0 comment '0: 正常, 1: 锁定, 2: 离职',
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
)engine=innodb default charset utf8mb4;

insert into user(staff_id, name, nickname, password, gender, tel, email, addr, department, status, created_at, updated_at) values
("K00001", "kk", "kk", "$2a$10$7Sm8HKryfpStF6Or2AE78eKHaChmkO8PO86O8kZN6cHOHawZcE33S", 1, "15222222222", "iamkk@outlook.com", "陕西省西安市", "研发一部", 0, now(), now());