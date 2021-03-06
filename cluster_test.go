/*
Copyright 2014 Rohith All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package marathon

import (
	"testing"
	"time"
)

var cluster Cluster

func GetFakeCluster() {
	if cluster == nil {
		cluster, _ = NewMarathonCluster(FAKE_MARATHON_URL)
	}
}

func TestUrl(t *testing.T) {
	GetFakeCluster()
	assertOnString(cluster.Url(), FAKE_MARATHON_URL, t)
}

func TestSize(t *testing.T) {
	GetFakeCluster()
	assertOnInteger(cluster.Size(), 2, t)
}

func TestActive(t *testing.T) {
	GetFakeCluster()
	assertOnInteger(len(cluster.Active()), 2, t)
}

func TestNonActive(t *testing.T) {
	GetFakeCluster()
	assertOnInteger(len(cluster.NonActive()), 0, t)
}

func TestGetMember(t *testing.T) {
	GetFakeCluster()
	member, err := cluster.GetMember()
	assertOnError(err, t)
	assertOnString(member, "http://127.0.0.1:3000", t)
}

func TestMarkdown(t *testing.T) {
	GetFakeCluster()
	assertOnInteger(len(cluster.Active()), 2, t)
	cluster.MarkDown()
	time.Sleep(10 * time.Millisecond)
	assertOnInteger(len(cluster.Active()), 2, t)
}
