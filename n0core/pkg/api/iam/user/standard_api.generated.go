// Code generated by "standard_api"; DO NOT EDIT.

package auser

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/n0stack/n0stack/n0core/pkg/datastore"
	grpcutil "github.com/n0stack/n0stack/n0core/pkg/util/grpc"
	piam "github.com/n0stack/n0stack/n0proto.go/iam/v0"
	"github.com/n0stack/n0stack/n0proto.go/pkg/transaction"
	"google.golang.org/grpc/codes"
)

func ListUsers(ctx context.Context, req *piam.ListUsersRequest, ds datastore.Datastore) (*piam.ListUsersResponse, error) {
	res := &piam.ListUsersResponse{}
	f := func(s int) []proto.Message {
		res.Users = make([]*piam.User, s)
		for i := range res.Users {
			res.Users[i] = &piam.User{}
		}

		m := make([]proto.Message, s)
		for i, v := range res.Users {
			m[i] = v
		}

		return m
	}

	if err := ds.List(f); err != nil {
		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, "Failed to list from db, please retry or contact for the administrator of this cluster")
	}
	if len(res.Users) == 0 {
		return nil, grpcutil.WrapGrpcErrorf(codes.NotFound, "")
	}

	return res, nil
}

func GetUser(ctx context.Context, req *piam.GetUserRequest, ds datastore.Datastore) (*piam.User, error) {
	resourse := &piam.User{}
	if err := ds.Get(req.Name, resourse); err != nil {
		if datastore.IsNotFound(err) {
			return nil, grpcutil.WrapGrpcErrorf(codes.NotFound, err.Error())
		}

		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, datastore.DefaultErrorMessage(err))
	}

	return resourse, nil
}

func GetAndPendExistingUser(tx *transaction.Transaction, ds datastore.Datastore, name string) (*piam.User, error) {
	resource := &piam.User{}
	if err := ds.Get(name, resource); err != nil {
		if datastore.IsNotFound(err) {
			return nil, grpcutil.WrapGrpcErrorf(codes.NotFound, err.Error())
		}

		return nil, grpcutil.WrapGrpcErrorf(codes.Internal, datastore.DefaultErrorMessage(err))
	}

	if resource.State == piam.User_PENDING {
		return nil, grpcutil.WrapGrpcErrorf(codes.FailedPrecondition, "User %s is pending", name)
	}

	current := resource.State
	resource.State = piam.User_PENDING
	if err := ApplyUser(ds, resource); err != nil {
		return nil, err
	}
	resource.State = current
	tx.PushRollback("free optimistic lock", func() error {
		resource.State = current
		return ds.Apply(resource.Name, resource)
	})

	return resource, nil
}

func PendNewUser(tx *transaction.Transaction, ds datastore.Datastore, name string) error {
	resource := &piam.User{}
	if err := ds.Get(name, resource); err == nil {
		return grpcutil.WrapGrpcErrorf(codes.AlreadyExists, "User %s is already exists", name)
	} else if !datastore.IsNotFound(err) {
		return grpcutil.WrapGrpcErrorf(codes.Internal, datastore.DefaultErrorMessage(err))
	}

	resource.Name = name
	resource.State = piam.User_PENDING
	if err := ApplyUser(ds, resource); err != nil {
		return err
	}
	tx.PushRollback("free optimistic lock", func() error {
		return ds.Delete(name)
	})

	return nil
}

func DeleteUser(ds datastore.Datastore, name string) error {
	if err := ds.Delete(name); err != nil {
		return grpcutil.WrapGrpcErrorf(codes.Internal, "failed to delete User %s from db: err='%s'", name, err.Error())
	}

	return nil
}

func ApplyUser(ds datastore.Datastore, resource *piam.User) error {
	if err := ds.Apply(resource.Name, resource); err != nil {
		return grpcutil.WrapGrpcErrorf(codes.Internal, "failed to apply User %s to db: err='%s'", resource.Name, err.Error())
	}

	return nil
}
