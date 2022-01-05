// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servicetest // import "go.opentelemetry.io/collector/service/servicetest"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmapprovider"
	"go.opentelemetry.io/collector/config/configunmarshaler"
	"go.opentelemetry.io/collector/service/internal/configprovider"
)

// LoadConfig loads a config.Config  from file, and does NOT validate the configuration.
func LoadConfig(fileName string, factories component.Factories) (*config.Config, error) {
	// Read yaml config from file
	cp, err := configprovider.NewExpand(configmapprovider.NewFile(fileName)).Retrieve(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	// Unmarshal the config using the given factories.
	m, err := cp.Get(context.Background())
	if err != nil {
		return nil, err
	}
	return configunmarshaler.NewDefault().Unmarshal(m, factories)
}

// LoadConfigAndValidate loads a config from the file, and validates the configuration.
func LoadConfigAndValidate(fileName string, factories component.Factories) (*config.Config, error) {
	cfg, err := LoadConfig(fileName, factories)
	if err != nil {
		return nil, err
	}
	return cfg, cfg.Validate()
}