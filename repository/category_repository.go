package repository

import "golang-pzn-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
