/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
type loginResult struct {
	Token       string        `json:"token,omitempty"`
	EcosystemID string        `json:"ecosystem_id,omitempty"`
	KeyID       string        `json:"key_id,omitempty"`
	Address     string        `json:"address,omitempty"`
	NotifyKey   string        `json:"notify_key,omitempty"`
	IsNode      bool          `json:"isnode,omitempty"`
	IsOwner     bool          `json:"isowner,omitempty"`
	IsVDE       bool          `json:"vde,omitempty"`
	Timestamp   string        `json:"timestamp,omitempty"`
	Roles       []rolesResult `json:"roles,omitempty"`
}

type rolesResult struct {
	RoleId   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}

type multiTxStatusResult struct {
	Results map[string]*txstatusResult `json:"results"`
}

type txstatusRequest struct {
	Hashes []string `json:"hashes"`
}

type txstatusError struct {
	Type  string `json:"type,omitempty"`
	Error string `json:"error,omitempty"`
	Id    string `json:"id,omitempty"`
}

type txstatusResult struct {
	BlockID string         `json:"blockid"`
	Message *txstatusError `json:"errmsg,omitempty"`
	Result  string         `json:"result"`
}

type contractField struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Optional bool   `json:"optional"`
}

type getContractResult struct {
	ID       uint32          `json:"id"`
	StateID  uint32          `json:"state"`
	Active   bool            `json:"active"`
	TableID  string          `json:"tableid"`
	WalletID string          `json:"walletid"`
	TokenID  string          `json:"tokenid"`
	Address  string          `json:"address"`
	Fields   []contractField `json:"fields"`
	Name     string          `json:"name"`
}

type sendTxResult struct {
	Hashes map[string]string `json:"hashes"`
}
