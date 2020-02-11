package messagequeue

import (
	"sns/models"
)

type QueueInterface interface {
	Push(string message, string actionType) error
}