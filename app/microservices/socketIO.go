package microservices

import (
	//Standard library packages
	"log"
	"encoding/json"

	//Third party packages
	"github.com/googollee/go-socket.io"
)

type (
	Channels struct {
		Channels	[]string	`json:"channels"`
	}

	Publish struct {
		Channel		string		`json:"channel"`
		Event		string		`json:"event"`
		Message		string		`json:"message"`
	}
)

func SocketIO() (*socketio.Server) {

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}


	// CONNECT
	server.On("connection", func(so socketio.Socket) {

		log.Println(so.Id() + " on connection")


		// SUBSCRIBE
		so.On("subscribe", func(paramsStr string) bool {

			var params Channels
			err := json.Unmarshal([]byte(paramsStr), &params)

			if err != nil {
				return false
			}

			if len(params.Channels) == 0 {
				return false
			}

			for _, channel := range params.Channels {
				err := so.Join(channel)
				if err != nil {
					return false
				}
			}

			return true
		})


		// UNSUBSCRIBE
		so.On("unsubscribe", func(paramsStr string) bool {

			var params Channels
			err := json.Unmarshal([]byte(paramsStr), &params)

			if err != nil {
				return false
			}

			if len(params.Channels) == 0 {
				return false
			}

			for _, channel := range params.Channels {
				err := so.Leave(channel)
				if err != nil {
					return false
				}
			}

			return true
		})


		// PUBLISH
		so.On("publish", func(paramsStr string) bool {

			var params Publish
			err := json.Unmarshal([]byte(paramsStr), &params)

			if err != nil {
				return false
			}

			so.BroadcastTo(params.Channel, params.Event, params.Message)

			return true
		})


		// DISCONNECT
		so.On("disconnection", func() {
			log.Println(so.Id() + " on disconnect")
		})

	})


	// ERROR
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}