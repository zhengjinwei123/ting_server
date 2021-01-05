package userservice

import (
	"app/yut/utils/fileutils"
	"errors"
	"path/filepath"
)

func helperGetImagePath(username string, image_name string) string {
	return "/upload/"+username+"/"+image_name
}

func helperDelUserImage(username string, image_name string) error {
	p, _ := filepath.Abs(filepath.Dir("./public/"))
	file_path := p + helperGetImagePath(username, image_name) //"/upload/" + username + "/" + request.ImageName

	if !fileutils.Exists(file_path) {
		return errors.New(file_path + " file not exists")
	}

	if err := fileutils.RmFile(file_path); err != nil {
		return err
	}

	return nil
}