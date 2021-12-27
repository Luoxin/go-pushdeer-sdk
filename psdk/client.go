package psdk

import (
	"encoding/json"
	"github.com/Luoxin/go-pushdeer-sdk/perror"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type PushdeerClient struct {
	apiBase    string
	token      string
	httpClient *resty.Client
}

func getBaseClient(apiBase string) *PushdeerClient {
	return &PushdeerClient{
		apiBase: apiBase,
		httpClient: resty.New().
			SetTimeout(time.Minute).
			SetRetryWaitTime(time.Second).
			SetRetryCount(5).
			SetHeaders(map[string]string{}).
			SetLogger(log.New()).
			SetBaseURL(apiBase),
	}
}

func New(apiBase string, token string) (*PushdeerClient, error) {
	p := getBaseClient(apiBase)
	p.token = token

	return p, nil
}

// 不需要传入token，会候通过 `/login/fake` 生成一个token
func NewFaker(apiBase string) (*PushdeerClient, error) {
	p := getBaseClient(apiBase)

	fakerRsp, err := p.LoginFaker()
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	p.token = fakerRsp.Content.Token

	return p, nil
}

func (p *PushdeerClient) LoginFaker() (*LoginFakeRsp, error) {
	var rsp LoginFakeRsp
	resp, err := p.httpClient.R().
		SetResult(&rsp).
		Get(LoginFake)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, &BaseRsp{
		Code: rsp.Code,
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) UserInfo(req *UserInfoReq) (*UserInfoRsp, error) {
	req.Token = p.token

	var rsp UserInfoRsp
	err := p.Post(UserInfo, *req, &rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) DeviceReg(req *DeviceRegReq) (*DeviceRegRsp, error) {
	req.Token = p.token

	var rsp DeviceRegRsp
	resp, err := p.httpClient.R().SetResult(&rsp).Post(DeviceReg)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) DeviceList(req *DeviceListReq) (*DeviceListRsp, error) {
	req.Token = p.token

	var rsp DeviceListRsp
	resp, err := p.httpClient.R().SetResult(&rsp).Post(DeviceList)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) DeviceRemove(req *DeviceRemoveReq) (*DeviceRemoveRsp, error) {
	req.Token = p.token

	var rsp DeviceRemoveRsp
	resp, err := p.httpClient.R().SetResult(&rsp).Post(DeviceRemove)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) KeyGen(req *KeyGenReq) (*KeyGenRsp, error) {
	req.Token = p.token

	var rsp KeyGenRsp
	resp, err := p.httpClient.R().SetResult(&rsp).Post(KeyGen)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) KeyRegen(req *KeyRegenReq) (*KeyRegenRsp, error) {
	req.Token = p.token

	var rsp KeyRegenRsp
	resp, err := p.httpClient.R().SetResult(&rsp).Post(KeyRegen)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) KeyList(req *KeyListReq) (*KeyListRsp, error) {
	req.Token = p.token
	var rsp KeyListRsp
	resp, err := p.httpClient.R().Post(KeyList)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) KeyRemove(req *KeyRemoveReq) (*KeyRemoveRsp, error) {
	req.Token = p.token
	var rsp KeyRemoveRsp
	resp, err := p.httpClient.R().Post(KeyRemove)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) MessagePush(req *MessagePushReq) (*MessagePushRsp, error) {
	req.Token = p.token

	var rsp MessagePushRsp
	resp, err := p.httpClient.R().Post(MessagePush)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) MessageList(req *MessageListReq) (*MessageListRsp, error) {
	req.Token = p.token
	var rsp MessageListRsp
	resp, err := p.httpClient.R().Post(MessageList)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) MessageRemove(req *MessageRemoveReq) (*MessageRemoveRsp, error) {
	req.Token = p.token
	var rsp MessageRemoveRsp
	resp, err := p.httpClient.R().Post(MessageRemove)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &rsp, nil
}

func (p *PushdeerClient) checkResp(resp *resty.Response, code Response) error {
	switch resp.StatusCode() {
	case http.StatusOK:
		switch code.GetCode() {
		case 0:
			return nil
		default:
			return perror.CreateErr(code.GetCode(), code.GetContent())
		}
	default:
		return perror.CreateErr(-resp.StatusCode(), resp.Status())
	}
}

func (p *PushdeerClient) Post(path string, req interface{}, rsp Response) error {
	buf, err := json.Marshal(req)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	var reqForm map[string]string
	err = json.Unmarshal(buf, &reqForm)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	resp, err := p.httpClient.R().SetFormData(reqForm).SetResult(rsp).Post(path)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	err = p.checkResp(resp, rsp)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}
