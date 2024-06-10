/*
 * Copyright 2020 The Yorkie Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"os"
)

/*
Document store for collaborative applications based on CRDT

Usage:
  yorkie [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  context     Manage contexts
  document    Manage documents
  help        Help about any command
  history     Show the history of a document
  login       Log in to Yorkie server
  logout      Log out from the Yorkie server
  project     Manage projects
  server      Start Yorkie server
  version     Print the version number of Yorkie

Flags:
  -h, --help              help for yorkie
      --rpc-addr string   Address of the rpc server (default "localhost:8080")

Use "yorkie [command] --help" for more information about a command.
*/

func main() {
	os.Exit(Run())
}
