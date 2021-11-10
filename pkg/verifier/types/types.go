/*
Copyright The Ratify Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"encoding/json"
	"io"

	"github.com/deislabs/ratify/pkg/verifier"
)

const (
	SpecVersion      string = "0.1.0"
	Version          string = "version"
	Name             string = "name"
	ArtifactTypes    string = "artifactTypes"
	NestedReferences string = "nestedReferences"
)

const (
	ErrUnknown                     uint = iota // 0
	ErrConfigParsingFailure                    // 1
	ErrInvalidVerifierConfig                   // 2
	ErrUnknownCommand                          // 3
	ErrMissingEnvironmentVariables             // 4
	ErrIOFailure                               // 5
	ErrVersionNotSupported                     // 6
	ErrArgsParsingFailure                      // 7
	ErrPluginCmdFailure                        // 8
	ErrInternalFailure             uint = 999
)

type VerifierResult struct {
	IsSuccess bool     `json:"isSuccess"`
	Results   []string `json:"results"`
	Name      string   `json:"name"`
}

func GetVerifierResult(result []byte) (*verifier.VerifierResult, error) {
	vResult := VerifierResult{}
	if err := json.Unmarshal(result, &vResult); err != nil {
		return nil, err
	}
	return &verifier.VerifierResult{
		IsSuccess: vResult.IsSuccess,
		Results:   vResult.Results,
		Name:      vResult.Name,
	}, nil
}

func WriteVerifyResultResult(result *verifier.VerifierResult, w io.Writer) error {
	return json.NewEncoder(w).Encode(result)
}
