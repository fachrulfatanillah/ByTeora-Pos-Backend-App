package request

type CreateUserRequest struct {
    Email        string `json:"email"`
    Password     string `json:"password"`
    NamaDepan    string `json:"nama_depan"`
    NamaBelakang string `json:"nama_belakang"`
}