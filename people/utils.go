package people

import (
	"context"

	"github.com/dpolimeni/fiber_app/ent"
	"github.com/dpolimeni/fiber_app/ent/user"
	_ "github.com/lib/pq"
)

func CheckSuperuser(DbClient *ent.Client, username string) bool {
	user, err := DbClient.User.Query().Where(user.UsernameEQ(username)).First(context.Background())
	if err != nil {
		return false
	}
	if user.IsAdmin {
		return true
	} else {
		return false
	}
}
