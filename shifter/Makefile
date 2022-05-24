# Copyright 2019 Google LLC
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#    http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


.PHONY: test lint templateTest clean build apply t1

test: yamlTest yamlMultiTest yamlDCTest yamlMultiOutputSingleTest yamlQNSTest templateTest
#test: serverTest

clean:

build: shifter_linux_amd64 shifter_darwin_amd64 shifter_win_amd64.exe

shifter_linux_amd64: fmt
	env GOOS=linux GOARCH=amd64 go build --ldflags '-linkmode external -extldflags "-static"' -o $@ -v shifter

shifter_darwin_amd64: fmt
	env GOOS=darwin GOARCH=amd64 go build -o $@ -v shifter

shifter_win_amd64.exe: fmt
	env GOOS=windows GOARCH=amd64 go build -o $@ -v shifter

fmt:
	go fmt ./...

apply:
	helm install  ./out -n test --generate-name

server:
	go run . server

# Tests
# ---------------

helmlint:
	helm lint ./out

yamlTest: fmt
	go run . convert -t yaml -f ./_test/yaml/multidoc/os-nginx.yaml -o ./out/t1_yaml -i yaml

yamlMultiTest: fmt
	go run . convert -t yaml -f ./_test/yaml/multifile/ -o ./out/t2_yaml_multifile/ -i yaml

yamlMultiOutputSingleTest: fmt
	go run . convert -t yaml -f ./_test/yaml/multifile/ -o ./out/t3_yaml_multiotput/files.yaml -i yaml

yamlDCTest: fmt
	go run . convert -t yaml -i yaml -f ./_test/yaml/deploymentconfig.yaml -o ./out/t4_deploymentconfig

yamlQNSTest: fmt
	go run . convert -t yaml -i yaml -f ./_test/yaml/quoted_nested_strings.yaml -o ./out/t5_quoted_nested_strings/quoted_nested_strings.yaml

templateTest:
	go run . convert -t helm -i template -f ./_test/os-nginx-template.yaml -o ./out/t6_helm

#serverTest:
#	go run . server -p 8081 -f ./data/source -o ./data/output
