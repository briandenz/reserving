package triangle

//Triangle : Defines the loss triangle structure used for paid and incurred development
type Triangle struct {
	Index               []byte   `json:"triangle_index"`
	Columns             []string `json:"triangle_columns"`
	Origin              []byte   `json:"triangle_origin"`
	Development         []byte   `json:"triangle_development"`
	KeyLabels           []byte   `json:"triangle_key_labels"`
	Valuation           []byte   `json:"triangle_valuation_dates"`
	OriginGrain         []string `json:"triangle_origin_grain"`
	DevelopmentGrain    []string `json:"triangle_development_grain"`
	Shape               [4]int   `json:"triangle_shape"`
	LatestValuationDate []byte   `json:"triangle_valuation_date"`
	LatestDiagonal      []int    `json:"triangle_latest_diagonal"`
	IsCumulative        bool     `json:"triangle_is_cumulative"`
	IsLDF               bool     `json:"triangle_is_ldf"`
	IsUltimate          bool     `json:"triangle_is_ultimate"`
	IsFull              bool     `json:"triangle_is_full"`
	IsVal               bool     `json:"triangle_is_val"`
}
