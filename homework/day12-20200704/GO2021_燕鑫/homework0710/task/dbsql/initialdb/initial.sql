use go_todolist;

create table task
(
    Id        int auto_increment,
    Name      varchar(20),
    StartTime varchar(50),
    EndTime   varchar(50),
    Status    varchar(20),
    User      varchar(20),
    primary key (id)
);

create table user
(
    Id         int auto_increment,
    Username   varchar(20),
    Password   varchar(255),
    Salt       varchar(255),
    CreateTime varchar(50),
    UpdateTime varchar(50),
    primary key (id)
);


REPLACE into task (Id, Name, StartTime, EndTime, Status, User)
values (1, "task1", "2020-04-18 18:56:00", "", "created", "yanxin1"),
       (2, "task2", "2020-04-18 18:57:00", "", "running", "yanxin1"),
       (3, "task3", "2020-04-18 18:58:00", "2020-04-18 19:02:00", "finished", "yanxin"),
       (4, "task4", "2020-04-18 19:01:00", "", "paused", "yanxin"),
       (5, "task5", "2020-04-18 18:56:00", "", "created", "yanxin2"),
       (6, "task6", "2020-04-18 18:57:00", "", "running", "yanxin"),
       (7, "task7", "2020-04-18 18:58:00", "2020-04-18 19:02:00", "finished", "yanxin3"),
       (8, "task8", "2020-04-18 19:01:00", "", "paused", "yanxin");

REPLACE into user (Id, Username, Password, Salt, CreateTime, UpdateTime)
values (1, "yanxin", "08c5f8701508dde64c281f83f4f47b8d", "ilEQKCsjIg", "2020-04-26 23:40:17", "2020-04-26 23:40:17"),
       (2, "yanxin1", "ad166203e2a5458156b68822dd1d187f", "upjlbBkIYW", "2020-04-26 23:46:39", "2020-04-26 23:46:39"),
       (3, "yanxin2", "639b0fa9b4cc86d3f1888fd0e6f02705", "XosEwpSTVh", "2020-04-27 00:48:10", "2020-04-27 00:48:10");