package auth

import (
	"context"

	"github.com/cahyobintarum/MayGram/pkg/domain/user"
)

type AuthRepo interface {
	LoginUser(ctx context.Context, username string) (result user.User, err error)
}
