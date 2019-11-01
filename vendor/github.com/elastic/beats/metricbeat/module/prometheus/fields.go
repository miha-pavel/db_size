// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package prometheus

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "prometheus", asset.ModuleFieldsPri, AssetPrometheus); err != nil {
		panic(err)
	}
}

// AssetPrometheus returns asset data.
// This is the base64 encoded gzipped contents of ./vendor/github.com/elastic/beats/metricbeat/module/prometheus.
func AssetPrometheus() string {
	return "eJyUkc1qwzAQhO9+ikG9hSQPoENfoYUeSwmKtba30R+7G0LeviQ2aQIttLppv9FoNNrgQGePJjWTTXTUDjC2RB7u9TZ0HRBJe+FmXIvHcwcAbxZMob2ERhGD1IyA71OgElvlYtsO0KmK7fpaBh49hpCUOkAoUVDyGMNFQ2ZcRvV4d6rJreEms+Y+OmBgSlH99d4nvEgkASs4tyoWimEioTVS2FNSnDgl5GD9hIFFbQ2bCEJqCEKI9bhPdPXaoIRM9w1sZ4/t6soBOzfyqPtP6m0ZzZvdTA50PlWJC/qhpsu6ayWTCfdL0t8yzKK/h7h70QPZ5dAal3GRuZX7Z84b2Tx81lcAAAD//3hXtDU="
}
