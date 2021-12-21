package models

import (
	"context"
	"time"

	"github.com/mises-id/sns-storagesvc/app/models/enum"
	"github.com/mises-id/sns-storagesvc/lib/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Attachment struct {
	ID        uint64        `bson:"_id"`
	Filename  string        `bson:"filename,omitempty"`
	FileType  enum.FileType `bson:"file_type"`
	CreatedAt time.Time     `bson:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty"`
	UserType  enum.UserType `bson:"user_type"`
	UserId    uint64        `bson:"user_id"`
}

func (a *Attachment) BeforeCreate(ctx context.Context) error {
	var err error
	a.ID, err = getNextSeq(ctx, "attachmentid")
	if err != nil {
		return err
	}
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Attachment) Create(ctx context.Context, in *Attachment) (*Attachment, error) {

	if err := in.BeforeCreate(ctx); err != nil {
		return nil, err
	}
	_, err := db.DB().Collection("attachments").InsertOne(ctx, in)
	return in, err
}

func FindAttachmentMap(ctx context.Context, ids []uint64) (map[uint64]*Attachment, error) {
	attachments := make([]*Attachment, 0)
	cursor, err := db.DB().Collection("attachments").Find(ctx,
		bson.M{
			"_id": bson.M{"$in": ids},
		})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &attachments); err != nil {
		return nil, err
	}
	result := make(map[uint64]*Attachment)
	for _, attachment := range attachments {
		result[attachment.ID] = attachment
	}
	return result, nil
}
