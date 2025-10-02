package dto

type CreateJobPostRequest struct {
	Title                 string   `json:"title" binding:"required"`
	Location              string   `json:"location" binding:"required"`
	WorkModel             string   `json:"workModel" binding:"required"`
	ContractType          string   `json:"contractType" binding:"required"`
	Description           string   `json:"description" binding:"required"`
	Responsibilities      string   `json:"responsibilities" binding:"required"`
	Requirements          string   `json:"requirements" binding:"required"`
	AccessibilityFeatures []string `json:"accessibilityFeatures"`
}
