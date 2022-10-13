set global local_infile = 1;

CREATE DATABASE IF NOT EXISTS earlGrey;
use earlGrey;

-- 何限目の情報を格納している
CREATE TABLE timers(
	time_no int not null,
	s_time time not null,
	e_time time not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(time_no)
);

-- data
INSERT INTO timers(time_no, s_time, e_time)
  VALUES(1, '09:15', '10:45'),
        (2, '11:00', '12:30'),
        (3, '13:30', '15:00'),
        (4, '15:15', '16:45'),
        (5, '17:00', '18:30');

-- 先生の情報
CREATE TABLE teachers(
	id int auto_increment,
	teacher_name varchar(100) not null,
	password varchar(255) not null,
	mail varchar(255) unique,
	permisson varchar(10) default 'common',
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(id)

);

-- sampel data
INSERT INTO teachers(teacher_name, password, mail, permisson)
  VALUES("内山豊彦","uchiyama", "test@gmail.com", "common"),
				("小戎冴茄", "koebisu", "test1@gmail.com", "admin");


-- 教室の情報
CREATE TABLE rooms(
	room_no varchar(4) not null,
	outlet varchar(5) not null,
	lan boolean not null,
	is_detected boolean not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(room_no)
);

-- data
INSERT INTO rooms(room_no, outlet, lan, is_detected)
  VALUES("1201", 'down', true, false),
				("2302", 'up', true, false);

-- 時間割り
CREATE TABLE timetables(
	class_no int auto_increment,
	subject_name varchar(255) not null,
	youbi char(3) not null,
	time_no int not null,
	room_no varchar(4) not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(class_no),
	foreign key(room_no) references rooms(room_no) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(time_no) references timers(time_no) ON DELETE CASCADE ON UPDATE CASCADE
);

-- data
-- INSERT INTO timetables(subject_name, youbi, time_no, room_no)
--   VALUES("ITゼミ演習", "Tue", 1, "2302"),
-- 				("ITゼミ演習", "Tue", 2, "1201"),
-- 				("ITゼミ演習", "Tue", 3, "2302"),
-- 				("システム設計演習", "Tue", 4, "2302");

-- 予約
CREATE TABLE reservations(
	rese_no int auto_increment,
	teacher_no int not null,
	room_no varchar(4) not null,
	rese_date date not null,
	s_time time not null,
	e_time time not null,
	purpose varchar(150) not null,
	request_date date not null,
	request_state varchar(5) not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(rese_no),
	foreign key(teacher_no) references teachers(id) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(room_no) references rooms(room_no) ON DELETE CASCADE ON UPDATE CASCADE
);

-- sample data
INSERT INTO reservations(teacher_no, room_no, rese_date, s_time, e_time, purpose, request_date, request_state)
  VALUES(1, "1201", '2022-06-01', "12:00", "13:00", "面談", "2022-05-27", "wait"),
	(1, "2302", "2022-07-07", "13:00", "15:00", "面談", "2022-07-01", "wait"),
	(1, "1201", "2022-06-01", "13:00", "15:00", "授業準備", "2022-05-26", "wait");