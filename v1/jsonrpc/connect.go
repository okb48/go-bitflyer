package jsonrpc

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/buger/jsonparser"

	"github.com/labstack/gommon/log"

	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"

	"github.com/go-numb/go-bitflyer/v1/public/board"
	"github.com/go-numb/go-bitflyer/v1/public/executions"
	"github.com/go-numb/go-bitflyer/v1/public/ticker"
	"golang.org/x/net/websocket"
)

const (
	All Types = iota
	Ticker
	Executions
	Board
	Error

	HeartbeatIntervalSecond time.Duration = 60
	ReadTimeoutSecond       time.Duration = 300
	WriteTimeoutSecond      time.Duration = 5

	ORIGIN  = "https://lightning.bitflyer.com"
	BASEURL = "wss://ws.lightstream.bitflyer.com/json-rpc"
)

type Types int

type Response struct {
	Type  Types
	Error error

	Ticker     ticker.Response
	Executions executions.Response
	Orderbook  board.Response
}

type ResponseTicker struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string          `json:"channel"`
		Message ticker.Response `json:"message"`
	} `json:"params"`
}
type ResponseExecution struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string              `json:"channel"`
		Message executions.Response `json:"message"`
	} `json:"params"`
}
type ResponseBoard struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string         `json:"channel"`
		Message board.Response `json:"message"`
	} `json:"params"`
}

func Get(channels []string, ch chan Response) {
	ws, err := websocket.Dial(BASEURL, "", ORIGIN)
	if err != nil {
		s := &Response{
			Error: errors.Wrap(err, "websocket connecting error: "),
		}
		ch <- *s
	}
	defer ws.Close()

	var (
		which Types

		requests []string
	)

	for _, channel := range channels {
		switch {
		case strings.HasPrefix(channel, "lightning_ticker"):
			fmt.Printf("type has %d\n", Ticker)
		case strings.HasPrefix(channel, "lightning_executions"):
			fmt.Printf("type has %d\n", Executions)
		case strings.HasPrefix(channel, "lightning_board"):
			fmt.Printf("type has %d\n", Board)
		}
		requests = append(requests, fmt.Sprintf(`{"method": "subscribe", "params": {"channel": "%s"}}`, channel))
	}

	if len(requests) == 4 {
		fmt.Printf("gets all channels %d\n", All)
	}

	for _, v := range requests {
		if _, err := ws.Write([]byte(v)); err != nil {
			s := &Response{
				Type:  Error,
				Error: errors.Wrap(err, "websocket write error: "),
			}
			ch <- *s
		}
	}

	var eg errgroup.Group

	// // recieves msg
	// eg.Go(func() error {
	// 	ticker := time.NewTicker(HeartbeatIntervalSecond * time.Second)
	// 	defer ticker.Stop()
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			if _, err := ws.Write([]byte("ping")); err != nil {
	// 				log.Error(err)
	// 				return errors.New("websocket heartbeat write, return nil")
	// 			}
	// 		}
	// 	}

	// 	return errors.New("websocket heartbeat write, return nil")
	// })

	eg.Go(func() error {
		var msg = make([]byte, 512)
		for {
			ws.SetReadDeadline(time.Now().Add(ReadTimeoutSecond * time.Second))
			if err = websocket.Message.Receive(ws, &msg); err != nil {
				return errors.Wrap(err, "can't receive error: ")
			}

			channelName, err := jsonparser.GetString(msg, "params", "channel")
			if err != nil {
				return errors.Wrap(err, "can't read channel name: ")
			}

			switch {
			case strings.HasPrefix(channelName, "lightning_ticker"):
				var res ResponseTicker
				json.Unmarshal(msg, &res)

				s := &Response{
					Type:   Ticker,
					Ticker: res.Params.Message,
				}
				// fmt.Printf("きたこれTicker: %+v\n", s.Type)
				ch <- *s

			case strings.HasPrefix(channelName, "lightning_executions"):
				var res ResponseExecution
				json.Unmarshal(msg, &res)
				s := &Response{
					Type:       Executions,
					Executions: res.Params.Message,
				}
				// fmt.Printf("きたこれExecute: %+v\n", s.Type)
				ch <- *s

			case strings.HasPrefix(channelName, "lightning_board"):
				var res ResponseBoard
				json.Unmarshal(msg, &res)
				s := &Response{
					Type:      Board,
					Orderbook: res.Params.Message,
				}
				// fmt.Printf("きたこれBoard: %+v\n", s.Orderbook)
				ch <- *s

			default:
				s := &Response{
					Type:  Error,
					Error: errors.New("read type error: " + string(which)),
				}
				ch <- *s
			}
		}
		return errors.New("websocket read has error")
	})

	if err := eg.Wait(); err != nil {
		log.Error(err)
		go func() {
			ch <- Response{
				Type:  Error,
				Error: errors.New("websocket type error: " + err.Error()),
			}
		}()
		ws.Close()
	}
}
