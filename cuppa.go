//
// Copyright © 2016 Bryan T. Meyers <bmeyers@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"github.com/DataDrake/cuppa/cmd/search"
	"github.com/DataDrake/cuppa/providers"
	"os"
)

func usage() {
	print("USAGE: cuppa CMD [OPTIONS]\n")
}

func main() {

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	ps := []providers.Provider{providers.CPANProvider{}}

	switch os.Args[1] {
	case "search":
		search.Execute(ps)
	}

	os.Exit(0)
}