package models

import "time"

// Flag model
type Flag struct {
	UID              string           `json:"uid" firestore:"uid"`
	Address          string           `json:"address" firestore:"address"`
	AddressReference string           `json:"address_reference" firestore:"address_reference"`
	Description      string           `json:"description" firestore:"description"`
	MediaContent     FlagMediaContent `json:"media_content" firestore:"media_content"`
	SenderName       string           `json:"sender_name" firestore:"sender_name"`
	SenderPhotoURL   string           `json:"sender_photo_url" firestore:"sender_photo_url"`
	Text             string           `json:"text" firestore:"text"`
	Timestamp        time.Time        `json:"timestamp" firestore:"timestamp"`
	Visibility       string           `json:"visibility" firestore:"visibility"`
}

// FlagMediaContent model
type FlagMediaContent struct {
	DownloadURL string `json:"download_url" firestore:"download_url"`
	MimeType    string `json:"mime_type" firestore:"mime_type"`
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
					Value string `json:"stringValue"`
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

// ParseFlag convert FlagFirestore to Flag
func (f *FlagFirestore) ParseFlag() Flag {
	flag := Flag{
		UID:              f.UID.Value,
		Address:          f.Address.Value,
		AddressReference: f.AddressReference.Value,
		Description:      f.Description.Value,
		MediaContent: FlagMediaContent{
			DownloadURL: f.MediaContent.Value.Field.DownloadURL.Value,
			MimeType:    f.MediaContent.Value.Field.MimeType.Value,
		},
		Visibility: f.Visibility.Value,
	}
	return flag
}
