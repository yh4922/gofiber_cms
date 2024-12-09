package ctx

import (
	"github.com/gofiber/fiber/v2"
)

type ManageUserLoginReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PasswordKey string `json:"password_key"`
	Code        string `json:"code"`
	CodeKey     string `json:"code_key"`
}

func ManageUserLogin(c *fiber.Ctx) error {
	req := new(ManageUserLoginReq)
	c.BodyParser(req)
	return c.JSON(CtxSuccess(req))
}
