package interfaces

type DetectChange interface {
	isFileChange(string) bool
}
