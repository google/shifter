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



test: yamlTest yamlMultiTest

lint:
	helm lint ./out

clean:

build: fmt
	env GOOS=linux go build -v shifter

fmt:
	go fmt ./...

apply:
	helm install  ./out -n test --generate-name

# Tests
# ---------------

yamlTest: fmt
	go run . convert -t yaml -i ./_test/yaml/multidoc/os-nginx.yaml -o ./out -k yaml

yamlMultiTest: fmt
	go run . convert -t yaml -i ./_test/yaml/multifile/ -o ./out/files -k yaml

templateTest:
	go run . convert -t template -i ./_test/os-nginx-template.yaml -o ./out -k helm
