// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admin

import (
	"context"

	"github.com/openimsdk/chat/pkg/protocol/admin"
)

func (o *adminServer) CreateToken(ctx context.Context, req *admin.CreateTokenReq) (*admin.CreateTokenResp, error) {
	token, err := o.Token.CreateToken(req.UserID, req.UserType)
	if err != nil {
		return nil, err
	}
	err = o.Database.CacheToken(ctx, req.UserID, token)
	if err != nil {
		return nil, err
	}
	return &admin.CreateTokenResp{
		Token: token,
	}, nil
}

func (o *adminServer) ParseToken(ctx context.Context, req *admin.ParseTokenReq) (*admin.ParseTokenResp, error) {
	userID, userType, err := o.Token.GetToken(req.Token)
	if err != nil {
		return nil, err
	}
	return &admin.ParseTokenResp{
		UserID:   userID,
		UserType: userType,
	}, nil
}

func (o *adminServer) GetUserToken(ctx context.Context, req *admin.GetUserTokenReq) (*admin.GetUserTokenResp, error) {
	tokensMap, err := o.Database.GetTokens(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	return &admin.GetUserTokenResp{TokensMap: tokensMap}, nil
}
