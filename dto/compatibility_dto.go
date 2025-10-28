package dto

type CompatibilityResponse struct {
	Score               float64  `json:"score"`
	MetNeeds            []string `json:"met_needs"`
	UnmetNeeds          []string `json:"unmet_needs"`
	TotalCandidateNeeds int      `json:"total_candidate_needs"`
}
