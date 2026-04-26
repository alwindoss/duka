package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginControllerRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type errResp struct {
	ErrCode    string `json:"err_code"`
	ErrMessage string `json:"err_message"`
}

type loginControllerResponse struct {
	Message string   `json:"message,omitempty"`
	Code    int      `json:"code"`
	Token   string   `json:"token,omitempty"`
	Err     *errResp `json:"err,omitempty"`
}

func LoginController(svc *LoginService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Here is login controller")
		inp := new(loginControllerRequest)
		ctx.BindJSON(inp)
		err := svc.Login(inp.UserName, inp.Password)
		resp := new(loginControllerResponse)

		if err != nil {
			resp.Code = http.StatusUnauthorized
			resp.Err = &errResp{
				ErrCode:    "ERR50001",
				ErrMessage: "unable to login",
			}
			ctx.JSON(http.StatusUnauthorized, resp)
			return
		}
		resp.Code = http.StatusOK
		resp.Message = "Login successful"
		resp.Token = "This should contain the token"
		ctx.JSON(http.StatusOK, resp)
	}
}
