package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var Envs *Config

func init() {
	cfg, err := Load()
	if err != nil {
		panic(err.Error())
	}
	Envs = cfg
}

type Config struct {
	Weight               float64 `yaml:"weight"`
	OfficeHours          int     `yaml:"office_hours"`
	LunchIntervalMinutes int     `yaml:"lunch_interval_minutes"`
	LunchIntervalStart   int     `yaml:"lunch_interval_start"`
}

func Load() (*Config, error) {
	filename := "config.yaml"
	config := &Config{}

	basicPath := "./"
	if _, err := os.Stat(filepath.Join(basicPath, filename)); os.IsNotExist(err) {
		basicPath, err = getProjectRoot()
		if err != nil {
			return nil, err
		}
	}

	path := filepath.Join(basicPath, filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("arquivo de configutação [%s] não encontrado", filename)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("erro ao parsear configuração: %w", err)
	}

	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("erro de validação da configuração: %w", err)
	}

	return config, nil
}

func (c *Config) validate() error {
	if c.Weight <= 0 {
		return fmt.Errorf("o peso deve ser maior que zero, valor fornecido: %f", c.Weight)
	}
	if c.OfficeHours <= 0 || c.OfficeHours > 24 {
		return fmt.Errorf("as horas de expediente devem ser entre 1 e 24, valor fornecido: %d", c.OfficeHours)
	}
	if c.LunchIntervalMinutes <= 0 || c.LunchIntervalMinutes > 180 {
		return fmt.Errorf("o intervalo de almoço deve ser entre 1 e 180 minutos, valor fornecido: %d", c.LunchIntervalMinutes)
	}
	if c.LunchIntervalStart < 0 || c.LunchIntervalStart > 24 {
		return fmt.Errorf("o horário de início do almoço deve ser entre 0 e 24, valor fornecido: %d", c.LunchIntervalStart)
	}

	return nil
}

func getProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}

	return "", fmt.Errorf("não foi possível encontrar a raiz do projeto")
}
