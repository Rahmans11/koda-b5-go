package bio

import "errors"

func CreateBio(
	nama string,
	foto string,
	email string,
	umur int,
	noTelp int,
	status bool,
	riwayat map[string]string) (Profile, error) {

	if nama == "" {
		return Profile{}, errors.New("nama tidak boleh kosong")
	}

	if email == "" {
		return Profile{}, errors.New("email tidak boleh kosong")
	}

	if umur <= 0 {
		return Profile{}, errors.New("umur harus diatas 0")
	}

	userInput := Profile{
		Nama:    nama,
		Foto:    foto,
		Email:   email,
		Umur:    umur,
		NoTelp:  noTelp,
		Status:  status,
		Riwayat: riwayat,
	}

	return userInput, nil
}
