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
metamodel_version:=v0.0.54
metamodel_url:=https://github.com/openshift-online/ocm-api-metamodel/releases/download/$(metamodel_version)/metamodel-linux-amd64
metamodel_sum:=76ea3d91a7d601c5e8891455c8c43a918e8dac21c9b1397f636c5b01af0afa7d

.PHONY: check
check: metamodel
	./metamodel check --model=model

.PHONY: openapi
openapi: metamodel
	./metamodel generate openapi --model=model --output=openapi

metamodel:
	wget --progress=dot:giga --output-document="$@" "$(metamodel_url)"
	echo "$(metamodel_sum) $@" | sha256sum --check
	chmod +x "$@"

# Enforce indentation by tabs. License contains 2 spaces, so reject 3+.
lint:
	find -name '*.model' -print0 | xargs -0 sh -c '! egrep --with-filename --line-number "^(   |\t+ )" "$$@"'

.PHONY: clean
clean:
	rm -rf \
		metamodel \
		openapi \
		$(NULL)
