package global

import (
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
	"sync"
)

var (
	GVAR_DB                   *gorm.DB
	GVAR_FILE_TYPE_MAP        sync.Map
	GVAR_ID_GENERATOR         *sonyflake.Sonyflake
	GVAR_AUTO_CREATE_DB       bool            = true                   //是否自动生成数据库
	GVAR_MAX_USERNAME_LENGTH  int             = 32                     // 用户名最大长度
	GVAR_MIN_PASSWORD_PATTERN string          = "^[_a-zA-Z0-9]{6,32}$" //密码格式
	GVAR_JWT_SigningKey       string          = "douyin-fighting"      //初始化全局签名
	GVAR_START_TIME           string          = "2022-05-21 00:00:01"  //固定启动时间，保证生成 ID 唯一性
	GVAR_FEED_NUM             int             = 10                     //每次返回视频数量
	GVAR_VIDEO_ADDR           string          = "./public/video/"      //视频存放位置
	GVAR_COVER_ADDR           string          = "./public/cover/"      //封面存放位置
	GVAR_MAX_FILE_SIZE        int64           = 10 << 20               // 上传文件大小限制为10MB
	GVAR_MAX_TITLE_LENGTH     int             = 140                    // 视频描述最大长度
	GVAR_MAX_COMMENT_LENGTH   int             = 300                    // 评论最大长度
	GVAR_WHITELIST_VIDEO      map[string]bool = map[string]bool{".mp4": true, ".avi": true, ".wmv": true, ".mpeg": true,
		".mov": true, ".flv": true, ".rmvb": true, ".3gb": true, ".vob": true, ".m4v": true}
)
