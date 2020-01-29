package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/olahol/melody"
)

var mel = melody.New() // se actualiza var mel *melody.Melody

func WebSocket(c echo.Context) error {
	mel.HandleRequest(c.Response().Writer, c.Request())
	mel.HandleConnect(hConnect)
	mel.HandleDisconnect(hDisconnect)
	mel.HandleMessage(hMessage)
	return nil
}

func hConnect(s *melody.Session) {
	if !validateToken(s) {
		return
	}
	nick := getNickFromURL(s)
	m := &WebSocketResponse{
		Type:    "connect",
		From:    nick,
		Data: "Conectado",
	}
	sendMessage(m)
}

func hDisconnect(s *melody.Session) {
	if !validateToken(s) {
		return
	}
	nick := getNickFromURL(s)
	m := &WebSocketResponse{
		Type:    "disconnect",
		From:    nick,
		Data: "Desconectado",
	}
	sendMessage(m)
}

func hMessage(s *melody.Session, msg []byte) {
	nick := getNickFromURL(s)

	var data = make(map[string]interface{}) //se corrige var data map[string]interface{}
	err := json.Unmarshal(msg, &data)       //se corrige err := json.Unmarshal(msg, data)
	if err != nil {
		log.Printf("no se pudo procesar el mensaje %v", err)
		return
	}

	if data["type"] == "ping" {
		mel.BroadcastFilter([]byte("pong"), func(ss *melody.Session) bool {
			return ss == s
		})
	}

	m := &WebSocketResponse{
		Type: data["type"].(string),
		From: nick,
		Data: data["data"].(string),
	}
	sendMessage(m)
}

func validateToken(s *melody.Session) bool {
	t := s.Request.URL.Query().Get("token")
	return t == token
}

func getNickFromURL(s *melody.Session) string {
	return s.Request.URL.Query().Get("nick")
}

func sendMessage(m *WebSocketResponse) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Printf("no se pudo convertir el mensaje a json: %v", err)
		return
	}

	mel.Broadcast(j)
}
