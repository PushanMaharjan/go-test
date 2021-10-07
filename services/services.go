package services

import (
	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewGmailService),
	fx.Provide(NewNotificationService),
	fx.Provide(NewRoleService),
)
