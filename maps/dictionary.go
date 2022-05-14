package maps

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNoItemInDict      = DictionaryErr("item not found in dict")
	ErrItemAlreadyInDict = DictionaryErr("item is already in dict")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(searchKey string) (string, error) {
	if val, in := d[searchKey]; in {
		return val, nil
	}
	return "", ErrNoItemInDict
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNoItemInDict:
		d[key] = val
		return nil
	case nil:
		return ErrItemAlreadyInDict
	default:
		return err
	}
}

func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNoItemInDict:
		return ErrNoItemInDict
	case nil:
		d[key] = val
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNoItemInDict:
		return ErrNoItemInDict
	case nil:
		delete(d, key)
		return nil
	default:
		return err
	}
}
