package provisioning

import (
	"context"
	"log"

	"github.com/n0stack/proto.go/budget/v0"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/n0stack/n0core/pkg/api/pool/node"
	"github.com/n0stack/n0core/pkg/datastore"
	"github.com/n0stack/proto.go/pool/v0"
	"github.com/n0stack/proto.go/provisioning/v0"
)

type VolumeAPI struct {
	dataStore datastore.Datastore

	// dependency APIs
	nodeAPI ppool.NodeServiceClient

	nodeConnections *node.NodeConnections
}

const (
	// Create のときに Node を指定したい時に利用
	AnnotationRequestNodeName = "n0core/provisioning/request_node_name"

	// Create のときに自動生成、消されると困る
	AnnotationVolumePath = "n0core/provisioning/volume_url"
)

func CreateVolumeAPI(ds datastore.Datastore, na ppool.NodeServiceClient, nc *node.NodeConnections) (*VolumeAPI, error) {
	a := &VolumeAPI{
		dataStore:       ds,
		nodeAPI:         na,
		nodeConnections: nc,
	}

	return a, nil
}

func (a *VolumeAPI) CreateEmptyVolume(ctx context.Context, req *pprovisioning.CreateEmptyVolumeRequest) (*pprovisioning.Volume, error) {
	prev := &pprovisioning.Volume{}
	if err := a.dataStore.Get(req.Metadata.Name, prev); err == nil {
		return nil, grpc.Errorf(codes.AlreadyExists, "Volume '%s' is already exists", req.Metadata.Name)
	} else if status.Code(err) != codes.NotFound {
		log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Metadata.Name)
	}

	res := &pprovisioning.Volume{
		Metadata: req.Metadata,
		Spec:     req.Spec,
		Status:   &pprovisioning.VolumeStatus{},
	}

	var rcr *ppool.ReserveStorageResponse
	var err error
	if node, ok := req.Metadata.Annotations[AnnotationRequestNodeName]; !ok {
		rcr, err = a.nodeAPI.ScheduleStorage(context.Background(), &ppool.ScheduleStorageRequest{
			StorageName: req.Metadata.Name,
			Storage: &pbudget.Storage{
				RequestBytes: req.Spec.RequestBytes,
				LimitBytes:   req.Spec.LimitBytes,
			},
		})
	} else {
		rcr, err = a.nodeAPI.ReserveStorage(context.Background(), &ppool.ReserveStorageRequest{
			Name:        node,
			StorageName: req.Metadata.Name,
			Storage: &pbudget.Storage{
				RequestBytes: req.Spec.RequestBytes,
				LimitBytes:   req.Spec.LimitBytes,
			},
		})
	}
	if err != nil {
		return nil, err // TODO: #89
	}
	res.Status.NodeName = rcr.Name
	res.Status.StorageName = rcr.StorageName

	conn, err := a.nodeConnections.GetConnection(res.Status.NodeName) // errorについて考える
	if err != nil {
		_, err = a.nodeAPI.ReleaseStorage(context.Background(), &ppool.ReleaseStorageRequest{
			Name:        res.Status.NodeName,
			StorageName: res.Status.StorageName,
		})
		if err != nil {
			return nil, err // TODO: #89
		}

		log.Printf("Fail to dial to node: err=%v.", err.Error())
		return nil, grpc.Errorf(codes.Internal, "")
	}
	defer conn.Close()
	cli := NewVolumeAgentServiceClient(conn)

	v, err := cli.CreateEmptyVolumeAgent(context.Background(), &CreateEmptyVolumeAgentRequest{
		Name:  req.Metadata.Name,
		Bytes: req.Spec.LimitBytes,
	})
	if err != nil && status.Code(err) != codes.AlreadyExists {
		log.Printf("Fail to create volume on node '%s': err='%s'", "", err.Error()) // TODO: #89
		return nil, grpc.Errorf(codes.Internal, "")
	}

	res.Metadata.Annotations[AnnotationVolumePath] = v.Path
	res.Status.State = pprovisioning.VolumeStatus_AVAILABLE

	if err := a.dataStore.Apply(req.Metadata.Name, res); err != nil {
		_, err = a.nodeAPI.ReleaseStorage(context.Background(), &ppool.ReleaseStorageRequest{
			Name:        res.Status.NodeName,
			StorageName: res.Status.StorageName,
		})
		if err != nil {
			return nil, err // TODO: #89
		}

		log.Printf("[WARNING] Failed to apply data for db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to store '%s' for db, please retry or contact for the administrator of this cluster", req.Metadata.Name)
	}

	return res, nil
}

func (a *VolumeAPI) CreateVolumeWithDownloading(ctx context.Context, req *pprovisioning.CreateVolumeWithDownloadingRequest) (*pprovisioning.Volume, error) {
	prev := &pprovisioning.Volume{}
	if err := a.dataStore.Get(req.Metadata.Name, prev); err == nil {
		return nil, grpc.Errorf(codes.AlreadyExists, "Volume '%s' is already exists", req.Metadata.Name)
	} else if status.Code(err) != codes.NotFound {
		log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Metadata.Name)
	}

	res := &pprovisioning.Volume{
		Metadata: req.Metadata,
		Spec:     req.Spec,
		Status:   &pprovisioning.VolumeStatus{},
	}

	var rcr *ppool.ReserveStorageResponse
	var err error
	if node, ok := req.Metadata.Annotations[AnnotationRequestNodeName]; !ok {
		rcr, err = a.nodeAPI.ScheduleStorage(context.Background(), &ppool.ScheduleStorageRequest{
			StorageName: req.Metadata.Name,
			Storage: &pbudget.Storage{
				RequestBytes: req.Spec.RequestBytes,
				LimitBytes:   req.Spec.LimitBytes,
			},
		})
	} else {
		rcr, err = a.nodeAPI.ReserveStorage(context.Background(), &ppool.ReserveStorageRequest{
			Name:        node,
			StorageName: req.Metadata.Name,
			Storage: &pbudget.Storage{
				RequestBytes: req.Spec.RequestBytes,
				LimitBytes:   req.Spec.LimitBytes,
			},
		})
	}
	if err != nil {
		return nil, err // TODO: #89
	}
	res.Status.NodeName = rcr.Name
	res.Status.StorageName = rcr.StorageName

	conn, err := a.nodeConnections.GetConnection(res.Status.NodeName) // errorについて考える
	if err != nil {
		_, err = a.nodeAPI.ReleaseStorage(context.Background(), &ppool.ReleaseStorageRequest{
			Name:        res.Status.NodeName,
			StorageName: res.Status.StorageName,
		})
		if err != nil {
			return nil, err // TODO: #89
		}

		log.Printf("Fail to dial to node: err=%v.", err.Error())
		return nil, grpc.Errorf(codes.Internal, "")
	}
	defer conn.Close()
	cli := NewVolumeAgentServiceClient(conn)

	v, err := cli.CreateVolumeAgentWithDownloading(context.Background(), &CreateVolumeAgentWithDownloadingRequest{
		Name:      req.Metadata.Name,
		Bytes:     req.Spec.LimitBytes,
		SourceUrl: req.SourceUrl,
	})
	if err != nil && status.Code(err) != codes.AlreadyExists {
		log.Printf("Fail to create volume on node '%s': err='%s'", "", err.Error()) // TODO: #89
		return nil, grpc.Errorf(codes.Internal, "")
	}

	res.Metadata.Annotations[AnnotationVolumePath] = v.Path
	res.Status.State = pprovisioning.VolumeStatus_AVAILABLE

	if err := a.dataStore.Apply(req.Metadata.Name, res); err != nil {
		_, err = a.nodeAPI.ReleaseStorage(context.Background(), &ppool.ReleaseStorageRequest{
			Name:        res.Status.NodeName,
			StorageName: res.Status.StorageName,
		})
		if err != nil {
			return nil, err // TODO: #89
		}

		log.Printf("[WARNING] Failed to apply data for db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to store '%s' for db, please retry or contact for the administrator of this cluster", req.Metadata.Name)
	}

	return res, nil
}

func (a *VolumeAPI) ListVolumes(ctx context.Context, req *pprovisioning.ListVolumesRequest) (*pprovisioning.ListVolumesResponse, error) {
	res := &pprovisioning.ListVolumesResponse{}
	f := func(s int) []proto.Message {
		res.Volumes = make([]*pprovisioning.Volume, s)
		for i := range res.Volumes {
			res.Volumes[i] = &pprovisioning.Volume{}
		}

		m := make([]proto.Message, s)
		for i, v := range res.Volumes {
			m[i] = v
		}

		return m
	}

	if err := a.dataStore.List(f); err != nil {
		return nil, grpc.Errorf(codes.Internal, "message:Failed to get from db\tgot:%v", err.Error())
	}
	if len(res.Volumes) == 0 {
		return nil, grpc.Errorf(codes.NotFound, "")
	}

	return res, nil
}

func (a *VolumeAPI) GetVolume(ctx context.Context, req *pprovisioning.GetVolumeRequest) (*pprovisioning.Volume, error) {
	res := &pprovisioning.Volume{}
	if err := a.dataStore.Get(req.Name, res); err != nil {
		log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Name)
	}
	if res == nil {
		return nil, grpc.Errorf(codes.NotFound, "")
	}

	return res, nil
}

func (a *VolumeAPI) UpdateVolume(ctx context.Context, req *pprovisioning.UpdateVolumeRequest) (*pprovisioning.Volume, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")

	// res := &pprovisioning.Volume{
	// 	Metadata: req.Metadata,
	// 	Spec:     req.Spec,
	// 	Status:   &pprovisioning.VolumeStatus{},
	// }

	// prev := &pprovisioning.Volume{}
	// if err := a.dataStore.Get(req.Metadata.Name, prev); err != nil {
	// 	log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
	// 	return nil, grpc.Errorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Metadata.Name)
	// }
	// var err error
	// res.Metadata.Version, err = datastore.CheckVersion(prev, req)
	// if err != nil {
	// 	return nil, grpc.Errorf(codes.InvalidArgument, "Failed to check version: %s", err.Error())
	// }

	// if prev.Spec.RequestBytes != req.Spec.RequestBytes {
	// 	return nil, grpc.Errorf(codes.Unimplemented, "Not supported change size: from='%s', to='%s'", prev.Spec.RequestBytes, req.Spec.RequestBytes)
	// }
	// if prev.Spec.LimitBytes != req.Spec.LimitBytes {
	// 	return nil, grpc.Errorf(codes.Unimplemented, "Not supported change size: from='%s', to='%s'", prev.Spec.LimitBytes, req.Spec.LimitBytes)
	// }

	// 	if err := a.dataStore.Apply(req.Metadata.Name, res); err != nil {
	// 		log.Printf("[WARNING] Failed to apply data for db: err='%s'", err.Error())
	// 		return nil, grpc.Errorf(codes.Internal, "Failed to store '%s' for db, please retry or contact for the administrator of this cluster", req.Metadata.Name)
	// 	}

	// return res, nil
}

func (a *VolumeAPI) DeleteVolume(ctx context.Context, req *pprovisioning.DeleteVolumeRequest) (*empty.Empty, error) {
	prev := &pprovisioning.Volume{}
	if err := a.dataStore.Get(req.Name, prev); err == nil {
		return nil, grpc.Errorf(codes.AlreadyExists, "Volume '%s' is already exists", req.Name)
	} else if status.Code(err) != codes.NotFound {
		log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Name)
	} else {
		return &empty.Empty{}, grpc.Errorf(codes.NotFound, "")
	}

	conn, err := a.nodeConnections.GetConnection(prev.Status.NodeName)
	if err != nil {
		log.Printf("[WARNING] Fail to dial to node: err=%v.", err.Error())
		return nil, grpc.Errorf(codes.Internal, "") // TODO: #89
	}
	if conn == nil {
		return nil, grpc.Errorf(codes.FailedPrecondition, "Node '%s' is not ready, so cannot delete: please wait a moment", prev.Status.NodeName)
	}
	defer conn.Close()
	cli := NewVolumeAgentServiceClient(conn)

	_, err = cli.DeleteVolumeAgent(context.Background(), &DeleteVolumeAgentRequest{Path: prev.Metadata.Annotations[AnnotationVolumePath]})
	if err != nil {
		log.Printf("Fail to delete volume on node, err:%v.", err.Error())
		return &empty.Empty{}, grpc.Errorf(codes.Internal, "Fail to delete volume on node") // TODO #89
	}

	d, err := a.dataStore.Delete(req.Name)
	if err != nil {
		return &empty.Empty{}, grpc.Errorf(codes.Internal, "message:Failed to delete from db.\tgot:%v", err.Error())
	}
	if d < 1 {
		return &empty.Empty{}, grpc.Errorf(codes.NotFound, "")
	}

	return &empty.Empty{}, nil
}

func (a *VolumeAPI) SetInuseVolume(ctx context.Context, req *pprovisioning.SetInuseVolumeRequest) (*pprovisioning.Volume, error) {
	res := &pprovisioning.Volume{}
	if err := a.dataStore.Get(req.Name, res); err != nil {
		log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Name)
	}

	res.Status.State = pprovisioning.VolumeStatus_IN_USE

	if err := a.dataStore.Apply(req.Name, res); err != nil {
		log.Printf("[WARNING] Failed to apply data for db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to store '%s' for db, please retry or contact for the administrator of this cluster", req.Name)
	}

	return res, nil
}

func (a *VolumeAPI) SetAvailableVolume(ctx context.Context, req *pprovisioning.SetAvailableVolumeRequest) (*pprovisioning.Volume, error) {
	res := &pprovisioning.Volume{}
	if err := a.dataStore.Get(req.Name, res); err != nil {
		log.Printf("[WARNING] Failed to get data from db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to get '%s' from db, please retry or contact for the administrator of this cluster", req.Name)
	}

	res.Status.State = pprovisioning.VolumeStatus_AVAILABLE

	if err := a.dataStore.Apply(req.Name, res); err != nil {
		log.Printf("[WARNING] Failed to apply data for db: err='%s'", err.Error())
		return nil, grpc.Errorf(codes.Internal, "Failed to store '%s' for db, please retry or contact for the administrator of this cluster", req.Name)
	}

	return nil, nil
}