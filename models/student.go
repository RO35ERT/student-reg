

package models

import "gorm.io/gorm"

type Student struct {
    gorm.Model
    Name  string `json:"name" gorm:"size:30;not null"`
    Email string `json:"email" gorm:"size:50;unique;not null"`
}
