// Copyright (c) 2012 - Cloud Instruments Co., Ltd.
// 
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met: 
// 
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer. 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution. 
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package seelog

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func countSequencedRowsInFile(filePath string) (int64, error) {
	bts, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	bufReader := bufio.NewReader(bytes.NewBuffer(bts))

	var gotCounter int64
	for {
		line, _, bufErr := bufReader.ReadLine()
		if bufErr != nil && bufErr != io.EOF {
			return 0, bufErr
		}

		lineString := string(line)
		if lineString == "" {
			break
		}

		intVal, atoiErr := strconv.ParseInt(lineString, 10, 64)
		if atoiErr != nil {
			return 0, atoiErr
		}

		if intVal != gotCounter {
			return 0, errors.New(fmt.Sprintf("Wrong order: %d Expected: %d\n", intVal, gotCounter))
		}

		gotCounter++
	}

	return gotCounter, nil
}

func tryRemoveFile(filePath string) (err error) {
	err = os.Remove(filePath)
	if os.IsNotExist(err) {
		err = nil
		return
	}
	return
}
