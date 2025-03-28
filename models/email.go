package models

import (
	"time"
)

type Message struct {
	ID  string `json:"id"`
	UID uint32 `json:"uid"`

	Subject         string   `json:"subject"`
	From            string   `json:"from"`
	FromName        string   `json:"from_name"`
	To              string   `json:"to"`
	Date            string   `json:"date"`
	AttachmentNames []string `json:"attachment_names"` // ✅ Store attachment names
}

// Attachment represents an email attachment
type Attachment struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	MIMEType string `json:"mimeType"`
}

// EmailRecord represents a stored email
type EmailRecord struct {
	Email string `json:"email"`
	Name  string `json:"name"` // Added recipient name field
}

// OTP represents an OTP and its expiration
type OTP struct {
	Code       string    `json:"code"`
	Expiration time.Time `json:"expiration"`
}

// UserSession holds session information for logged-in users
type UserSession struct {
	Name      string    `json:"name"` // Add this field
	Email     string    `json:"email"`
	LoginTime time.Time `json:"login_time"`
}
type OpdModel struct {
	PatientName  string
	DoctorName   string
	OPDDate      string
	OPDNotes     string
	Prescription string
	FollowupDate string
	CreatedOn    string
	GeneratedOn  string
}
