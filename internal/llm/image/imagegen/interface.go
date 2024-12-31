package imagegen

type Provider interface {
	GetModel() string

	GetURL() string

	GenerateImage(imageDescription string) ([]string, error)
}
