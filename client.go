package go_xen_client

import (
	"errors"
	"fmt"
	"log"

	"github.com/nilshell/xmlrpc"
)

type XenClient struct {
	Session  interface{}
	Host     string
	URL      string
	Username string
	Password string
	RPC      *xmlrpc.Client
}

func NewXenClient(host, username, password string) *XenClient {
	url := fmt.Sprintf("http://%s", host)
	rpc, _ := xmlrpc.NewClient(url, nil)
	return &XenClient{
		Host:     host,
		URL:      url,
		Username: username,
		Password: password,
		RPC:      rpc,
	}
}

func (client *XenClient) Login() (err error) {
	//Do loging call
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)
	params[0] = client.Username
	params[1] = client.Password

	err = client.RPCCall(&result, "session.login_with_password", params)
	if err == nil {
		if resultError, ok := result["ErrorDescription"]; ok {
			errorDescription := resultError.([]interface{})
			errorName := errorDescription[0]

			if errorName == "HOST_IS_SLAVE" {
				master := errorDescription[1].(string)

				//log.Infof("Host:%s is slave. Attempting relogin to master:%s", client.Host, master)
				client.Host = master
				client.URL = fmt.Sprintf("http://%s", master)
				client.RPC, _ = xmlrpc.NewClient(client.URL, nil)

				return client.Login()
			}
		}
		// err might not be set properly, so check the reference
		if result["Value"] == nil {
			return errors.New("Invalid credentials supplied")
		}
	}
	client.Session = result["Value"]
	return err
}

func (c *XenClient) RPCCall(result interface{}, method string, params []interface{}) (err error) {
	//log.Debugf("RPCCall method=%v params=%v\n", method, params)
	p := new(xmlrpc.Params)
	p.Params = params
	err = c.RPC.Call(method, *p, result)
	return err
}

func (client *XenClient) APICall(method string, params ...interface{}) (result interface{}, err error) {
	if client.Session == nil {
		log.Printf("[ERR] no session\n")
		return result, fmt.Errorf("No session. Unable to make call")
	}

	//Make a params slice which will include the session
	p := make([]interface{}, len(params)+1)
	p[0] = client.Session

	for idx, element := range params {
		p[idx+1] = element
	}

	res := xmlrpc.Struct{}

	err = client.RPCCall(&res, method, p)
	if err != nil {
		return result, err
	}

	status := res["Status"].(string)
	if status != "Success" {
		//log.Errorf("Encountered an API error: %v %v", result.Status, res["ErrorDescription"])
		return nil, fmt.Errorf("API Error: %s", res["ErrorDescription"])
	} else {
		result = res["Value"]
	}
	return result, nil
}
