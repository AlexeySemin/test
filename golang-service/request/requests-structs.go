package request

type CreateNews struct {
	Count int `validate:"required,max=500000"`
}
