package wireguard

import (
	"context"
	"errors"
	"fmt"
	"net/netip"
	"strings"

	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/log"
	"golang.zx2c4.com/wireguard/device"
)

//go:generate go run github.com/xtls/xray-core/common/errors/errorgen

var wgLogger = &device.Logger{
	Verbosef: func(format string, args ...any) {
		log.Record(&log.GeneralMessage{
			Severity: log.Severity_Debug,
			Content:  fmt.Sprintf(format, args...),
		})
	},
	Errorf: func(format string, args ...any) {
		log.Record(&log.GeneralMessage{
			Severity: log.Severity_Error,
			Content:  fmt.Sprintf(format, args...),
		})
	},
}

func init() {
	common.Must(common.RegisterConfig((*DeviceConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		deviceConfig := config.(*DeviceConfig)
		if deviceConfig.IsClient {
			return New(ctx, deviceConfig)
		} else {
			return NewServer(ctx, deviceConfig)
		}
	}))
}

// convert endpoint string to netip.Addr
func parseEndpoints(conf *DeviceConfig) ([]netip.Addr, bool, bool, error) {
	var hasIPv4, hasIPv6 bool
}

