package accounts

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"yatter-backend-go/app/handler/auth"

	"github.com/google/uuid"
)

type UpdateRequest struct {
	DisplayName string
	Note        string
	Avatar      string
	Header      string
}

// Handle request for `POST /v1/accounts/update_credentials`
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(1024)

	account := auth.AccountOf(r)

	account.DisplayName = &r.MultipartForm.Value["display_name"][0]
	account.Note = &r.MultipartForm.Value["note"][0]

	if path, err := saveFileFromFormName(w, r, "avatar", "files/images/avatar/"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if path != nil {
		account.Avatar = path
	}

	if path, err := saveFileFromFormName(w, r, "header", "files/images/header/"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if path != nil {
		account.Header = path
	}

	account, err := h.ar.UpdateCredentials(ctx, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func saveFileFromFormName(w http.ResponseWriter, r *http.Request, name string, folder_path string) (*string, error) {
	fileHeader := r.MultipartForm.File[name][0]
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	path, err := saveFile(folder_path, &file)
	if err != nil {
		return nil, err
	}

	return path, nil
}

func saveFile(folder_path string, file *multipart.File) (*string, error) {

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	path := folder_path + uid.String() + ".jpg"
	output_file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer output_file.Close()

	_, err = io.Copy(output_file, *file)
	if err != nil {
		return nil, err
	}

	return &path, nil
}
