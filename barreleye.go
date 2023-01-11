package barreleye

// Barreleye -
type Barreleye struct {
	URL *string `json:"-"`
}

// NewBarreleye -
func NewBarreleye(
	url string,
) *Barreleye {
	return &Barreleye{
		URL: &url,
	}
}
