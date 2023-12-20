package api

import (
	"io"
	"net/http"
	"path"

	ae "github.com/vault-thirteen/auxie/errors"
)

func (r *Repository) getTagsJSON() (data []byte, err error) {
	url := r.api.proto + "://" +
		path.Join(r.api.host, r.api.reposPath, r.owner, r.path, r.api.tagsPath)

	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer func() {
		derr := resp.Body.Close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	return data, nil
}
