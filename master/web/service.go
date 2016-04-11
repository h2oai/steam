package web

import (
	// "fmt"
	// "github.com/h2oai/steam/lib/az"
	// "github.com/h2oai/steam/lib/fs"
	// "github.com/h2oai/steam/lib/grid"
	// "github.com/h2oai/steam/lib/sec"
	// "github.com/h2oai/steam/master/cli"
	// "github.com/h2oai/steam/master/ctl"
	"github.com/h2oai/steamY/master/db"
	// "github.com/h2oai/steam/master/job"
	// "github.com/h2oai/steam/master/proc"
	// "github.com/h2oai/steam/master/usr/auth"
	// "github.com/h2oai/steam/srv/h2ov3"
	"github.com/h2oai/steamY/srv/web"
	// "os"
	// "regexp"
	// "strconv"
	// "strings"
	// "time"
)

type Service struct {
	wd string
	ds *db.DS
}

func NewService(wd string, ds *db.DS) *web.Impl {
	return &web.Impl{&Service{wd, ds}}
}

func (s *Service) Ping(status bool) (bool, error) {
	return false, nil
}
