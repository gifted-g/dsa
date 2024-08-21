package main

func IsUnique(str string) bool {
	ht := NewHashTable()

	for _, v := range str {
		_, exist := ht.Get(string(v))

		if exist {
			return false
		}

		ht.Put(string(v), "exist")
	}

	return true
}
