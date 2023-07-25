// Copyright 2022 Chainguard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package confloader

import (
	"os"

	"gopkg.in/yaml.v3"
)

type LoadedConfig struct {
	Root    yaml.Node
}

type Option func(ctx *LoadedConfig) error

func LoadConfig(configFile string) (*LoadedConfig, error) {
	cfg := LoadedConfig{}

	configData, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(configData, &cfg.Root); err != nil {
		return nil, err
	}

	return &cfg, nil
}
