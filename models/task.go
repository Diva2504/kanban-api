package models

import "gorm.io/gorm"

type Task struct {
  gorm.Model
  Title string
  Description string
  Modelss string
  Status bool
}
