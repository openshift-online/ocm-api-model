#
# Copyright (c) 2019 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Details of the metamodel used to check the model:
metamodel_version:=v0.0.9
metamodel_url:=https://github.com/openshift-online/ocm-api-metamodel.git

.PHONY: check
check: metamodel
	metamodel/ocm-metamodel-tool check --model=model

.PHONY: metamodel
metamodel:
	[ -d "$@" ] || git clone "$(metamodel_url)" "$@"
	cd "$@" && git fetch --tags origin
	cd "$@" && git checkout -B build "$(metamodel_version)"
	make -C "$@"

.PHONY: clean
clean:
	rm -rf metamodel
