package json

func Escape(src string) (string, error) {
	bytes, err := Marshal(src)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
