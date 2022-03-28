package keyring

type keychain struct {
	path    string
	service string

	passwordFunc PromptFunc

	isSynchronizable         bool
	isAccessibleWhenUnlocked bool
	isTrusted                bool
}

func (k *keychain) Get(key string) (item Item, err error) {
	return
}

func (k *keychain) GetMetadata(key string) (md Metadata, err error) {
	return
}

func (k *keychain) Set(item Item) error {
	return nil
}

func (k *keychain) Remove(key string) error {
	return nil
}

func (k *keychain) Keys() (accountNames []string, err error) {
	return
}
