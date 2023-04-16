package dto

type UserDTO struct {
	ID       uint   `json:"id" valid:"optional"`
	Username string `json:"username" valid:"required,type(string)"`
	Email    string `json:"email" valid:"required~Email should not be empty,email~Please provide a valid email,type(string)"`
	Password string `json:"password" valid:"required~Password cannot be empty,minstringlength(6)~Password cannot be less than 6 characters; type(string)"`
	Age      int    `json:"age" valid:"required~Age should not be empty,type(int)"`
}
