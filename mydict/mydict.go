package mydict

import (
	"errors"
)

type Dictionary map[string]string

// 에러 문구
var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("존재하지 않는 단어는 업데이트 할 수 없습니다.")
	errWordExists = errors.New("That word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] //exists true or false return

	if exists {
		return value, nil
	}
	return "", errNotFound
}

// dictionary 에 단어 추가하기
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExists
		}
		return nil
	*/
}

// 단어 업데이트
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

// 단어 삭제
func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantUpdate
	}

	return nil
}
