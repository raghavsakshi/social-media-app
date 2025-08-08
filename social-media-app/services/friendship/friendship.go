package friendship

import (
	"context"
	"encoding/json"
	"fmt"
	"social-media-app/internals/cache"
	"social-media-app/internals/dto"
	"social-media-app/models/friendship"
	"time"

	"github.com/google/uuid"
)
type Friends struct {
	UserID   uuid.UUID
	FriendsID uuid.UUID

	Friends  *dto.Friends

	AllFriends []*dto.AllFriends
}


func New() *Friends {
	return &Friends{}
}


func (f *Friends) Create(ctx context.Context) error {
    m := friendships.New()
    m.Friends = f.Friends
    err := m.Create(ctx)
    f.Friends.UpdatedAt = nil  
    return err
}

func (f *Friends) GetAll(ctx context.Context) error {
    val, err := cache.Client().Get(ctx, f.UserID.String()).Result()
    if val != "" && err == nil {
        json.Unmarshal([]byte(val), &f.AllFriends)
        return nil
    }

    m := friendships.New()
    m.UserID = f.UserID
    m.Get(ctx)

    f.AllFriends = m.AllFriends
    b, _ := json.Marshal(f.AllFriends)
    if err := cache.Client().Set(ctx, f.UserID.String(), b, 24*time.Hour).Err(); err != nil {
        fmt.Println("Error setting cache for friends:", err)
    }
    return nil
}

func (u *Friends) Delete(ctx context.Context) error {
    m := friendships.New()
    m.UserID = u.UserID
    m.FriendID = u.FriendsID
    if err := m.Delete(ctx); err != nil {
        return err
    }
    return nil
}