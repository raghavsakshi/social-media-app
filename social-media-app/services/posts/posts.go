package posts

import (
	"github.com/google/uuid"
	"time"
	"social-media-app/models/posts"   

	"context"
	"social-media-app/internals/dto"
	
)

type Posts struct {
    ID        uuid.UUID   `json:"id"`
    UserID    uuid.UUID   `json:"user_id"`
    Content   string      `json:"content"`
    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
    Post      *dto.Post   
    Posts     []*dto.Post 
}


func New() *Posts {
	return &Posts{}
}

func (p *Posts) Create(ctx context.Context) error {
	m := posts.New()
	m.Post = p.Post
err :=	m.Create(ctx)
	p.Post.UpdatedAt = nil
	return err
}

func (p *Posts) GetAll(ctx context.Context) error {
	m := posts.New()
	m.UserID = p.UserID
	m.Posts = p.Posts
	err := m.Get(ctx)
	return err
}

func (p *Posts) Delete(ctx context.Context) error {

	m := posts.New()
	m.ID = p.ID
	m.UserID = p.UserID
	if err := m.Delete(ctx); err != nil {
		return err
	}
	return nil
}