// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/sacloud/libsacloud/v2/sacloud/trace"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

const (
	// CharSetAlphaNum アフファベット(小文字)+数値
	CharSetAlphaNum = "abcdefghijklmnopqrstuvwxyz012346789"

	// CharSetAlpha アフファベット(小文字)
	CharSetAlpha = "abcdefghijklmnopqrstuvwxyz"
)

// TestResourcePrefix テスト時に作成するリソースの名称に付与するプレフィックス
//
// このプレフィックスを持つリソースは受入テスト実行後に削除される
const TestResourcePrefix = "libsacloud-test-"

// ResourceName テスト時に作成するリソースの名称
func ResourceName(name string) string {
	return fmt.Sprintf("%s%s", TestResourcePrefix, name)
}

// RandomPrefix テスト時に作成するリソースに付与するランダムなプレフィックスを生成する
func RandomPrefix() string {
	return fmt.Sprintf("%s-%s", RandomName(5,CharSetAlpha),TestResourcePrefix)
}

// WithRandomPrefix ランダムなプレフィックスをつけて返す
func WithRandomPrefix(name string) string {
	return fmt.Sprintf("%s%s", RandomPrefix(), name)
}

// RandomName ランダムな文字列を生成して返す
func RandomName(strlen int, charSet string) string {
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}

var apiCaller *sacloud.Client
var httpTrace bool

var accTestOnce sync.Once
var accTestMu sync.Mutex

// SingletonAPICaller 環境変数からシングルトンAPICallerを作成する
func SingletonAPICaller() sacloud.APICaller {
	accTestMu.Lock()
	defer accTestMu.Unlock()
	accTestOnce.Do(func() {
		if !IsAccTest() {
			sacloud.DefaultStatePollingInterval = 100 * time.Millisecond
			sacloud.DefaultDBStatusPollingInterval = 10 * time.Millisecond
			fake.SwitchFactoryFuncToFake()
			os.Setenv("SAKURACLOUD_ACCESS_TOKEN", "dummy")
			os.Setenv("SAKURACLOUD_ACCESS_TOKEN_SECRET", "dummy")
		}

		if IsEnableTrace() || IsEnableAPITrace() {
			trace.AddClientFactoryHooks()
		}

		if IsEnableTrace() || IsEnableHTTPTrace() {
			httpTrace = true
		}

		//環境変数にトークン/シークレットがある場合のみテスト実施
		accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
		accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

		if accessToken == "" || accessTokenSecret == "" {
			log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
			os.Exit(0) // exit normal
		}
		client := sacloud.NewClient(accessToken, accessTokenSecret)
		client.UserAgent = fmt.Sprintf("test-libsacloud/%s", libsacloud.Version)
		client.AcceptLanguage = "en-US,en;q=0.9"

		client.RetryMax = 20
		client.HTTPClient = &http.Client{
			Transport: &sacloud.RateLimitRoundTripper{RateLimitPerSec: 1},
		}
		if httpTrace {
			client.HTTPClient.Transport = &sacloud.TracingRoundTripper{
				Transport: client.HTTPClient.Transport,
			}
		}

		apiCaller = client
	})
	return apiCaller
}

// TestZone SAKURACLOUD_ZONE環境変数からテスト対象のゾーンを取得 デフォルトはtk1v
func TestZone() string {
	testZone := os.Getenv("SAKURACLOUD_ZONE")
	if testZone == "" {
		testZone = "tk1v"
	}
	return testZone
}

// IsAccTest TESTACC環境変数が指定されているか
func IsAccTest() bool {
	return os.Getenv("TESTACC") != ""
}

// IsEnableTrace SAKURACLOUD_TRACE環境変数が指定されているか
func IsEnableTrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE") != ""
}

// IsEnableAPITrace SAKURACLOUD_TRACE_API環境変数が指定されているか
func IsEnableAPITrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE_API") != ""
}

// IsEnableHTTPTrace SAKURACLOUD_TRACE_HTTP環境変数が指定されているか
func IsEnableHTTPTrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE_HTTP") != ""
}

