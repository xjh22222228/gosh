// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// Private

package ginternal

// "1" => "01"
// "10" => "10"
func AddZero(str string) string {
    if len(str) == 1 {
        return "0" + str
    }
    return str
}
