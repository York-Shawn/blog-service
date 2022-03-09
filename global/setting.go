package global

import (
	"github.com/York-Shawn/blog-service/pkg/logger"
	"github.com/York-Shawn/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	JWTSetting      *setting.JWTSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	EmailSetting    *setting.EmailSettingS
)
