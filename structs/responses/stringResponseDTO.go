package responses

type StringResponseFDTO struct {
	StatusCode int
	Value      *string
	StatusDesc string
}

type StringResponseDTO struct {
	StatusCode int
	Value      string
	StatusDesc string
}

type StringOriResponseDTO struct {
	StatusCode int
	Value      string
	StatusDesc string
}
