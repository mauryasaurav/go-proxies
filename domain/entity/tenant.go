package entity

/* Tenant Validation*/
type TenantSchema struct {
	ID   int64  `json:"id"`
	NAME string `json:"name"`
}

type Tenant struct {
	ID   int64  `json:"id"`
	NAME string `json:"name"`
}
