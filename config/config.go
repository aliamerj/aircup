package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

type Config struct {
	Addr string `json:"addr"`
	Root string `json:"root"`
}

const (
	AppDirName     = "Aircup"
	ConfigFileName = "config.json"
	DefaultAddr    = "0.0.0.0:8080"
)

func Parse(importPath string, overrides Config) (*Config, error) {
	configPath, err := resolveAppConfigPath()
	if err != nil {
		return nil, err
	}

	cfg, _, err := loadOrCreate(configPath)
	if err != nil {
		return nil, err
	}

	if importPath != "" {
		imported, err := loadExisting(filepath.Clean(importPath))
		if err != nil {
			return nil, err
		}

		if imported.Addr != "" {
			cfg.Addr = imported.Addr
		}
		if imported.Root != "" {
			cfg.Root = imported.Root
		}
	}

	if overrides.Addr != "" {
		cfg.Addr = overrides.Addr
	}
	if overrides.Root != "" {
		cfg.Root = overrides.Root
	}

	if err := validate(cfg); err != nil {
		return nil, err
	}

	if err := save(configPath, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func resolveAppConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home dir: %w", err)
	}

	return filepath.Join(home, ".local", "share", AppDirName, ConfigFileName), nil
}

func loadOrCreate(path string) (*Config, bool, error) {
	data, err := os.ReadFile(path)
	if err == nil {
		var cfg Config
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, false, fmt.Errorf("parse config %q: %w", path, err)
		}
		return &cfg, false, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return nil, false, fmt.Errorf("read config %q: %w", path, err)
	}

	cfg := Config{
		Addr: DefaultAddr,
		Root: "",
	}

	if err := save(path, &cfg); err != nil {
		return nil, false, err
	}

	return &cfg, true, nil
}

func loadExisting(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read import config %q: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse import config %q: %w", path, err)
	}

	return &cfg, nil
}

func save(path string, cfg *Config) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create config dir %q: %w", dir, err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write config %q: %w", path, err)
	}

	return nil
}

func validate(cfg *Config) error {
	if cfg.Addr == "" {
		return errors.New("config addr is required")
	}
	if cfg.Root == "" {
		return errors.New("config root is required")
	}

	info, err := os.Stat(cfg.Root)
	if err != nil {
		return fmt.Errorf("stat root %q: %w", cfg.Root, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("root %q is not a directory", cfg.Root)
	}

	ln, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return fmt.Errorf("cannot bind addr %q: %w", cfg.Addr, err)
	}
	_ = ln.Close()

	return nil
}
