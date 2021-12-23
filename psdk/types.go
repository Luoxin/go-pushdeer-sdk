package psdk

import "time"

type BaseReq struct {
	token string `json:"token"`
}

type BaseRsp struct {
	Code int `json:"code"`
}

type MsgType string

const (
	MsgMarkdown MsgType = "markdown"
	MsgText     MsgType = "text"
	MsgImage    MsgType = "image"
)

type LoginFakeRsp struct {
	Code int `json:"code"`

	Content struct {
		Token string `json:"token"`
	} `json:"content"`
}

type UserInfoReq struct {
	BaseReq
}

type UserInfoRsp struct {
	BaseRsp
	Uid   int    `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Level int    `json:"level"`
}

type DeviceRegReq struct {
	BaseReq
	Name     string `json:"name,omitempty"`
	DeviceId string `json:"device_id,omitempty"`
	IsClip   int    `json:"is_clip,omitempty"`
}

type DeviceRegRsp struct {
	BaseRsp
	Content struct {
		Devices []struct {
			Id       int    `json:"id"`
			Uid      string `json:"uid"`
			Name     string `json:"name"`
			Type     string `json:"type"`
			DeviceId string `json:"device_id"`
			IsClip   int    `json:"is_clip"`
		} `json:"devices"`
	} `json:"content"`
}

type DeviceListReq struct {
	BaseReq
}

type DeviceListRsp struct {
	BaseRsp
	Content struct {
		Devices []struct {
			Id       int    `json:"id"`
			Uid      string `json:"uid"`
			Name     string `json:"name"`
			Type     string `json:"type"`
			DeviceId string `json:"device_id"`
			IsClip   int    `json:"is_clip"`
		} `json:"devices"`
	} `json:"content"`
}

type DeviceRemoveReq struct {
	BaseReq
	Id int64 `json:"id,omitempty"`
}

type DeviceRemoveRsp struct {
	BaseRsp
	Content struct {
		Message string `json:"message"`
	} `json:"content"`
}

type KeyGenReq struct {
	BaseReq
}

type KeyGenRsp struct {
	BaseRsp
	Content struct {
		Keys []struct {
			Id  int    `json:"id"`
			Uid string `json:"uid"`
			Key string `json:"key"`
		} `json:"keys"`
	} `json:"content"`
}

type KeyRegenReq struct {
	BaseReq
}

type KeyRegenRsp struct {
	BaseRsp
	Content struct {
		Message string `json:"message"`
	} `json:"content"`
}

type KeyListReq struct {
	BaseReq
	Id int64 `json:"id,omitempty"`
}

type KeyListRsp struct {
	BaseRsp
	Content struct {
		Keys []struct {
			Id  int    `json:"id"`
			Uid string `json:"uid"`
			Key string `json:"key"`
		} `json:"keys"`
	} `json:"content"`
}

type KeyRemoveReq struct {
	BaseReq
	Id int64 `json:"id,omitempty"`
}

type KeyRemoveRsp struct {
	BaseRsp
	Content struct {
		Message string `json:"message"`
	} `json:"content"`
}

type MessagePushReq struct {
	BaseReq
	PushKey string  `json:"pushkey,omitempty"`
	Text    string  `json:"text,omitempty"`
	Desp    string  `json:"desp,omitempty"`
	Type    MsgType `json:"type,omitempty"`
}

type MessagePushRsp struct {
	BaseRsp
	Result []string `json:"result"`
}

type MessageListReq struct {
	BaseReq
	Limit uint32 `json:"limit,omitempty" validate:"max=100"`
}

type MessageListRsp struct {
	BaseRsp
	Content struct {
		Messages []struct {
			Id        int       `json:"id"`
			Uid       string    `json:"uid"`
			Text      string    `json:"text"`
			Desp      string    `json:"desp"`
			Type      string    `json:"type"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"messages"`
	} `json:"content"`
}

type MessageRemoveReq struct {
	BaseReq
	Id int64 `json:"id,omitempty"`
}

type MessageRemoveRsp struct {
	BaseRsp
	Content struct {
		Message string `json:"message"`
	} `json:"content"`
}
