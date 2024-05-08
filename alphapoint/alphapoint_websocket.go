package alphapoint

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/aaabigfish/goex"
	"github.com/aaabigfish/gopkg/log"
)

const (
	alphapointDefaultWebsocketURL = "wss://sim3.alphapoint.com:8401/v1/GetTicker/"
)

// WebsocketClient starts a new webstocket connection
func (a *Alphapoint) WebsocketClient() {
	for a.Enabled {
		var dialer websocket.Dialer
		var err error
		var httpResp *http.Response
		endpoint, err := a.API.Endpoints.GetURL(goex.WebsocketSpot)
		if err != nil {
			log.Errorf("WebsocketMgr err(%v)", err)
		}
		a.WebsocketConn, httpResp, err = dialer.Dial(endpoint, http.Header{})
		httpResp.Body.Close() // not used, so safely free the body

		if err != nil {
			log.Errorf("ExchangeSys %s Unable to connect to Websocket. err(%v)", a.Name, err)
			continue
		}

		if a.Verbose {
			log.Debugf("ExchangeSys %s Connected to Websocket.", a.Name)
		}

		err = a.WebsocketConn.WriteMessage(websocket.TextMessage, []byte(`{"messageType": "logon"}`))

		if err != nil {
			log.Errorf("ExchangeSys err(%v)", err)
			return
		}

		for a.Enabled {
			msgType, resp, err := a.WebsocketConn.ReadMessage()
			if err != nil {
				log.Errorf("ExchangeSys err(%v)", err)
				break
			}

			if msgType == websocket.TextMessage {
				type MsgType struct {
					MessageType string `json:"messageType"`
				}

				msgType := MsgType{}
				err := json.Unmarshal(resp, &msgType)
				if err != nil {
					log.Errorf("ExchangeSys err(%v)", err)
					continue
				}

				if msgType.MessageType == "Ticker" {
					ticker := WebsocketTicker{}
					err = json.Unmarshal(resp, &ticker)
					if err != nil {
						log.Errorf("ExchangeSys err(%v)", err)
						continue
					}
				}
			}
		}
		a.WebsocketConn.Close()
		log.Infof("ExchangeSys %s Websocket client disconnected.", a.Name)
	}
}
