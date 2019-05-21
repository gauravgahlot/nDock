package converters

import (
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/constants"
	"github.com/gauravgahlot/dockerdoodle/pb"
	"github.com/gauravgahlot/dockerdoodle/types"
)

// ToHostsViewModel returns a collection of Host view model
func ToHostsViewModel(r *pb.GetContainersCountResponse, hosts []types.Host) *[]vm.Host {
	res := []vm.Host{}
	for i, host := range hosts {
		h := vm.Host{
			Name:           host.Name,
			IP:             host.IP,
			ContainerCount: int(r.HostContainers[i].Containers[host.IP]),
		}
		res = append(res, h)
	}
	return &res
}

// ToContainersViewModel returns pointers to collection of Container view model and GetStatsRequest
func ToContainersViewModel(r *pb.GetContainersResponse, host string) (*[]vm.Container, *pb.GetStatsRequest) {
	res := []vm.Container{}
	req := pb.GetStatsRequest{Host: host, Containers: map[string]int32{}}

	for i, c := range r.Containers {
		res = append(res, vm.Container{
			ID:        c.Id,
			Name:      c.Name,
			Image:     c.Image,
			Command:   c.Command,
			Created:   c.Created,
			State:     c.State,
			Status:    c.Status,
			Ports:     *getPorts(c.Ports),
			Mounts:    *getMounts(c.Mounts),
			ColorCode: constants.BGCodes[i],
		})
		req.Containers[c.Id] = int32(i)
	}
	return &res, &req
}

func getPorts(ports []*pb.Port) *[]vm.Port {
	ps := []vm.Port{}
	for _, p := range ports {
		ps = append(ps, vm.Port{
			IP:          p.IP,
			Type:        p.Type,
			PrivatePort: p.PrivatePort,
			PublicPort:  p.PublicPort,
		})
	}
	return &ps
}

func getMounts(mounts []*pb.MountPoint) *[]vm.Mount {
	ms := []vm.Mount{}
	for _, m := range mounts {
		ms = append(ms, vm.Mount{
			Type:        m.Type,
			Name:        m.Name,
			Source:      m.Source,
			Destination: m.Destination,
			Mode:        m.Mode,
			RW:          m.RW,
		})
	}
	return &ms
}
