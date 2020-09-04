package reserving

//Triangle : Defines the loss triangle structure used for paid and incurred development
type Triangle struct {
	Index            []byte   `json:"triangle_index"`
	Columns          []string `json:"triangle_columns"`
	Origin           []byte   `json:"triangle_origin"`
	Development      []byte   `json:"triangle_development"`
	KeyLabels        []byte   `json:"triangle_key_labels"`
	Valuation        []byte   `json:"triangle_valuation_dates"`
	OriginGrain      []string `json:"triangle_origin_grain"`
	DevelopmentGrain []string `json:"triangle_development_grain"`
	Shape            [4]int   `json:"triangle_shape"`
}
