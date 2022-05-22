package dto

import "time"

/* TENANT ACTIVITY*/
type TenantActivity struct {
	ID                int64     `json:"id"`
	TENANT_ID         string    `json:"tenant_id"`
	CREATED_AT        time.Time `json:"created_at"`
	API_END_POINT     string    `json:"api_endpoint"`
	BYTES_TRANSFERRED int       `json:"bytes_transferred"`
	IS_ERROR          int       `json:"is_error"`
}
