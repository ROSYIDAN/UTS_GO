package dataDiri

type Informasi struct {
	Nama   string `json:"name"`
	NIM    int    `json:"nim"`
	Alamat string `json:"alamat"`
}

func NewInfo(name string, nim int, alamat string) *Informasi {
	return &Informasi{
		Nama:   name,
		NIM:    nim,
		Alamat: alamat,
	}
}
