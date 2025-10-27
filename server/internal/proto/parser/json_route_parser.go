package parser

import (
	"encoding/json"
	"fmt"
	"proto_buffer_example/server/generated/json_api"
	"proto_buffer_example/server/third-party/antnet"
	"reflect"
)

// JsonRouteParser implements antnet.IParser for JSON messages with route-based dispatch.
type JsonRouteParser struct {
	*antnet.Parser // Embed antnet.Parser to inherit common fields/methods

	// Map to store message types based on their route string
	routeMap map[string]reflect.Type
}

// NewJsonRouteParser creates a new JsonRouteParser.
func NewJsonRouteParser(baseParser *antnet.Parser) *JsonRouteParser {
	return &JsonRouteParser{
		Parser:   baseParser,
		routeMap: make(map[string]reflect.Type),
	}
}

// GenericJsonMessage is a temporary struct to extract the route from incoming JSON.
type GenericJsonMessage struct {
	Route string `json:"route"`
	// RequestId string `json:"request_id"` // If request_id is re-enabled
}

// ParseC2S parses an incoming C2S message based on its route.
func (r *JsonRouteParser) ParseC2S(msg *antnet.Message) (antnet.IMsgParser, error) {
	if msg == nil || len(msg.Data) == 0 {
		return nil, antnet.ErrMsgPackUnPack // Or a more specific error
	}

	// 1. Extract the route from the raw JSON data
	var genericMsg GenericJsonMessage
	err := json.Unmarshal(msg.Data, &genericMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal generic JSON message: %w", err)
	}

	if genericMsg.Route == "" {
		return nil, fmt.Errorf("JSON message missing 'route' field")
	}

	// 2. Look up the actual message type based on the route
	msgType, ok := r.routeMap[genericMsg.Route]
	if !ok {
		return nil, fmt.Errorf("unknown route: %s", genericMsg.Route)
	}

	// 3. Create a new instance of the target message type
	c2sMsg := reflect.New(msgType.Elem()).Interface()

	// 4. Unmarshal the raw JSON data into the target message type
	err = json.Unmarshal(msg.Data, c2sMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON for route %s into type %s: %w", genericMsg.Route, msgType.String(), err)
	}

	// Create an antnet.MsgParser to hold the deserialized message
	return &jsonRouteMsgParser{c2s: c2sMsg, parser: r}, nil
}

// jsonRouteMsgParser is a simple implementation of antnet.IMsgParser for our custom parser.
type jsonRouteMsgParser struct {
	c2s    interface{}
	s2c    interface{} // Not used for incoming, but needed for interface
	parser antnet.IParser
}

func (r *jsonRouteMsgParser) C2S() interface{}  { return r.c2s }
func (r *jsonRouteMsgParser) S2C() interface{}  { return r.s2c }
func (r *jsonRouteMsgParser) C2SData() []byte   { return r.parser.PackMsg(r.c2s) }
func (r *jsonRouteMsgParser) S2CData() []byte   { return r.parser.PackMsg(r.s2c) }
func (r *jsonRouteMsgParser) C2SString() string { return string(r.C2SData()) }
func (r *jsonRouteMsgParser) S2CString() string { return string(r.S2CData()) }

// PackMsg packs an outgoing S2C message into JSON bytes.
func (r *JsonRouteParser) PackMsg(v interface{}) []byte {
	jsonData, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("Error marshalling S2C message to JSON: %v\n", err)
		return nil // Or handle error appropriately
	}
	return jsonData
}

// GetType returns the parser type.
func (r *JsonRouteParser) GetType() antnet.ParserType {
	return antnet.ParserTypeCustom
}

// GetErrType returns the error type for the parser.
func (r *JsonRouteParser) GetErrType() antnet.ParseErrType {
	return r.Parser.GetErrType() // Inherit from base parser
}

// GetRemindMsg creates a reminder message for errors.
func (r *JsonRouteParser) GetRemindMsg(err error, t antnet.MsgType) *antnet.Message {
	// For JSON, we can send back an ErrorResponse
	errorResponse := &json_api.ErrorResponse{
		Code:    500, // Internal Server Error
		Message: err.Error(),
	}
	jsonData, _ := json.Marshal(errorResponse)
	return &antnet.Message{Data: jsonData, Head: nil}
}

// RegisterMsg registers a C2S message type with its route.
// This method will be called from NewWebSocketServer.
func (r *JsonRouteParser) RegisterMsg(route string, c2s interface{}, s2c interface{}) {
	c2sType := reflect.TypeOf(c2s)
	if c2sType.Kind() != reflect.Ptr {
		panic("C2S message must be a pointer to a struct")
	}

	if route == "" {
		panic(fmt.Sprintf("Provided route for C2S message type %s is empty", c2sType.String()))
	}

	r.routeMap[route] = c2sType
}
