package customers

/**
 * Created by boy.sirichai on 23/8/2022 AD
 */
type Servicer interface {
	GetSymbols() (map[string]interface{}, error)
}

// Service type of announce
type Service struct {
}

func (s *Service) GetSymbols() (map[string]interface{}, error) {

	return map[string]interface{}{
		"id":   1,
		"name": "test",
	}, nil
}
