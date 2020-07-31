package forms

type FormUser struct {
	ID         int    `form:"id"`
	StaffID    string `form:"staff_id"`
	Name       string `form:"name"`
	Nickname   string `form:"nickname"`
	Password   string `form:"password"`
	Gender     int    `form:"gender"`
	Tel        string `form:"tel"`
	Addr       string `form:"addr"`
	Email      string `form:"email"`
	Department string `form:"department"`
	Status     int    `form:"status"`
	CreatedAt  string `form:"created_at" `
	UpdatedAt  string `form:"updated_at" `
	DeletedAt  string `form:"deleted_at" `
	IsAdmin    int    `form:"isAdmin"`
}
