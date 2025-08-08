package friendship

import (
	"fmt"

	"social-media-app/internals/cache"
	"social-media-app/internals/dto"
	"social-media-app/internals/validator"
	"social-media-app/models/friendship"
	"social-media-app/services/users"
	"github.com/google/uuid"
	 "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	
)

func Add(c *fiber.Ctx) error{
	ctx:=c.UserContext()
	var friend dto.FriendsCreate
	if err := c.BodyParser(&friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}
	if err := validator.Payload(&friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user data")
	}

	us :=users.New()
us.User =&dto.User{}
    us.User.ID = friend.UserID
    if err := us.Get(ctx); err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON("user not found!")
        }
        return c.Status(fiber.StatusInternalServerError).JSON("internal server error!")
    }

    friendUser := users.New()
    friendUser.User = &dto.User{}
    friendUser.User.ID = friend.FriendID
    if err := friendUser.Get(ctx); err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON("user not found!")
        }
        return c.Status(fiber.StatusInternalServerError).JSON("internal server error!")
    }

    fs := friendships.New()
    fs.UserID = friend.UserID
    fs.FriendID = friend.FriendID
    fs.Create(ctx)
    if err := cache.Client().Del(ctx, fs.UserID.String()).Err(); err != nil {
        fmt.Printf("Error invalidating cache : %v\n", err)
    }
    return c.Status(fiber.StatusCreated).JSON(fs.Friends) 
}

func Get(c *fiber.Ctx) error {
	ctx:= c.UserContext()
	id :=c.Params("id")
	userID,err:=uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}
	
us :=users.New()
us.User =&dto.User{}
us.User.ID= userID
if err :=us.Get(ctx);err !=nil{
	if err==gorm.ErrRecordNotFound{
return  c.Status(fiber.StatusNotFound).JSON("user not found!")
	}
	return  c.Status(fiber.StatusInternalServerError).JSON("internal server error!")
}

fs := friendships.New()
fs.UserID =userID
fs.GetAll(ctx)
return c.Status(fiber.StatusOK).JSON(fs.AllFriends)
}


func Delete(c *fiber.Ctx) error {
	ctx:= c.UserContext()
	id :=c.Params("id")
	userID,err:=uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}
	fid:=c.Query("f_id")
	friendID,err :=uuid.Parse(fid)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}
	
us :=users.New()
us.User =&dto.User{}
us.User.ID= userID
if err :=us.Get(ctx);err !=nil{
	if err==gorm.ErrRecordNotFound{
return  c.Status(fiber.StatusNotFound).JSON("user not found!")
	}
		return  c.Status(fiber.StatusInternalServerError).JSON("internal server error!")
}


us.User =&dto.User{}
us.User.ID= friendID
if err :=us.Get(ctx);err !=nil{
	if err==gorm.ErrRecordNotFound{
return  c.Status(fiber.StatusNotFound).JSON("user not found!")
	}
	return  c.Status(fiber.StatusInternalServerError).JSON("internal server error!")
}

fs := friendships.New()
fs.UserID =userID
fs.FriendID =friendID
fs.Delete(ctx)
if err := cache.Client().Del(ctx, fs.UserID.String()).Err(); err != nil {
	fmt.Println("Error invalidating cache : %v\n", err)
}
return c.Status(fiber.StatusNoContent).JSON("friendship deleted")

}

