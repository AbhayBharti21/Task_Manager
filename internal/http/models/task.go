package types

type Task struct {
	TaskId      int    `json:"taskId"`
	OwnerId     int    `json:"ownerId"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

//type TaskManager struct {
//	Tasks   map[int]Task
//	OwnerId int
//	TaskId  int
//	Mu      *sync.Mutex
//}
