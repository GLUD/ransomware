/*
Ransomware example in Golang
Copyright (C) 2017 Gustavo Henrique

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/
package util_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavohenrique/ransomware/util"
)

func TestGenerateRansomwareHtmlPage(t *testing.T) {
	tempDir := os.TempDir()
	htmlFile := filepath.Join(tempDir, "test.html")
	util.GenerateRansomwareHtmlPage(htmlFile, tempDir)

	expected := fmt.Sprintf(util.HTML_TEMPLATE, tempDir)
	generated, _ := ioutil.ReadFile(htmlFile)
	assert.Equal(t, expected, string(generated))
}
