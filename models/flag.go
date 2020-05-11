package models

import "time"

// Flag model
type Flag struct {
	UID              string    `json:"uid" firestore:"uid"`
	Address          string    `json:"address" firestore:"address"`
	AddressReference string    `json:"address_reference" firestore:"address_reference"`
	Description      string    `json:"description" firestore:"description"`
	SenderName       string    `json:"sender_name" firestore:"sender_name"`
	SenderPhotoURL   string    `json:"sender_photo_url" firestore:"sender_photo_url"`
	Text             string    `json:"text" firestore:"text"`
	Timestamp        time.Time `json:"timestamp" firestore:"timestamp"`
	Visibility       string    `json:"visibility" firestore:"visibility"`
}

// FlagFirestore model
type FlagFirestore struct {
	UID struct {
		Value string `json:"stringValue"`
	} `json:"uid"`
	Address struct {
		Value string `json:"stringValue"`
	} `json:"address"`
	AddressReference struct {
		Value string `json:"stringValue"`
	} `json:"address_reference"`
	Description struct {
		Value string `json:"stringValue"`
	} `json:"description"`
	MediaContent struct {
		Value struct {
			Field struct {
				DownloadURL struct {
					Value bool `json:"stringValue"`
				} `json:"download_url"`
				MimeType struct {
					Value string `json:"stringValue"`
				} `json:"mime_type"`
			} `json:"fields"`
		} `json:"mapValue"`
	} `json:"media_content"`
	Timestamp struct {
		Value string `json:"timestampValue"`
	} `json:"timestamp"`
	Visibility struct {
		Value string `json:"stringValue"`
	} `json:"visibility"`
}
