CREATE DATABASE IF NOT EXISTS earlGrey;
use earlGrey;

-- 何限目の情報を格納している
CREATE TABLE timers(
	time_no char(10) not null,
	s_time char(5) not null,
	e_time char(5)not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(time_no)
);

-- data
INSERT INTO timers(time_no, s_time, e_time)
  VALUES("1限目", "09:15", "10:45"),
        ("2限目", "11:00", "12:30"),
        ("3限目", "13:30", "15:00"),
        ("4限目", "15:15", "16:45"),
        ("5限目", "17:00", "18:30");

-- 先生たちの権限情報
CREATE TABLE permissions(
	per_no int auto_increment,
	permission char(10) not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(per_no)
);

-- data
INSERT INTO permissions(permission)
  VALUES("予約"),
        ("申請");

-- 先生の情報
CREATE TABLE teachers(
	id int auto_increment,
	teacher_name varchar(255) not null,
	password varchar(255) not null,
	per_no int not null,
	mail varchar(255) unique,
	updated_at datetime null,
	created_at datetime null,
	deleted_at datetime null,

	primary key(id),
	foreign key(per_no) references permissions(per_no) ON DELETE CASCADE ON UPDATE CASCADE

);

-- sampel data
INSERT INTO teachers(teacher_name, per_no, password, mail)
  VALUES("", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test1@gmail.com"),
  		("内山豊彦", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test2@gmail.com"),
        ("武次順平", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test3@gmail.com"),
		("小戎冴茄", 2, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test4@gmail.com"),
		("杉原宏", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test5@gmail.com"),
		("上村香代子", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test6@gmail.com");

INSERT INTO teachers(teacher_name, per_no, password, mail)
	VALUES("山本太", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test7@gmail.com"),
			("加藤昌", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test8@gmail.com"),
			("石田雄太", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test9@gmail.com"),
			("小出操", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test10@gmail.com"),
			("曽根国雄", 2, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test11@gmail.com"),
			("先生名前（仮）", 1, "$2a$12$r.Mj6LHidxSpvaIncoMY0OiKP3OgoywlkO9xydXNLQs/iF8G/Mhsi", "test12@gmail.com");

-- 予約申請の状態を格納している（承認など）
CREATE TABLE states(
	state_no int auto_increment,
	state_name char(20) not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(state_no)
);

-- data
INSERT INTO states(state_name)
  VALUES("承認済み"),
        ("承認待ち"),
        ("否認");

-- 教室の情報
CREATE TABLE rooms(
	id int(255) auto_increment unique,
	room_no varchar(4),
	memo varchar(255),
	is_detected boolean,
	updated_at datetime null,
	created_at datetime null,
	deleted_at datetime null,

	primary key(room_no)
);

-- data
INSERT INTO rooms(room_no, memo, is_detected, created_at)
  VALUES("1201", "コンセントプラグ：床", 0, '2020-08-01'),
		("1202", "コンセントプラグ：床", 0, '2020-08-01'),
		("1203", "コンセントプラグ：床", 0, '2020-08-01'),
		("1204", "コンセントプラグ：床", 0, '2020-08-01'),
        ("1205", "コンセントプラグ：床", 0, '2020-08-01'),
		("2301", "コンセントプラグ：机の上, ネットワーク機器あり", 0, '2020-08-01'),
		("4301", "コンセントプラグ：床", 0, '2020-08-01'),
		("3301", "コンセントプラグ：机の横", 0, '2020-08-01'),
		("2302", "コンセントプラグ：机の上", 0, '2020-08-01'),
		("4203", "コンセントプラグ：床", 0, '2020-08-01'),
        ("2031", "コンセントプラグ：机の上, ネットワーク機器あり", 0, '2020-08-01');


INSERT INTO rooms(room_no, memo, is_detected, created_at)
	VALUES("2405", "コンセントプラグ：机の横", 0, '2020-08-01'),
		  ("2402", "コンセントプラグ：床", 0, '2020-08-01'),
		  ("3502", "コンセントプラグ：机の上", 0, '2020-08-01'),
		  ("1403", "コンセントプラグ：机の上", 0, '2020-08-01'),
		  ("4202", "コンセントプラグ：床", 0, '2020-08-01'),
		  ("2303", "コンセントプラグ：床", 0, '2020-08-01'),
		  ("2501", "コンセントプラグ：床", 0, '2020-08-01');

-- 時間割り
CREATE TABLE timetables(
	No int auto_increment,
	room_no varchar(4) not null,
	subject_name varchar(255) not null,
	youbi char(3) not null,
	teacher_no int not null,
	time_no char(10) not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(No),
	foreign key(room_no) references rooms(room_no) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(teacher_no) references teachers(id) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(time_no) references timers(time_no) ON DELETE CASCADE ON UPDATE CASCADE
);

-- sample data
INSERT INTO timetables(room_no, subject_name, youbi, teacher_no, time_no)
  VALUES("4301", "セキュリティ演習_A", "Mon", 2, "1限目"),
				("4301", "セキュリティ演習_A", "Mon", 2, "2限目"),
				("1205", "ITシステム開発演習V", "Tue", 2, "1限目"),
				("1205", "ITシステム開発演習V", "Tue", 2, "2限目"),
				("2301", "ITゼミ演習", "Tue", 2, "3限目"),
				("2301", "ITゼミ演習", "Tue", 2, "4限目"),
				("3301", "就職対策", "Wed", 6, "3限目"),
				("2302", "システム設計演習", "Wed", 5, "3限目"),
				("2302", "システム設計演習", "Wed", 5, "4限目"),
				("2301", "ITゼミ演習", "Thu", 3, "1限目"),
				("2301", "ITゼミ演習", "Thu", 3, "2限目"),
				("4203", "AIシステム開発演習", "Thu", 3, "3限目"),
				("4203", "AIシステム開発演習", "Thu", 3, "4限目"),
				("1205", "ITシステム開発演習V", "Fri", 2, "3限目"),
        		("1205", "ITシステム開発演習V", "Fri", 2, "4限目"),
				("1205", "", "Fri", 1, "5限目"),
				("3301", "ハイプロフェッショナルゼミ", "Fri", 2, "5限目");

INSERT INTO timetables(room_no, subject_name, youbi, teacher_no, time_no)
	VALUES("2405", "就職対策Ⅱ", "Mon", 7, "1限目"),
		  ("4202", "セキュリティ演習_A", "Mon", 8, "2限目"),
		  ("4202", "セキュリティ演習_A", "Mon", 8, "3限目"),
		  ("1204", "ITシステム開発演習Ⅴ", "Tue", 2, "1限目"),
		  ("1204", "ITシステム開発演習Ⅴ", "Tue", 2, "2限目"),
		  ("3502", "ITゼミ演習Ⅲ", "Tue", 9, "3限目"),
		  ("3502", "ITゼミ演習Ⅲ", "Tue", 9, "4限目"),
		  ("2302", "システム設計演習", "Wed", 5, "1限目"),
		  ("2302", "システム設計演習", "Wed", 5, "2限目"),
		  ("3502", "ITゼミ演習Ⅲ", "Thu", 10, "3限目"),
		  ("3502", "ITゼミ演習Ⅲ", "Thu", 10, "4限目"),
		  ("1403", "AIシステム開発演習Ⅰ", "Fri", 3, "1限目"),
		  ("1403", "AIシステム開発演習Ⅰ", "Fri", 3, "2限目"),
		  ("1204", "ITシステム開発演習Ⅴ", "Tue", 2, "3限目"),
		  ("1204", "ITシステム開発演習Ⅴ", "Tue", 2, "4限目"),
		  ("1204", "", "Tue", 1, "5限目");


INSERT INTO timetables(room_no, subject_name, youbi, teacher_no, time_no)
	VALUES("2302", "ITシステム開発演習Ⅰ", "Mon", 12, "1限目"),
		  ("2302", "ITシステム開発演習Ⅰ", "Mon", 12, "2限目"),
		  ("2301", "Ciscoネットワーク演習Ⅰ", "Mon", 12, "3限目"),
		  ("2301", "Ciscoネットワーク演習Ⅰ", "Mon", 12, "4限目"),
		  ("2302", "ITシステム開発演習Ⅰ", "Fri", 12, "1限目"),
		  ("2302", "ITシステム開発演習Ⅰ", "Fri", 12, "2限目"),
		  ("2301", "Ciscoネットワーク演習Ⅰ", "Tue", 12, "1限目"),
		  ("2301", "Ciscoネットワーク演習Ⅰ", "Tue", 12, "2限目"),
		  ("2303", "ITシステム開発演習Ⅰ", "Mon", 12, "1限目"),
		  ("2303", "ITシステム開発演習Ⅰ", "Mon", 12, "2限目"),
		  ("2303", "ITシステム開発演習Ⅰ", "Fri", 12, "1限目"),
		  ("2303", "ITシステム開発演習Ⅰ", "Fri", 12, "2限目"),
		  ("2301", "Ciscoネットワーク演習Ⅰ", "Fri", 12, "3限目"),
		  ("2301", "Ciscoネットワーク演習Ⅰ", "Fri", 12, "4限目"),
		  ("2302", "AWSクラウド演習Ⅰ", "Thu", 12, "3限目"),
		  ("2302", "AWSクラウド演習Ⅰ", "Thu", 12, "4限目"),
		  ("2501", "AIシステム開発演習Ⅰ", "Thu", 12, "1限目"),
		  ("2501", "AIシステム開発演習Ⅰ", "Thu", 12, "2限目"),
		  ("2501", "AIシステム開発演習Ⅰ", "Thu", 12, "3限目");

-- 各教室の空のデータ挿入
INSERT INTO timetables(room_no, subject_name, youbi, teacher_no, time_no)
	VALUES("2301", "", "Mon", 1, "1限目"),
	      ("2301", "", "Mon", 1, "2限目"),
		  ("2301", "", "Mon", 1, "5限目"),
		  ("2302", "", "Mon", 1, "3限目"),
		  ("2302", "", "Mon", 1, "4限目"),
		  ("2302", "", "Mon", 1, "5限目"),
		  ("2303", "", "Mon", 1, "3限目"),
		  ("2303", "", "Mon", 1, "4限目"),
		  ("2303", "", "Mon", 1, "5限目"),
		  ("2301", "", "Tue", 1, "5限目"),
		  ("2302", "", "Tue", 1, "1限目"),
		  ("2302", "", "Tue", 1, "2限目"),
		  ("2302", "", "Tue", 1, "3限目"),
		  ("2302", "", "Tue", 1, "4限目"),
		  ("2302", "", "Tue", 1, "5限目"),
		  ("2303", "", "Tue", 1, "1限目"),
		  ("2303", "", "Tue", 1, "2限目"),
		  ("2303", "", "Tue", 1, "3限目"),
		  ("2303", "", "Tue", 1, "4限目"),
		  ("2303", "", "Tue", 1, "5限目"),
		  ("2301", "", "Wed", 1, "1限目"),
		  ("2301", "", "Wed", 1, "2限目"),
		  ("2301", "", "Wed", 1, "3限目"),
		  ("2301", "", "Wed", 1, "4限目"),
		  ("2301", "", "Wed", 1, "5限目"),
		  ("2302", "", "Wed", 1, "5限目"),
		  ("2303", "", "Wed", 1, "1限目"),
		  ("2303", "", "Wed", 1, "2限目"),
		  ("2303", "", "Wed", 1, "3限目"),
		  ("2303", "", "Wed", 1, "4限目"),
		  ("2303", "", "Wed", 1, "5限目"),
		  ("2301", "", "Thu", 1, "3限目"),
		  ("2301", "", "Thu", 1, "4限目"),
		  ("2301", "", "Thu", 1, "5限目"),
		  ("2302", "", "Thu", 1, "1限目"),
		  ("2302", "", "Thu", 1, "2限目"),
		  ("2302", "", "Thu", 1, "5限目"),
		  ("2303", "", "Thu", 1, "1限目"),
		  ("2303", "", "Thu", 1, "2限目"),
		  ("2303", "", "Thu", 1, "3限目"),
		  ("2303", "", "Thu", 1, "4限目"),
		  ("2303", "", "Thu", 1, "5限目"),
		  ("2301", "", "Fri", 1, "1限目"),
		  ("2301", "", "Fri", 1, "2限目"),
		  ("2301", "", "Fri", 1, "5限目"),
		  ("2302", "", "Fri", 1, "3限目"),
		  ("2302", "", "Fri", 1, "4限目"),
		  ("2302", "", "Fri", 1, "5限目"),
		  ("2303", "", "Fri", 1, "3限目"),
		  ("2303", "", "Fri", 1, "4限目"),
		  ("2303", "", "Fri", 1, "5限目");

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
	state_no int not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(rese_no),
	foreign key(teacher_no) references teachers(id) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(room_no) references rooms(room_no) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(state_no) references states(state_no) ON DELETE CASCADE ON UPDATE CASCADE
);

-- sample data
INSERT INTO reservations(teacher_no, room_no, rese_date, s_time, e_time, purpose, request_date, state_no)
  VALUES(1, "1204", "2022-06-01", "12:00", "13:00", "面談", "2022-05-27", 2),
	(2, "2301", "2022-07-07", "13:00", "15:00", "面談", "2022-07-01", 1),
	(7, "1204", "2022-06-01", "13:00", "15:00", "授業準備", "2022-05-26", 1);
