package auth

type UseCase interface {
	GetJwtUser(jwtToken string) (*ContextUser, error)
	GetContextUserKey() ContextUserKey
}
