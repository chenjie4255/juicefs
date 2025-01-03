/*
 * JuiceFS, Copyright 2024 Juicedata, Inc.
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

package meta

import (
	"fmt"
	"sync/atomic"

	"google.golang.org/protobuf/proto"
)

func (m *kvMeta) dump(ctx Context, opt *DumpOption, ch chan<- *dumpedResult) error {
	return nil
}

func (m *kvMeta) load(ctx Context, typ int, opt *LoadOption, val proto.Message) error {
	return nil
}

func (m *kvMeta) prepareLoad(ctx Context, opt *LoadOption) error {
	return nil
}

func printSums(sums map[int]*atomic.Uint64) string {
	var p string
	for typ, sum := range sums {
		p += fmt.Sprintf("%s num: %d\n", getMessageNameFromType(typ), sum.Load())
	}
	return p
}
