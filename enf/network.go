package enf

type NetworkService service

type NetworkRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type Network struct {
	Name        *string `json:"name"`
	Network     *string `json:"network"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}
