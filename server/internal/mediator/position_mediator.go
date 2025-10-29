package mediator

import (
	"context"
	"proto_buffer_example/server/generated/json_api"
	"proto_buffer_example/server/internal/player/interface/igamer"
)

type IPositionModelMediator interface {
	OnUpdatePosition(ctx context.Context, gamer igamer.IGamer, c2s *json_api.C2SPositionUpdate, s2c *json_api.S2CPositionUpdate)
}
