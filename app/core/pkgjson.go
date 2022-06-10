/*

Redefine - UI component documentation

MIT License.
Copyright (c) 2022, Sandeep Gupta.

Use of this source code is governed by a MIT style license
that can be found in LICENSE file in the code repository.

*/

package core

type PackageAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Url   string `json:"url"`
}

type PackageJson struct {
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Description string        `json:"description"`
	HomePage    string        `json:"homePage"`
	Author      PackageAuthor `json:"author"`
	License     string        `json:"license"`
	MainFile    string        `json:"main"`
}
