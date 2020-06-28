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
metamodel_version:=v0.0.30
metamodel_url:=https://github.com/openshift-online/ocm-api-metamodel.git

.PHONY: check
check: metamodel
	metamodel/metamodel check --model=model

.PHONY: openapi
openapi: metamodel
	metamodel/metamodel generate openapi --model=model --output=openapi

.PHONY: metamodel
metamodel:
	rm -rf "$@"
	if [ -d "$(metamodel_url)" ]; then \
		cp -r "$(metamodel_url)" "$@"; \
	else \
		git clone "$(metamodel_url)" "$@"; \
		cd "$@"; \
		git fetch --tags origin; \
		git checkout -B build "$(metamodel_version)"; \
	fi
	make -C "$@"

# Enforce indentation by tabs. License contains 2 spaces, so reject 3+.
lint:
	find -name '*.model' -print0 | xargs -0 sh -c '! egrep --with-filename --line-number "^(   |\t+ )" "$$@"'

.PHONY: clean
clean:
	rm -rf \
		metamodel \
		openapi \
		$(NULL)
