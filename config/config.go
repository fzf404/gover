/*
 * @Author: fzf404
 * @Date: 2022-02-10 11:12:02
 * @LastEditTime: 2022-02-27 08:18:49
 * @Description: 配置文件结构
 */
package config

/**
* @description: 配置结构
 */
type Config struct {
	Common Common `mapstructure:"common"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Email  Email  `mapstructure:"email"`
	JWT    JWT    `mapstructure:"jwt"`
}

/**
* @description: 基础配置
 */
type Common struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`

	UploadPath string `mapstructure:"upload_path"`
	UrlPrefix  string `mapstructure:"url_prefix"`
}

/**
* @description: Mysql 配置
 */
type Mysql struct {
	Host     string `mapstructure:"host"`     // 服务器地址
	Port     string `mapstructure:"port"`     // 端口
	Username string `mapstructure:"username"` // 数据库用户名
	Password string `mapstructure:"password"` // 数据库密码
	Database string `mapstructure:"database"` // 数据库名
	Config   string `mapstructure:"config"`   // 高级配置
}

/**
* @description: 数据库配置转换
 */
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Config
}

/**
* @description: Redis 配置
 */
type Redis struct {
	Addr     string `mapstructure:"addr"`     // 服务器地址
	Password string `mapstructure:"password"` // 数据库密码
	DB       int    `mapstructure:"db"`       // 数据库
}

/**
* @description: 邮件配置
 */
type Email struct {
	Host     string `mapstructure:"host"`     // 发件人邮箱
	Smtp     string `mapstructure:"smtp"`     // smtp服务器
	Addr     string `mapstructure:"addr"`     // 发件人邮箱地址
	Name     string `mapstructure:"name"`     // 发件人名称
	Password string `mapstructure:"password"` // 发件人邮箱密码
	Expires  int    `mapstructure:"expires"`  // 过期时间
}

/**
* @description: 鉴权配置
 */
type JWT struct {
	Key     string `mapstructure:"key"`     // 签名密钥
	Issuer  string `mapstructure:"issuer"`  // 发行人
	Expires int64  `mapstructure:"expires"` // 过期时间
}
