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
	env GOOS=linux GOARCH=amd64 go build --ldflags '-linkmode external -extldflags "-static" -X shifter/cmd.Version=v0.3.0 -X shifter/cmd.Platform=amd64_GNU/Linux' -o $@ -v shifter

shifter_darwin_amd64: fmt
	env GOOS=darwin GOARCH=amd64 go build -o $@ -v shifter

shifter_win_amd64.exe: fmt
	env GOOS=windows GOARCH=amd64 go build --ldflags '-X shifter/cmd.Version=v0.3.0 -X shifter/cmd.Platform=amd64_Windows' -o $@ -v shifter

fmt:
	go fmt ./...

secScan:
	gosec -no-fail -fmt=json -out=results.json -stdout -verbose=text ./...

apply:
	helm install  ./out -n test --generate-name

server:
	go run . server

# Tests
# ---------------

helmlint:
	helm lint ./out

yamlTest: fmt
	go run . convert -o yaml -i yaml ./_test/yaml/multidoc/os-nginx.yaml ./out/t1_yaml

yamlMultiTest: fmt
	go run . convert -o yaml -i yaml ./_test/yaml/multifile/ ./out/t2_yaml_multifile/

yamlMultiOutputSingleTest: fmt
	go run . convert -o yaml -i yaml ./_test/yaml/multifile/ ./out/t3_yaml_multiotput/files.yaml

yamlDCTest: fmt
	go run . convert -o yaml -i yaml ./_test/yaml/deploymentconfig.yaml ./out/t4_deploymentconfig

yamlQNSTest: fmt
	go run . convert -o yaml -i yaml ./_test/yaml/quoted_nested_strings.yaml ./out/t5_quoted_nested_strings/quoted_nested_strings.yaml

templateTest:
	go run . convert -o helm -i template ./_test/os-nginx-template.yaml ./out/t6_helm

#serverTest:
#	go run . server -p 8081 -f ./data/source -o ./data/output
