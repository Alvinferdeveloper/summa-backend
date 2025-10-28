package dto

type DisabilityInsight struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type DashboardStats struct {
	ActiveJobs      int64 `json:"active_jobs"`
	TotalApplicants int64 `json:"total_applicants"`
	NewApplicants7d int64 `json:"new_applicants_7d"`
}

type PipelineStep struct {
	Status string `json:"status"`
	Count  int64  `json:"count"`
}

type SkillInsight struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type LocationInsight struct {
	Location  string  `json:"location"`
	Count     int64   `json:"count"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
