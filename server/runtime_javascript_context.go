// Copyright 2018 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/dop251/goja"
)

const (
	__RUNTIME_JAVASCRIPT_CTX_ENV              = "env"
	__RUNTIME_JAVASCRIPT_CTX_MODE             = "execution_mode"
	__RUNTIME_JAVASCRIPT_CTX_NODE             = "node"
	__RUNTIME_JAVASCRIPT_CTX_QUERY_PARAMS     = "query_params"
	__RUNTIME_JAVASCRIPT_CTX_USER_ID          = "user_id"
	__RUNTIME_JAVASCRIPT_CTX_USERNAME         = "username"
	__RUNTIME_JAVASCRIPT_CTX_VARS             = "vars"
	__RUNTIME_JAVASCRIPT_CTX_USER_SESSION_EXP = "user_session_exp"
	__RUNTIME_JAVASCRIPT_CTX_SESSION_ID       = "session_id"
	__RUNTIME_JAVASCRIPT_CTX_CLIENT_IP        = "client_ip"
	__RUNTIME_JAVASCRIPT_CTX_CLIENT_PORT      = "client_port"
	__RUNTIME_JAVASCRIPT_CTX_MATCH_ID         = "match_id"
	__RUNTIME_JAVASCRIPT_CTX_MATCH_NODE       = "match_node"
	__RUNTIME_JAVASCRIPT_CTX_MATCH_LABEL      = "match_label"
	__RUNTIME_JAVASCRIPT_CTX_MATCH_TICK_RATE  = "match_tick_rate"
)

func NewRuntimeJsContext(r *goja.Runtime, node string, env goja.Value, mode RuntimeExecutionMode, queryParams map[string][]string, sessionExpiry int64, userID, username string, vars map[string]string, sessionID, clientIP, clientPort string) *goja.Object {
	ctxObj := r.NewObject()
	ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_NODE, node)
	ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_ENV, env)
	ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_MODE, mode)
	if queryParams != nil {
		ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_QUERY_PARAMS, queryParams)
	}
	if sessionExpiry != 0 {
		ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_USER_SESSION_EXP, sessionExpiry)
	}
	if userID != "" {
		ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_USER_ID, userID)
	}
	if username != "" {
		ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_USERNAME, username)
	}
	if vars != nil {
		ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_VARS, vars)
	}
	if sessionID != "" {
		ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_SESSION_ID, sessionID)
	}
	if clientIP != "" {
		ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_CLIENT_IP, clientIP)
	}
	if clientPort != "" {
		ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_CLIENT_PORT, clientPort)
	}

	return ctxObj
}

func NewRuntimeJsInitContext(r *goja.Runtime, node string, env map[string]string) *goja.Object {
	ctxObj := r.NewObject()
	ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_NODE, node)
	ctxObj.Set(__RUNTIME_JAVASCRIPT_CTX_ENV, env)

	return ctxObj
}
