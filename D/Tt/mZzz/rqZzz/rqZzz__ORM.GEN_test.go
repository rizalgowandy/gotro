package rqZzz

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	`testing`
	`github.com/stretchr/testify/assert`
	`github.com/kokizzu/gotro/D/Tt`
)

func TestGeneratedZzzOrmHelpers(t *testing.T) {
	q := NewZzz(nil)
	assert.NotNil(t, q)
	assert.NotEmpty(t, q.SpaceName())
	assert.NotEmpty(t, q.SqlTableName())
	q.Id = uint64(1)
	q.CreatedAt = int64(1)
	q.Coords = []any{1.1, 2.2}
	q.Name = "sample"
	q.HeightMeter = 1.5
	q.Counter = uint64(1)
	q.IsActive = true
	arr := q.ToArray()
	assert.Len(t, arr, 7)
	assert.NotNil(t, q.ToUpdateArray())
	decoded := (&Zzz{}).FromArray(arr)
	assert.Equal(t, q.Id, decoded.Id)
	assert.Equal(t, q.CreatedAt, decoded.CreatedAt)
	assert.Equal(t, q.Coords, decoded.Coords)
	assert.Equal(t, q.Name, decoded.Name)
	assert.Equal(t, q.HeightMeter, decoded.HeightMeter)
	assert.Equal(t, q.Counter, decoded.Counter)
	assert.Equal(t, q.IsActive, decoded.IsActive)
	decoded2 := (&Zzz{}).FromUncensoredArray(arr)
	assert.Equal(t, q.Id, decoded2.Id)
	assert.Equal(t, q.CreatedAt, decoded2.CreatedAt)
	assert.Equal(t, q.Coords, decoded2.Coords)
	assert.Equal(t, q.Name, decoded2.Name)
	assert.Equal(t, q.HeightMeter, decoded2.HeightMeter)
	assert.Equal(t, q.Counter, decoded2.Counter)
	assert.Equal(t, q.IsActive, decoded2.IsActive)
	assert.Equal(t, 0, q.IdxId())
	assert.Equal(t, `"id"`, q.SqlId())
	assert.Equal(t, 1, q.IdxCreatedAt())
	assert.Equal(t, `"createdAt"`, q.SqlCreatedAt())
	assert.Equal(t, 2, q.IdxCoords())
	assert.Equal(t, `"coords"`, q.SqlCoords())
	assert.Equal(t, 3, q.IdxName())
	assert.Equal(t, `"name"`, q.SqlName())
	assert.Equal(t, 4, q.IdxHeightMeter())
	assert.Equal(t, `"heightMeter"`, q.SqlHeightMeter())
	assert.Equal(t, 5, q.IdxCounter())
	assert.Equal(t, `"counter"`, q.SqlCounter())
	assert.Equal(t, 6, q.IdxIsActive())
	assert.Equal(t, `"isActive"`, q.SqlIsActive())
	_, ok := ZzzFieldTypeMap[`id`]
	assert.True(t, ok)
	assert.NotEmpty(t, q.UniqueIndexId())
	assert.NotEmpty(t, q.UniqueIndexName())
	assert.Equal(t, `coords`, q.SpatialIndexCoords())
}

func TestGeneratedZzzDbMethodsPanic(t *testing.T) {
	q := NewZzz(&Tt.Adapter{})
	q.Id = uint64(1)
	q.CreatedAt = int64(1)
	q.Coords = []any{1.1, 2.2}
	q.Name = "sample"
	q.HeightMeter = 1.5
	q.Counter = uint64(1)
	q.IsActive = true
	assert.Panics(t, func() { _ = q.FindById() })
	assert.Panics(t, func() { _ = q.FindByName() })
	assert.Panics(t, func() { _ = q.FindOffsetLimit(0, 1, "") })
	assert.Panics(t, func() { _, _ = q.FindArrOffsetLimit(0, 1, "") })
	assert.Panics(t, func() { _ = q.Total() })
}

