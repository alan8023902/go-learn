package notifier

import (
	"encoding/json"
	"fmt"
	"os"
)

// 工厂方法，根据类型创建不同的的实现
func CreateNotifier(kind string) (Notifier, error) {
	switch kind {
	case "email":
		return EmailNotifier{}, nil
	case "sms":
		return SmsNotifier{}, nil
	case "ai":
		return AINotifier{}, nil
	default:
		return nil, fmt.Errorf("未知通知类型： %s", kind)
	}
}

type NotifierConfig struct {
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

// 从配置文件加载

func LoadNotifierFromConfig(filePath string) ([]Notifier, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败： %v", err)
	}

	var configs []NotifierConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, fmt.Errorf("解析json 失败 ：%v", err)
	}

	var list []Notifier
	for _, cfg := range configs {
		if !cfg.Enabled {
			continue
		}
		n, err := CreateNotifier(cfg.Type)
		if err != nil {
			fmt.Println("", err)
			continue
		}
		list = append(list, n)
	}
	return list, nil
}
