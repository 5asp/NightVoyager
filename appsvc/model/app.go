package model

type AppInfo struct {
	ID        int    `json:"id"`
	AppKey    string `json:"app_key"`
	AppName   string `json:"app_name"`
	AppSecret string `json:"app_secret"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (t AppInfo) Table() string {
	return "t_sms_appinfo"
}
