package setting

import "time"

type ServerSettingS struct {
	AppMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	JwtKey       string
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType   string
	UserName string
	Password string
	Host     string
	DBName   string
}

// 绑定到setting类型的指针
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
