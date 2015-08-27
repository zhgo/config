// Copyright 2014 The zhgo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
    "log"
    "os"
)

func WorkingDir() string {
    w, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    return w
}
