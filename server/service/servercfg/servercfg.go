package servercfg

import (
	"fmt"

	"github.com/FlowingSPDG/ProjectDeus/server/entity"
)

// this package helps generating server-specific configration for server.cfg

const logAddressHTTPFormat = `log on;
logaddress_delall_http;
logaddress_add_http "%s";
`

type ServerCfg interface {
	// GenerateLogAddressHTTP generates logaddress_add_http for server.cfg
	GenerateLogAddressHTTP(serverID entity.GameServerID) string
}

type serverCfg struct {
	endpointFormat string // e.g. "http://ProjectDeusHostedAt:8080/game_server/%s/log";
}

func NewServerCfg(endpointFormat string) ServerCfg {
	return &serverCfg{
		endpointFormat: endpointFormat,
	}
}

// GenerateLogAddressHTTP implements ServerCfg.
// TODO: Support delayed stream(logaddress_add_http_delayed)
func (s *serverCfg) GenerateLogAddressHTTP(serverID entity.GameServerID) string {
	endpoint := fmt.Sprintf(s.endpointFormat, serverID)
	return fmt.Sprintf(logAddressHTTPFormat, endpoint)
}
