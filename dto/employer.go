package dto

type EmployerRegisterRequest struct {
	CompanyName    string `json:"company_name" binding:"required,min=3"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	PhoneNumber    string `json:"phone_number"`
	Country        string `json:"country"`
	FoundationDate string `json:"foundation_date"` // String to parse into time.Time
	Industry       string `json:"industry"`
	Size           string `json:"size"`
	Description    string `json:"description"`
	Address        string `json:"address"`
	Website        string `json:"website"`
}
