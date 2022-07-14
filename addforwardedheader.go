package AddForwardedHeader

import (
	"context"
	"net/http"
	"strings"
)

// Config holds configuration to be passed to the plugin
type Config struct {
	By        string `yaml:"by"`
	ForHeader string `yaml:"forHeader"`
}

// CreateConfig populates the Config data object
func CreateConfig() *Config {
	return &Config{}
}

// Plugin holds the necessary components of a Traefik plugin
type Plugin struct {
	by        string
	forHeader string
	name      string
	next      http.Handler
}

// New instantiates and returns the required components used to handle a HTTP request
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Plugin{
		by:        strings.TrimSpace(config.By),
		forHeader: strings.TrimSpace(config.ForHeader),
		name:      name,
		next:      next,
	}, nil
}

func (plugin *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	Forwarded := []string{}

	By := plugin.by
	For := strings.TrimSpace(req.Header.Get(plugin.forHeader))
	Host := strings.TrimSpace(req.Header.Get("X-Forwarded-Host"))
	Proto := strings.TrimSpace(req.Header.Get("X-Forwarded-Proto"))

	if len(By) > 0 {
		Forwarded = append(Forwarded, "by="+By)
	}
	if len(For) > 0 {
		Forwarded = append(Forwarded, "for="+For)
	}
	if len(Host) > 0 {
		Forwarded = append(Forwarded, "host="+Host)
	}
	if len(Proto) > 0 {
		Forwarded = append(Forwarded, "proto="+Proto)
	}

	req.Header.Set("Forwarded", strings.Join(Forwarded, ";"))

	plugin.next.ServeHTTP(rw, req)

}
