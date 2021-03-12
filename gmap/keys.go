// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Since: v0.0.7

package gmap

func Keys(m map[string]string) []string {
    keys := make([]string, 0, len(m))

    for k, _ := range m {
        keys = append(keys, k)
    }

    return keys
}
