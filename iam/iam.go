package iam

import (
	"context"

	admin "cloud.google.com/go/iam/admin/apiv1"
)

type IamContext struct {
	ctx   context.Context
	admin *admin.IamClient
}

// func (i *IamClient) ListServiceAccounts() error {
// 	ctx := context.Background()
// 	c, err := admin.NewIamClient(ctx)
// 	//it := &ServiceAccountIterator{}
// 	var resp *adminpb.ListServiceAccountsResponse
//
// 	if err != nil {
// 		return err
// 	}
// 	//req := &adminpb.ListServiceAccountsRequest{}
// 	req := &adminpb.ListServiceAccountsRequest{}
//
// 	it := c.ListServiceAccounts(ctx, req)
//
// 	for {
// 		resp, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		log.Println("User:", resp.GetName())
// 	}
// 	return nil
// }

// // NewIAMContext will create a new IAM context ready to use
// func NewIAMContext() (*IamContext, error) {
// 	ctx := context.Background()
// 	c, err := admin.NewIamClient(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IamContext{
// 		ctx:   ctx,
// 		admin: c,
// 	}, nil
// }

// func (i *IamContext) Close() error {
// 	return i.admin.Close()
// }
//
// func (i *IamContext) ListAccounts() error {
// 	req := &adminpb.ListServiceAccountsRequest{}
//
// 	it := i.admin.ListServiceAccounts(i.ctx, req)
// 	for {
// 		resp, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		log.Println("User:", resp.GetName())
// 	}
// 	return nil
// }
