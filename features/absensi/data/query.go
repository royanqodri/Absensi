package data

import (
	"Absensi-App/features/absensi"

	"gorm.io/gorm"
)

type absensiQuery struct {
	db *gorm.DB
}

// SelectUserById implements absensi.AbsensiDataInterface
func (*absensiQuery) SelectUserById(idUser string) (absensi.PenggunaEntity, error) {
	panic("unimplemented")
}
