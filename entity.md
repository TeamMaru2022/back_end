# DBのEntity

## データベース名

earlGray

## テーブル

update_at, create_at, delete_atはすべてのテーブルに存在する。

データ型：datetime

### timetable(時間割)
- No
  - primary key
  - int
  - auto_increment
  - 時間割を特定するための値

- class
  - char(5)
  - not null
  - クラスを指定

- roomNo
  - char(4)
  - not null
  - foreign key：rooms表のroomNo
  - 教室の番号

- name
  - varchar(40)
  - not null
  - 授業名

- youbi
  - char(3)
  - not null
  - 曜日
  - 書き方：英語表記の先頭3文字(例->Mon)

- teacherNo
  - int
  - not null
  - foreign key：teachers表のteacherNo
  - 先生の番号

- timeNo
  - char(3)
  - not null
  - foreign key：timer表のtimeNo
  - ○限目の表示

### reservation(予約)

- reseNo
  - primary key
  - int
  - auto_increment
  - 予約を特定するための値

- teacherNo
  - int
  - not null
  - foreign key：teachers表のteacherNo
  - 予約申請した先生
- roomNo
  - char(4)
  - not null
  - foreign key：rooms表のroomNo
  - 予約したい教室番号
- reseDate
  - date
  - not null
  - 予約したい日程
- sTime
  - time
  - not null
  - 開始時間
- eTime
  - time
  - not null
  - 終了時間
- purpose
  - varchar(150)
  - not null
  - 予約理由
- requestDate
    - date
    - not null
    - 予約申請をした日時
- stateNo
  - int
  - not null
  - foreign key：state表のstateNo
  - 予約の状態を表す（予約権限を持った先生が承認したかどうか）

### timer(時間)
※何限目と決まっている時間のみ(時間割作成で使用する)
- timeNo
  - char(3)
  - not null
  - ○限目
- sTime
  - char(5)
  - not null
  - 開始時間
- eTime
  - char(5)
  - not null
  - 終了時間

### permission(権限)
※先生が予約する時の権限
- perNo
  - primary key
  - int
  - auto_increment
  - 権限を特定するための値
- permission
  - char(2)
  - not null
  - 権限

### teachers(先生)
- teacherNo
  - primary key
  - int
  - auto_increment
  - 先生を特定するための値
- name
  - varchar(20)
  - not null
  - 先生の名前
- perNo
  - int
  - not null
  - foreign key：parmission表のperNo
  - 先生の権限を表す

### state(予約申請の状態)
- stateNo
  - primary key
  - int
  - auto_increment
  - 予約申請の状態を特定するための値
- stateName
  - char(4)
  - not null
  - 申請の状態

### rooms(教室)
- roomNo
  - primary key
  - char(4)
  - 教室番号
- memo
  - varchar(255)
  - 教室設備などを書き留めるメモ欄