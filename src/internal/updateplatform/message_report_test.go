// SPDX-FileCopyrightText: 2026 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package updateplatform

import "testing"

func TestHasDeliveryRepo(t *testing.T) {
	manager := &UpdatePlatformManager{
		repoInfos: []repoInfo{
			{Source: "deb https://packages.example.com/desktop beige main"},
			{Source: "deb delivery://packages.example.com/apps beige main"},
		},
	}

	if !manager.HasDeliveryRepo() {
		t.Fatal("expected delivery repo to be detected")
	}
}
