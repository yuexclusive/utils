package dao

type IDType interface {
	string | int64
}

type IEntity interface {
}

type Result struct {
	Error        error
	RowsAffected int64
}

type Where struct {
	Query interface{}
	Args  []interface{}
}

type IDao[ID IDType, Entity IEntity] interface {
	Insert(entity Entity) Result
	Inserts(entities []Entity, batchSize int) Result
	Delete(id ID) Result
	Deletes(ids []ID) Result
	Update(entity *Entity, wheres []Where, selects ...string) Result
	Get(id ID) (Entity, Result)
	Find(wheres []Where, selects ...string) []Entity
}
