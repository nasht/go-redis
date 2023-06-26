package service


type StorageService struct {
	store map[string]string
}


func NewStorageService() StorageService {
	return StorageService { store: make(map[string]string) }
}
func (s *StorageService)Put(key string, value string) {
	s.store[key] = value
}

func (s *StorageService)Get(key string) string {
	return s.store[key]
}
