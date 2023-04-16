package dto

type UserDTO struct {
	ID       uint   `json:"id" form:"id" valid:"optional"`
	Username string `json:"username" form:"username" valid:"required,type(string)"`
	Email    string `json:"email" form:"email" valid:"required~Email should not be empty,email~Please provide a valid email,type(string)"`
	Password string `json:"password" form:"password" valid:"required~Password cannot be empty,minstringlength(6)~Password cannot be less than 6 characters; type(string)"`
	Age      int    `json:"age" form:"age" valid:"required~Age should not be empty,type(int)"`
}
