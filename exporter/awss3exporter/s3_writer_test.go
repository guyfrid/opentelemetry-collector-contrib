// Copyright The OpenTelemetry Authors
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

package awss3exporter

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestS3TimeKey(t *testing.T) {
	const layout = "2006-01-02"

	tm, err := time.Parse(layout, "2022-06-05")
	timeKey := getTimeKey(tm, "hour")

	assert.NoError(t, err)
	require.NotNil(t, tm)
	assert.Equal(t, "year=2022/month=06/day=05/hour=00", timeKey)

	timeKey = getTimeKey(tm, "minute")
	assert.Equal(t, "year=2022/month=06/day=05/hour=00/minute=00", timeKey)
}

func TestS3Key(t *testing.T) {
	const layout = "2006-01-02"

	tm, err := time.Parse(layout, "2022-06-05")

	assert.NoError(t, err)
	require.NotNil(t, tm)

	re := regexp.MustCompile(`keyprefix/year=2022/month=06/day=05/hour=00/minute=00/fileprefixlogs_([0-9]+).json`)
	s3Key := getS3Key(tm, "keyprefix", "minute", "fileprefix", "logs", "json")
	matched := re.MatchString(s3Key)
	assert.Equal(t, true, matched)
}

func TestGetAwsEndpoint(t *testing.T) {
	const endpointURL = "https://endpoint.com"
	config := &Config{
		S3Uploader: S3UploaderConfig{
			AwsEndpoint: endpointURL,
		},
	}
	ep := getAwsEndpoint(config)
	assert.Equal(t, ep, endpointURL)
}

func TestGetAwsEndpointNoConfigure(t *testing.T) {
	const endpointURL = "https://default_endpoint.com"
	t.Setenv("AWS_ENDPOINT_URL", endpointURL)
	config := &Config{}
	ep := getAwsEndpoint(config)
	assert.Equal(t, ep, endpointURL)
}
