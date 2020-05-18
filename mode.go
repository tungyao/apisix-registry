package main

// easyjson:json
type _nodes struct {
	Addr  string `json:"addr"`
	Weigh int32  `json:"weigh"`
}

// easyjson:json
type Upstream struct {
	Desc  string    `json:"string"`
	Nodes []*_nodes `json:"nodes"`
	Name  string    `json:"name"`
}
