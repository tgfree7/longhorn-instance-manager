package process

import (
	"net"

	"github.com/longhorn/longhorn-instance-manager/pkg/rpc"
	"golang.org/x/net/context"
)

func (pm *Manager) GetInterfaceIP(ctx context.Context, req *rpc.Interface) (ret *rpc.DataIP, err error) {
	inter, err := net.InterfaceByName(req.Interface)
	if err != nil {
		return nil, err
	}
	if (inter.Flags & net.FlagUp) != 0 {
		addrs, err := inter.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
				if ip.IP.To4() != nil {
					return &rpc.DataIP{
						DataIP: ip.IP.String(),
					}, nil
				}
			}
		}
	}
	return nil, nil
}
