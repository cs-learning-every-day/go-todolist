package serializer

import "todo-list/model"

type Task struct {
	ID        uint   `json:"id"`      // 任务ID
	Title     string `json:"title"`   // 题目
	Content   string `json:"content"` // 内容
	View      uint64 `json:"view"`    // 浏览量
	Status    int    `json:"status"`  // 状态(0未完成，1已完成)
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		View:      item.View(),
		CreatedAt: item.CreatedAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return
}
