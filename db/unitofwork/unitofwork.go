package unitofwork

type IUnitOfWork interface {
	Begin()
	Commit()
	Rollback()
}
