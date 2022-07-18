package tggo

import (
	"encoding/json"
	"fmt"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type TypeApiResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
}

func (t *TypeApiResponse) decode(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *TypeApiResponse) isApiError() bool {
	return t.ErrorCode != 0
}

func (t *TypeApiResponse) generateApiError() error {
	return fmt.Errorf("APIERROR{ http_code: %d, description: %s}", t.ErrorCode, t.Description)
}

type TypeGetMe struct {
	Id                     int    `json:"id"`
	IsBot                  bool   `json:"is_bot"`
	FirstName              string `json:"first_name"`
	Username               string `json:"username"`
	CanJoinGroups          bool   `json:"can_join_groups"`
	CanReadAllGroupMessage bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries  bool   `json:"supports_inline_queries"`
}

func (t *TypeGetMe) decode(data []byte) error {
	resp := TypeApiResponse{}
	resp.decode(data)

	if resp.isApiError() {
		return resp.generateApiError()
	}

	err := json.Unmarshal(resp.Result, t)
	if err != nil {
		return err
	}

	return nil
}
