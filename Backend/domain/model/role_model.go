package model

type Role struct {
	Roleid   int `json:"id"`
	Rolename string
	Leader   bool
	Modules  []*struct {
		Modulerolid       int
		Moduleid          int
		Modulepermissions []struct {
			Modulepermissionid int
		}
		Permissions []int
	}
}
