package users

type Service struct {
	repository *Repository
}

func newService(repo *Repository) *Service {
	return &Service{repository: repo}
}
