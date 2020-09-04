package reserving

//Triangle : Defines the loss triangle structure used for paid and incurred development
type Triangle struct {
	Origin      []byte `json:"triangle_origin"`
	Development []byte `json:"triangle_development"`

	Cumulative bool `json:"triangle_cum_incr"`
}
