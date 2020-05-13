package banderablancafunctions

import (
	"banderablancafunctions/models"
	"banderablancafunctions/repository/comment"
	"banderablancafunctions/repository/flag"
	"banderablancafunctions/repository/notification"
	"banderablancafunctions/utils"
	"context"
	"log"
)

// CreateNotifications trigger to create a notification to user that created a white flag
func CreateNotifications(ctx context.Context, e models.EventComment) error {
	path := e.Value.Name
	flagID := utils.GetParams(path, "comments")
	flagComment := e.Value.Fields.ParseComment()

	flagNotification := map[string]interface{}{
		"sender_name":      flagComment.SenderName,
		"sender_photo_url": flagComment.SenderPhotoURL,
		"message":          flagComment.Text,
		"uid":              flagComment.UID,
		"flag_id":          flagID,
	}

	flagSelected := <-flag.GetByID(flagID)
	commenters := <-comment.GetUsers(flagID)

	// Create Notifications
	for _, userID := range commenters {
		if userID != flagComment.UID && userID != flagSelected.UID {
			err := <-notification.Save(userID, flagNotification)
			if err != nil {
				log.Fatalf("Error save notification: %v", err)
				return err
			}
		}
	}

	if flagComment.UID != flagSelected.UID {
		err := <-notification.Save(flagSelected.UID, flagNotification)
		if err != nil {
			log.Fatalf("Error send notification to flag creator: %v", err)
			return err
		}
	}

	err := <-notification.Send(flagID, &flagComment)
	return err
}

// CreateImageMarked trigger to create a image with logo
func CreateImageMarked(ctx context.Context, e models.EventFlag) error {
	path := e.Value.Name
	flagID := utils.GetParams(path, "flags")
	flagData := e.Value.Fields.ParseFlag()

	imageMarkedURL, err := flag.UploadImageMarked(flagID, flagData)
	if err != nil {
		log.Fatalf("Error to upload flag image marked: %v", err)
		return err
	}

	err = <-flag.Update(flagID, map[string]interface{}{"image_marked": imageMarkedURL})
	if err != nil {
		log.Fatalf("Error to update flag: %v", err)
		return err
	}

	return nil
}
