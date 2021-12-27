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

package beat

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "beat", asset.ModuleFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded zlib format compressed contents of module/beat.
func AssetBeat() string {
	return "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"
}
