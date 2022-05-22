package entity

import "time"

/* Tenant Activity Validation */
type TenantActivity struct {
	ID                int64     `json:"id"`
	TENANT_ID         string    `json:"tenant_id"`
	CREATED_AT        time.Time `json:"created_at"`
	API_END_POINT     string    `json:"api_endpoint"`
	BYTES_TRANSFERRED int64     `json:"bytes_transferred"`
	IS_ERROR          int64     `json:"is_error"`
}

/* Tenant Activity Response */
type TenantActivityResponse struct {
	BYTES_TRANSFERRED int64 `json:"bytes_transferred"`
	IS_ERROR          int64 `json:"is_error"`
}
