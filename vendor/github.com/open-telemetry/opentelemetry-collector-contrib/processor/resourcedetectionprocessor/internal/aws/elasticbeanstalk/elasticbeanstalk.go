// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticbeanstalk // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/elasticbeanstalk"

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/elasticbeanstalk/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "elastic_beanstalk"

	linuxPath   = "/var/elasticbeanstalk/xray/environment.conf"
	windowsPath = "C:\\Program Files\\Amazon\\XRay\\environment.conf"
)

var _ internal.Detector = (*Detector)(nil)

type Detector struct {
<<<<<<< HEAD
	fs                 fileSystem
	resourceAttributes metadata.ResourceAttributesConfig
=======
	fs fileSystem
	rb *metadata.ResourceBuilder
>>>>>>> main
}

type EbMetaData struct {
	DeploymentID    int    `json:"deployment_id"`
	EnvironmentName string `json:"environment_name"`
	VersionLabel    string `json:"version_label"`
}

func NewDetector(_ processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
<<<<<<< HEAD
	return &Detector{fs: &ebFileSystem{}, resourceAttributes: cfg.ResourceAttributes}, nil
}

func (d Detector) Detect(context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	res := pcommon.NewResource()
=======
	return &Detector{fs: &ebFileSystem{}, rb: metadata.NewResourceBuilder(cfg.ResourceAttributes)}, nil
}

func (d Detector) Detect(context.Context) (resource pcommon.Resource, schemaURL string, err error) {
>>>>>>> main
	var conf io.ReadCloser

	if d.fs.IsWindows() {
		conf, err = d.fs.Open(windowsPath)
	} else {
		conf, err = d.fs.Open(linuxPath)
	}

	// Do not want to return error so it fails silently on non-EB instances
	if err != nil {
		// TODO: Log a more specific message with zap
<<<<<<< HEAD
		return res, "", nil
=======
		return pcommon.NewResource(), "", nil
>>>>>>> main
	}

	ebmd := &EbMetaData{}
	err = json.NewDecoder(conf).Decode(ebmd)
	conf.Close()

	if err != nil {
		// TODO: Log a more specific error with zap
<<<<<<< HEAD
		return res, "", err
	}

	attr := res.Attributes()
	if d.resourceAttributes.CloudProvider.Enabled {
		attr.PutStr(conventions.AttributeCloudProvider, conventions.AttributeCloudProviderAWS)
	}
	if d.resourceAttributes.CloudPlatform.Enabled {
		attr.PutStr(conventions.AttributeCloudPlatform, conventions.AttributeCloudPlatformAWSElasticBeanstalk)
	}
	if d.resourceAttributes.ServiceInstanceID.Enabled {
		attr.PutStr(conventions.AttributeServiceInstanceID, strconv.Itoa(ebmd.DeploymentID))
	}
	if d.resourceAttributes.DeploymentEnvironment.Enabled {
		attr.PutStr(conventions.AttributeDeploymentEnvironment, ebmd.EnvironmentName)
	}
	if d.resourceAttributes.ServiceVersion.Enabled {
		attr.PutStr(conventions.AttributeServiceVersion, ebmd.VersionLabel)
	}
	return res, conventions.SchemaURL, nil
=======
		return pcommon.NewResource(), "", err
	}

	d.rb.SetCloudProvider(conventions.AttributeCloudProviderAWS)
	d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformAWSElasticBeanstalk)
	d.rb.SetServiceInstanceID(strconv.Itoa(ebmd.DeploymentID))
	d.rb.SetDeploymentEnvironment(ebmd.EnvironmentName)
	d.rb.SetServiceVersion(ebmd.VersionLabel)

	return d.rb.Emit(), conventions.SchemaURL, nil
>>>>>>> main
}
