package notifications

import (
	"context"
	"fmt"
	"log"
	"social-media-app/internals/dto"
	friendships "social-media-app/models/friendship"
	"social-media-app/models/users"
	"sync"
"github.com/google/uuid"
)

var Store  map[uuid.UUID] chan string

var mu sync.Mutex

func InitNotificationsSystem() {
	Store = make(map[uuid.UUID]chan string)
}
func Register(userID uuid.UUID) chan string {
	mu.Lock()
	defer mu.Unlock()
	if _,ok :=Store[userID]; !ok {
		Store[userID] = make(chan string)
	}
	return Store[userID]
}
func ListenForNotifications(ctx context.Context, userID uuid.UUID) <-chan string {
	mu.Lock()
	channel, ok := Store[userID]
	mu.Unlock()
	if !ok {
		fmt.Printf("No Notifications channel registered for user %v", userID)
		return nil
	}   

	us :=users.New()
us.User =&dto.User{}
us.User.ID = userID
if err :=us.Get(ctx);err !=nil{
	return   nil

}


	for {
		select {
		case message := <-channel:
			// Handle the incoming message
			fmt.Printf("Hey,%v you have a new notification : %v\n", us.User.Name,message)

		case <-ctx.Done():
			// Context was cancelled, exit the goroutine
			fmt.Printf("Stopping notification listener for user %v\n", userID)
			return  nil
		}
	}
}

func NotifyUsers(ctx context.Context, userID uuid.UUID, msg string) {
    // get all friends
    fs := friendships.New()
    fs.UserID = userID
    fs.GetAll(ctx)
    mu.Lock()
    defer mu.Unlock()
    for _, f := range fs.AllFriends {
        if ch, ok := Store[f.FriendID]; ok {
            go func(ch chan string) {
                ch <- msg
            }(ch)
        }
    }
}

func Hydrate() {
	ctx := context.Background()
	us :=users.New()
us.User =&dto.User{}
if err :=us.GetAll(ctx);err !=nil{
	
	log.Fatalf("Internal error :%v ",err)
}
for _, u:= range us.AllUsers {
	Register (u.ID)
	go ListenForNotifications(ctx,us.ID)
}

}