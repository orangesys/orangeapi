// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

import (
	"reflect"
	"testing"
)

func Test_getConfigMap(t *testing.T) {
	type args struct {
		namespace string
		name      string
	}
	tests := []struct {
		name    string
		args    args
		want    *ConfigMap
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConfigMap(tt.args.namespace, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfigMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfigMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConfigMaps(t *testing.T) {
	type args struct {
		namespace string
	}
	tests := []struct {
		name    string
		args    args
		want    *ConfigMapList
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConfigMaps(tt.args.namespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfigMaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfigMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
