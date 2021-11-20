package model

type Login struct {
	Email    string `json:"email" example:"example@grupogit.com"`
	Password string `json:"password" example:"example15"`
}
type LoginMovil struct {
	Email      string `json:"email" example:"example@grupogit.com"`
	Password   string `json:"password" example:"example15"`
	TokenMovil string `json:"tokenmovil" example:""`
}

type Authenticate struct {
	Userid         int    `json:"id" example:"1"`
	Resourcedataid int    `json:"resourcedataid" example:"1"`
	Roleid         int    `json:"roleid" example:"2"`
	Rolename       string `json:rolename example:"internship"`
	Status         bool   `json:status example:"true"`
}
type AuthenticateMovil struct {
	Userid         int  `json:"id" example:"1"`
	Resourcedataid int  `json:"resourcedataid" example:"1"`
	Roleid         int  `json:"roleid" example:"2"`
	Status         bool `json:status example:"true"`
}

type RecoveryPassword struct {
	Userid    int    `json:"-"`
	Email     string `json:"email" example:"example@grupogit.com"`
	Firstname string `json:"-"`
	Lastname  string `json:"-"`
}

type ChangePassword struct {
	Hashuserid  string `example:"8b459f91bacc7157222c4c1894c7229ae2d58eec745084bfdf1aa87b65"`
	Newpassword string `example:"ex@ample16"`
}
