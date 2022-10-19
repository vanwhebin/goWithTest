package maps

type Dict map[string]string

var (
	ErrNotFoundMapValue        = DictErr("not found value")
	ErrExistedMapValue         = DictErr("cannot assign value, while it already existed")
	ErrKeyDoesNotExistMapValue = DictErr("cannot assign value, while it already existed")
)

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

func (d Dict) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFoundMapValue
	}
	return value, nil
}

func (d Dict) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFoundMapValue:
		d[key] = value
	case nil:
		return ErrExistedMapValue
	default:
		return err
	}
	return nil
}

func (d Dict) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFoundMapValue:
		return ErrKeyDoesNotExistMapValue
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dict) Delete(key string) {
	delete(d, key)
}
