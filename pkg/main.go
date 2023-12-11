package pkg

import (
	"ace/model"
)

var Upload model.Upload

func Init(u model.Upload) {
	Upload = u
}
