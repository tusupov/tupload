package handle

type (

	Result struct {
		Image	string		`json:"img"`
		Thumb	string		`json:"thumb"`
		Err		string		`json:"err"`
	}

	Results struct {
		Results	[]Result	`json:"results"`
	}

)

func (r *Results) Add(img, thumb string) {
	r.Results = append(
		r.Results,
		Result{
			Image: img,
			Thumb: thumb,
		},
	)
}

func (r *Results) AddError(img string, err error) {
	r.Results = append(
		r.Results,
		Result{
			Image: img,
			Err: err.Error(),
		},
	)
}

func NewResult() Results {
	return Results{
		make([]Result, 0),
	}
}

