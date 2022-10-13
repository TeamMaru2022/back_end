package json

type JsonRequestUser struct {
	UserMail     string `json:"mail"`
	UserPassword string `json:"pass"`
	UserAge      uint8  `json:"age"`
}

type JsonRequestUserLogin struct {
	UserMail     string `json:"mail"`
	UserPassword string `json:"pass"`
}

type UpdateUserJson struct {
	Mail        string `json:"mail"`
	Password    string `json:"pass"`
	NewMail     string `json:"new-mail"`
	NewPassword string `json:"new-pass"`
}

type DeleteUserJson struct {
	Mail     string `json:"mail"`
	Password string `json:"pass"`
}

type SampleValidationJson struct {
	Token string `json:"token"`
}

type SensorInfoJson struct {
	RoomNo     string `json:"room_no"`
	IsDetected bool   `json:"is_detected"`
}

type RegTeacher struct {
	TeacherName string `json:"teacher_name"`
	Password    string `json:"password"`
	TeacherMail string `json:"mail"`
}

type TeacherLogin struct {
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type JsonReservation struct{
	TeacherNo	int			`json:"teacher_no"`
	RoomNo		string	`json:"room_no"`
	ReseDate	string	`json:"rese_date"`
	StartT		string	`json:"s_time"`
	EndT			string	`json:"e_time"`
	StateNo		int			`json:"state_no"`
}