package handlers

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aliamerj/aircup/types"
	"github.com/labstack/echo/v5"
)

func (c Config) AddFilesEndPoints(g *echo.Group) {
	files := g.Group("/files")
	files.GET("", c.getAllFiles)
	files.GET("/content", c.streamFiles)
}

func (c *Config) getAllFiles(e *echo.Context) error {
	reqPath := e.QueryParamOr("path", "/")
	fullPath, cleanRelPath, err := c.resolve(reqPath)
	if err != nil {
		return withErr(e, err)
	}

	info, err := os.Stat(fullPath)
	if err != nil {
		return withErr(e, err)
	}

	if !info.IsDir() {
		return withErr(e, fmt.Errorf("path is not a directory"))
	}

	items, err := os.ReadDir(fullPath)
	if err != nil {
		return withErr(e, err)
	}

	entries := make([]types.FileEntry, 0, len(items))

	for _, item := range items {

		if item.Type()&os.ModeSymlink != 0 {
			continue
		}

		itemInfo, err := item.Info()
		if err != nil {
			return withErr(e, err)
		}

		itemRelPath := filepath.ToSlash(filepath.Join(cleanRelPath, item.Name()))
		if !strings.HasPrefix(itemRelPath, "/") {
			itemRelPath = "/" + itemRelPath
		}

		size := itemInfo.Size()
		if item.IsDir() {
			size = 0
		}

		entries = append(entries, types.FileEntry{
			Name:       item.Name(),
			Path:       itemRelPath,
			IsDir:      item.IsDir(),
			Size:       size,
			ModifiedAt: itemInfo.ModTime(),
		})

	}

	return e.JSON(http.StatusOK, types.FileListResponse{
		Path:    cleanRelPath,
		Entries: entries,
	})
}

func (c *Config) streamFiles(e *echo.Context) error {
	reqPath := e.QueryParam("path")
	if reqPath == "" {
		return withErr(e, fmt.Errorf("path is required"))
	}

	fullPath, _, err := c.resolve(reqPath)
	if err != nil {
		return withErr(e, err)
	}

	info, err := os.Lstat(fullPath)
	if err != nil {
		return withErr(e, err)
	}

	if info.Mode()&os.ModeSymlink != 0 {
		return withErr(e, fmt.Errorf("symlinks are not allowed"))
	}

	if info.IsDir() {
		return withErr(e, fmt.Errorf("path is a directory"))
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return withErr(e, err)
	}
	defer file.Close()

	filename := info.Name()

	contentType := mime.TypeByExtension(filepath.Ext(fullPath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	header := e.Response().Header()
	header.Set(echo.HeaderContentType, contentType)
	header.Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filename))
	header.Set("X-Content-Type-Options", "nosniff")

	http.ServeContent(e.Response(), e.Request(), filename, info.ModTime(), file)
	return nil
}

func (c *Config) resolve(relPath string) (fullPath string, cleanRelPath string, err error) {
	if relPath == "" {
		relPath = "/"
	}

	rootAbs, err := filepath.Abs(c.Config.Root)
	if err != nil {
		return "", "", fmt.Errorf("resolve root: %w", err)
	}

	cleanRelPath = filepath.Clean("/" + relPath)
	trimmed := strings.TrimPrefix(cleanRelPath, "/")

	fullPath = filepath.Join(rootAbs, trimmed)
	fullPath, err = filepath.Abs(fullPath)
	if err != nil {
		return "", "", fmt.Errorf("resolve path: %w", err)
	}

	if fullPath != rootAbs && !strings.HasPrefix(fullPath, rootAbs+string(os.PathSeparator)) {
		return "", "", fmt.Errorf("path escapes configured root")
	}

	return fullPath, cleanRelPath, nil
}
