package main

// Dictionary は辞書を表す
type Dictionary map[string]string

// DictionaryErr は辞書関連のエラー文字列
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	//ErrNotFound は文字列を検索できなかった場合のエラー内容
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	//ErrWordExists は文字列が既に存在する場合のエラー内容
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	//ErrWordDoesNotExist は文字の更新時に存在しない場合のエラー内容
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// Search は文字列を検索します
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add は辞書に文字を追加します
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update は辞書の文字を更新をする
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

//Delete は辞書の単語を削除します
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
