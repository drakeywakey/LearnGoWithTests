package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	WordExistsError     = DictionaryErr("tried to add an existing word")
	ErrWordDoesNotExist = DictionaryErr("tried to update a word that doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dict Dictionary) Search(key string) (string, error) {
	def, ok := dict[key]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (dict Dictionary) Add(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		dict[key] = value
	case nil:
		return WordExistsError
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Update(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dict[key] = value
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Delete(key string) {
	delete(dict, key)
}
