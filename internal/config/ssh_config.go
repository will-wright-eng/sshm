package config

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

const (
    DelimiterStart = "### BEGIN MANAGED BLOCK ###"
    DelimiterEnd   = "### END MANAGED BLOCK ###"
)

type SSHConfigReader interface {
    Read() (string, error)
    Write(content string) error
}

type FileConfig struct {
    path string
}

func NewFileConfig(path string) *FileConfig {
    if strings.HasPrefix(path, "~/") {
        home, _ := os.UserHomeDir()
        path = filepath.Join(home, path[2:])
    }
    return &FileConfig{path: path}
}

func (f *FileConfig) Read() (string, error) {
    content, err := os.ReadFile(f.path)
    if err != nil {
        if os.IsNotExist(err) {
            return "", nil
        }
        return "", fmt.Errorf("reading config: %w", err)
    }
    return string(content), nil
}

func (f *FileConfig) Write(content string) error {
    dir := filepath.Dir(f.path)
    if err := os.MkdirAll(dir, 0700); err != nil {
        return fmt.Errorf("creating config directory: %w", err)
    }
    return os.WriteFile(f.path, []byte(content), 0600)
}
