package usernodejs

type User struct {
	ID          int      `json:"id"`
	Email       string   `json:"email"`
	Username    string   `json:"username"`
	Level       string   `json:"level"`
	IsActivated bool     `json:"is_activated"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	IsDeleted   bool     `json:"is_deleted"`
	Document    Document `json:"document"`
}

type Document struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Fullname  string `json:"fullname"`
	Nik       string `json:"nik"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`
}

type ResponseAllUser struct {
	StatusCode int    `json:"status_code"`
	Result     string `json:"result"`
	Message    string `json:"message"`
	Data       []User `json:"data"`
}

type ResponseUserById struct {
	StatusCode int    `json:"status_code"`
	Result     string `json:"result"`
	Message    string `json:"message"`
	Data       User   `json:"data"`
}

func UserByteToResponse(user User) User {
	return User{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		Level:       user.Level,
		IsActivated: user.IsActivated,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		IsDeleted:   user.IsDeleted,
		Document:    DocumentByteToResponse(user.Document),
	}
}

func DocumentByteToResponse(doc Document) Document {
	return Document{
		ID:        doc.ID,
		UserID:    doc.UserID,
		Fullname:  doc.Fullname,
		Nik:       doc.Nik,
		Phone:     doc.Phone,
		Gender:    doc.Gender,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
		IsDeleted: doc.IsDeleted,
	}
}