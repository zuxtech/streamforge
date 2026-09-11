/**
 * Copyright (c) 2026 ZuxTech.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * This source code is licensed under the Apache License, Version 2.0
 * found in the LICENSE file in the root directory of this source tree.
 */

package http

import (
	"encoding/json"
	"log"
	"net/http"
)



func writeJSON(w http.ResponseWriter, status int, value any) {
   w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(status)

    if err := json.NewEncoder(w).Encode(value); err != nil {
        log.Printf("http response write error: %v", err)
    }
}