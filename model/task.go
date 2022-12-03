package model

import (
	"strconv"
	"todo-list/cache"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}

func (t *Task) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(t.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (t *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(t.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(t.ID)))
}
