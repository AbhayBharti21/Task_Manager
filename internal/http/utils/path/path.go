package path

import (
	"net/http"
	"strconv"
	"strings"
)

func ExtractID(r *http.Request, segmentIndex int) (int, error) {
	path := r.URL.Path
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(pathParts) <= segmentIndex {
		return 0, ErrInvalidPath
	}

	id, err := strconv.Atoi(pathParts[segmentIndex])
	if err != nil {
		return 0, ErrInvalidID
	}

	return id, nil
}

func ExtractTaskID(r *http.Request) (int, error) {
	// For path like /api/tasks/123, we want segment index 2 (0-indexed: api, tasks, 123)
	return ExtractID(r, 2)
}
