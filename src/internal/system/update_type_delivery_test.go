// SPDX-FileCopyrightText: 2026 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package system

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceMatchedReposWithDelivery(t *testing.T) {
	localContent := strings.Join([]string{
		"deb https://packages.example.com/desktop beige main",
		"deb https://packages.example.com/custom beige main",
		"# keep comments untouched",
		"deb https://security.example.com beige-security main",
	}, "\n")

	platformRepos := []string{
		"deb https://packages.example.com/desktop beige main",
	}

	got := replaceMatchedReposWithDelivery(localContent, platformRepos)
	lines := strings.Split(got, "\n")

	assert.Equal(t, "deb delivery://packages.example.com/desktop beige main", lines[0])
	assert.Equal(t, "deb https://packages.example.com/custom beige main", lines[1])
	assert.Equal(t, "# keep comments untouched", lines[2])
	assert.Equal(t, "deb https://security.example.com beige-security main", lines[3])
}

func TestReplaceMatchedReposWithDeliveryRequiresExactMatch(t *testing.T) {
	localContent := "deb https://packages.example.com/desktop beige main"
	platformRepos := []string{
		"deb http://packages.example.com/desktop beige main",
	}

	got := replaceMatchedReposWithDelivery(localContent, platformRepos)

	assert.Equal(t, localContent, got)
}

func TestReplaceMatchedReposWithDeliveryWithoutPlatformReposKeepsOriginal(t *testing.T) {
	localContent := "deb https://packages.example.com/desktop beige main"

	got := replaceMatchedReposWithDelivery(localContent, nil)

	assert.Equal(t, localContent, got)
}
