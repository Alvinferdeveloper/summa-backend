package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

func GenerateICSContent(interview models.Interview) string {
	// format time for ics
	const icsTimeFormat = "20060102T150405Z"

	startTime := interview.ScheduledAt.UTC().Format(icsTimeFormat)

	endTime := interview.ScheduledAt.UTC().Add(1 * time.Hour).Format(icsTimeFormat)
	timestamp := time.Now().UTC().Format(icsTimeFormat)

	var builder strings.Builder

	builder.WriteString("BEGIN:VCALENDAR\r\n")
	builder.WriteString("VERSION:2.0\r\n")
	builder.WriteString("PRODID:-//Summa//Job Interview//EN\r\n")
	builder.WriteString("BEGIN:VEVENT\r\n")
	builder.WriteString(fmt.Sprintf("UID:%d@summa.com\r\n", interview.ID))
	builder.WriteString(fmt.Sprintf("DTSTAMP:%s\r\n", timestamp))
	builder.WriteString(fmt.Sprintf("DTSTART:%s\r\n", startTime))
	builder.WriteString(fmt.Sprintf("DTEND:%s\r\n", endTime))
	builder.WriteString(fmt.Sprintf("SUMMARY:Entrevista para %s\r\n", interview.JobApplication.JobPost.Title))

	description := fmt.Sprintf("Entrevista con %s para el puesto de %s.\nFormato: %s.",
		interview.JobApplication.JobPost.Employer.CompanyName,
		interview.JobApplication.JobPost.Title,
		interview.Format)
	if interview.Notes != "" {
		description += fmt.Sprintf("\nNotas del empleador: %s", interview.Notes)
	}

	description = strings.ReplaceAll(description, "\n", "\\n")
	builder.WriteString(fmt.Sprintf("DESCRIPTION:%s\r\n", description))

	builder.WriteString("END:VEVENT\r\n")
	builder.WriteString("END:VCALENDAR\r\n")

	return builder.String()
}
