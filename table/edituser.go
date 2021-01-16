package table

type EditedUser struct {
	TableName struct{} `sql:"edited_user_details" json:"-"`
	Id        string   `param:"id"`
	Phone     string   `param:"phone"`
	FirstName string   `param:"first_name"`
	LastName  string   `param:"last_name"`
	Role      string   `param:"role"`
	Email     string   `param:"email"`
	UserId    string   `param:"userId"`
	AdminId   string   `param:"adminId"`
}


