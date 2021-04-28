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


.PHONY: test lint templateTest clean build apply

test: yamlTest yamlMultiTest yamlDCTest yamlMultiOutputSingleTest yamlQNSTest

clean:

build: shifter_linux_amd64 shifter_darwin_amd64 

shifter_linux_amd64: fmt
	env GOOS=linux GOARCH=amd64 go build --ldflags '-linkmode external -extldflags "-static"' -o $@ -v shifter

shifter_darwin_amd64: fmt
	env GOOS=darwin GOARCH=amd64 go build --ldflags '-linkmode external -extldflags "-static"' -o $@ -v shifter

fmt:
	go fmt ./...

apply:
	helm install  ./out -n test --generate-name

# Tests
# ---------------

helmlint:
	helm lint ./out

yamlTest: fmt
	go run . convert -t yaml -f ./_test/yaml/multidoc/os-nginx.yaml -o ./out -i yaml

yamlMultiTest: fmt
	go run . convert -t yaml -f ./_test/yaml/multifile/ -o ./out/files -i yaml

yamlMultiOutputSingleTest: fmt
	go run . convert -t yaml -f ./_test/yaml/multifile/ -o ./out/files.yaml -i yaml

yamlDCTest: fmt
	go run . convert -t yaml -i yaml -f ./_test/yaml/deploymentconfig.yaml -o ./out/dc

yamlQNSTest: fmt
	go run . convert -t yaml -i yaml -f ./_test/yaml/quoted_nested_strings.yaml -o ./out/quoted_nested_strings.yaml

templateTest:
	go run . convert -t template -f ./_test/os-nginx-template.yaml -o ./out/helm -k helm
