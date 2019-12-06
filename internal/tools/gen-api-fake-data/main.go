// Copyright 2016-2019 The Libsacloud Authors
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

package main

import (
	"context"
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud/v2/internal/tools"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/search/keys"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/archive"
)

const destination = "sacloud/fake/zz_init_archive.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-fake-data: ")
}

var fakeDataDefines = []struct {
	destination   string
	template      string
	parameterFunc func() interface{}
}{
	{
		destination:   "sacloud/fake/zz_init_archive.go",
		template:      tmplArchive,
		parameterFunc: collectArchives,
	},
	{
		destination:   "sacloud/fake/zz_init_cdrom.go",
		template:      tmplCDROM,
		parameterFunc: collectCDROMs,
	},
}

func main() {
	for _, generator := range fakeDataDefines {
		param := generator.parameterFunc()
		tools.WriteFileWithTemplate(&tools.TemplateConfig{
			OutputPath: filepath.Join(tools.ProjectRootPath(), generator.destination),
			Template:   generator.template,
			Parameter:  param,
		})
		log.Printf("generated: %s\n", filepath.Join(generator.destination))
	}

}

func collectArchives() interface{} {
	caller, err := sacloud.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	tmplParam := map[string][]*sacloud.Archive{}
	zones := []string{"is1a", "is1b", "tk1a", "tk1v"}
	archiveOp := sacloud.NewArchiveOp(caller)
	ctx := context.Background()

	for _, zone := range zones {
		var archives []*sacloud.Archive
		for _, ost := range ostype.ArchiveOSTypes {
			archive, err := archive.FindByOSType(ctx, archiveOp, zone, ost)
			if err != nil {
				log.Fatal(err)
			}
			archives = append(archives, archive)
		}
		tmplParam[zone] = archives
	}
	return tmplParam
}

const tmplArchive = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-fake-data'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var initArchives = map[string][]*sacloud.Archive{
{{ range $zone, $archives := . -}}
	"{{$zone}}": {
{{ range $archives -}}
		{
			ID:                   types.ID({{.ID}}),
			Name:                 "{{.Name}}",
			Description:          "fake",
			Tags:                 types.Tags{ {{range .Tags}}"{{.}}",{{ end }} },
			DisplayOrder:         {{.DisplayOrder}},
			Availability:         types.EAvailability("{{.Availability}}"),
			Scope:                types.EScope("{{.Scope}}"),
			SizeMB:               {{.SizeMB}},
			DiskPlanID:           types.ID({{.DiskPlanID}}),
			DiskPlanName:         "{{.DiskPlanName}}",
			DiskPlanStorageClass: "{{.DiskPlanStorageClass}}",
		},
{{ end -}}
	},
{{ end -}}
}
`

func collectCDROMs() interface{} {
	caller, err := sacloud.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	tmplParam := map[string][]*sacloud.CDROM{}
	zones := []string{"is1a", "is1b", "tk1a", "tk1v"}
	cdromOp := sacloud.NewCDROMOp(caller)
	ctx := context.Background()

	for _, zone := range zones {
		var cdroms []*sacloud.CDROM

		searched, err := cdromOp.Find(ctx, zone, &sacloud.FindCondition{
			Filter: search.Filter{
				search.Key(keys.Scope): string(types.Scopes.Shared),
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		cdroms = append(cdroms, searched.CDROMs...)
		tmplParam[zone] = cdroms
	}
	return tmplParam
}

const tmplCDROM = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-fake-data'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var initCDROM = map[string][]*sacloud.CDROM{
{{ range $zone, $data := . -}}
	"{{$zone}}": {
{{ range $data -}}
		{
			ID:                   types.ID({{.ID}}),
			Name:                 "{{.Name}}",
			Description:          "fake",
			Tags:                 types.Tags{ {{range .Tags}}"{{.}}",{{ end }} },
			DisplayOrder:         {{.DisplayOrder}},
			Availability:         types.EAvailability("{{.Availability}}"),
			Scope:                types.EScope("{{.Scope}}"),
			SizeMB:               {{.SizeMB}},
		},
{{ end -}}
	},
{{ end -}}
}
`