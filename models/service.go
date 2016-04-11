package models

type Plan struct {
	IPs   []string `json:"ips"`
	Users []User   `json:"users"`
}

type User struct {
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Meta     map[string]interface{} `json:"meta"`
}

type Service struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ExternalIP string `json:"external_ip"`
	InternalIP string `json:"internal_ip`
	Plan       Plan   `json:"plan"`
	Started bool
}