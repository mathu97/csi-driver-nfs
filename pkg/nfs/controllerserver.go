package nfs

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
)

type newControllerServer struct {
	*csicommon.DefaultControllerServer
}

func (cs newControllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return &csi.ControllerExpandVolumeResponse{}, nil
}

func getControllerServer(csiDriver *csicommon.CSIDriver) newControllerServer {
	return newControllerServer{
		csicommon.NewDefaultControllerServer(csiDriver),
	}
}
