// Package config 站点配置信息
package config

import "gen-resume/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认是 腾讯邮件推送 SES 的配置
			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_HOST", "smtp.qcloudmail.com"),
				"port":     config.Env("MAIL_PORT", 465),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},

			"from": map[string]interface{}{
				"address": config.Env("MAIL_FROM_ADDRESS", "no-replay@seifwu.com"),
				"name":    config.Env("MAIL_FROM_NAME", "no-replay"),
			},
		}
	})
}
