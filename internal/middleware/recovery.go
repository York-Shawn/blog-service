package middleware

import (
	"fmt"
	"time"

	"github.com/York-Shawn/blog-service/global"
	"github.com/York-Shawn/blog-service/pkg/app"
	"github.com/York-Shawn/blog-service/pkg/email"
	"github.com/York-Shawn/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	detailMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		PassWord: global.EmailSetting.PassWord,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)

				err := detailMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("throw exception, time: %d", time.Now().Unix()),
					fmt.Sprintf("error info: %v", err),
				)
				if err != nil {
					global.Logger.Panic("mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
