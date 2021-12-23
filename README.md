# go-pushdeer-sdk

[pushdeer](https://github.com/easychen/pushdeer) 的golanga版本的sdk

![GitHub forks](https://img.shields.io/github/forks/Luoxin/go-pushdeer-sdk?style=social)
![GitHub Repo stars](https://img.shields.io/github/stars/Luoxin/go-pushdeer-sdk?style=social)
![GitHub](https://img.shields.io/github/license/yinguobing/cnn-facial-landmark)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Luoxin/go-pushdeer-sdk)

## 使用方式

在模块中安装

`go get github.com/Luoxin/go-pushdeer-sdk`

### 注册设备

```go
package main

import (
	"github.com/Luoxin/go-pushdeer-sdk/psdk"
	log "github.com/sirupsen/logrus"
)

func main() {
	p, err := psdk.New("http://127.0.0.1:8800", "token")
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	resp, err := p.DeviceList(&psdk.DeviceListReq{})
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	for _, device := range resp.Content.Devices {
		log.Infof("%+v", device)
	}
}
```

## TODO

- [ ] 完善测试
- [ ] 完善自动化检测
- [ ] `/message/push` 接口中的result的处理
