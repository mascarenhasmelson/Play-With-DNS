package model

type RecordRequest struct {
	Subdomain string            `json:"subdomain"`
	Typ       string            `json:"type"`
	TTL       uint32            `json:"ttl"`
	Values    map[string]string `json:"values"`
}

type RecordResponse struct {
	Subdomain  string            `json:"subdomain"`
	Type       string            `json:"type"`
	TTL        uint32            `json:"ttl"`
	Content    string            `json:"content"`
	// DomainName string            `json:"domain_name"`
	Values     []string `json:"values"`
}
