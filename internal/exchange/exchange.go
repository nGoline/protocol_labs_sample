package exchange

import "github.com/jinzhu/gorm"

// Worker is a worker interface for exchanges
type Worker interface {
	Init() *gorm.DB
	SyncData(db *gorm.DB)
	GetName() string
}
