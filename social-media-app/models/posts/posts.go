package posts

import (
	"social-media-app/internals/dto"
	"social-media-app/models/users"
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
			"social-media-app/internals/database"
			
	"context"
	"fmt"
)
type Posts struct  {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Content string     `json:"content"`
	UserID uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User users.Users  `gorm:"foreignKey:UserID;references:ID" json:"-"`
	 Post *dto.Post   `gorm:"-"`
	 Posts []*dto.Post    `gorm:"-"`

}

func New() *Posts{
	return &Posts{}
}


func (p *Posts) Create(ctx context.Context) error {
	if err := database.Client().Table("posts").Create(&p).Error; err != nil {
		fmt.Printf("Unable to create post : %v", err)
		return err
	}
	return nil
}

func (p *Posts) Get(ctx context.Context) error {
	if err := database.Client().Table("posts").Where("user_id=?", p.UserID).Find(&p.Posts).Error; err != nil {
		fmt.Printf("Unable to get posts : %v", err)
		return err
	}
	return nil
}
func (P *Posts) Delete(ctx context.Context) error {
	if err := database.Client().
		Where("user_id=? AND _id=?", P.UserID, P.ID).

		Delete(P).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
		fmt.Printf("Error getting user : %v\n", err)
		return err
	}
}

	return nil
}