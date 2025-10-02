package dto

type EmployerLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type OnboardingRequest struct {
	FirstName         string `json:"first_name" binding:"required"`
	LastName          string `json:"last_name" binding:"required"`
	DisabilityTypeIDs []uint `json:"disability_type_ids" binding:"required,min=1"`
}
