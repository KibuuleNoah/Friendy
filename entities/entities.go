package entities

type Friend struct{
	Id 			 string `db:"id" json:"id"`
	FullName string `db:"fullname" json:"fullname"`
	Tel 		 string `db:"tel" json:"tel"`
	Desc 		 string `db:"desc" json:"desc"`
	FirstMetOn string `db:"first_met_on" json:"first_met_on"`
	MetPlace string `db:"met_place" json:"met_place"`
	Tags string `db:"tags" json:"tags"`
	UserId string `db:"user_id" json:"user_id"`
}
