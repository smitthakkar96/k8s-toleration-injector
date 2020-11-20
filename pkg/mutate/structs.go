package mutate


type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type Toleration struct {
	Key string `json:"key"`
	Value string `json:"value"`
	Effect string `json:"effect"`
	Operator string `json:"operator"`
}
