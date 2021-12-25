package dao

import (
	"gorm.io/gorm"
)

type Dao[ID IDType, Entity IEntity] struct {
	DB *gorm.DB
}

func (d *Dao[ID, Entity]) Insert(entity Entity) Result {
	tx := d.DB.Create(entity)

	return Result{Error: tx.Error, RowsAffected: tx.RowsAffected}
}

func (d *Dao[ID, Entity]) Inserts(entities []Entity, batchSize int) Result {
	tx := d.DB.CreateInBatches(entities, batchSize)

	return Result{Error: tx.Error, RowsAffected: tx.RowsAffected}
}

func (d *Dao[ID, Entity]) Delete(id ID) Result {
	var e Entity
	tx := d.DB.Delete(&e, id)

	return Result{Error: tx.Error, RowsAffected: tx.RowsAffected}
}

func (d *Dao[ID, Entity]) Deletes(ids []ID) Result {
	var e Entity
	tx := d.DB.Delete(&e, ids)

	return Result{Error: tx.Error, RowsAffected: tx.RowsAffected}
}

func (d *Dao[ID, Entity]) Update(entity *Entity, wheres []Where, selects ...string) Result {
	var tx *gorm.DB
	if len(wheres) == 0 {
		c := d.DB.Model(entity)
		if len(selects) == 0 {
			tx = c.Updates(entity)
		} else {
			tx = c.Select(selects).Updates(entity)
		}
	} else {
		var e Entity
		c := d.DB.Model(&e)
		for _, v := range wheres {
			c = c.Where(v.Query, v.Args...)
		}
		if len(selects) == 0 {
			tx = c.Updates(entity)
		} else {
			tx = c.Select(selects).Updates(entity)
		}
	}
	return Result{Error: tx.Error, RowsAffected: tx.RowsAffected}
}

func (d *Dao[ID, Entity]) Get(id ID) (Entity, Result) {
	var res Entity
	tx := d.DB.First(&res, id)
	return res, Result{Error: tx.Error, RowsAffected: tx.RowsAffected}
}

func (d *Dao[ID, Entity]) Find(wheres []Where, selects ...string) []Entity {
	res := make([]Entity, 0, 100)

	c := d.DB

	if len(selects) > 0 {
		c = c.Select(selects)
	}

	tx := c.Where("1 = 1")
	for _, v := range wheres {
		tx = tx.Where(v)
	}
	tx.Find(&res)
	return res
}

func NewDao[T IEntity](db *gorm.DB) IDao[int64, T] {
	return &Dao[int64, T]{
		DB: db,
	}
}
