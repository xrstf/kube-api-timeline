// SPDX-FileCopyrightText: 2023 Christoph Mewes
// SPDX-License-Identifier: MIT

package timeline

import "time"

type Timeline struct {
	APIGroups []APIGroup
	Releases  []ReleaseMetadata
}

type ReleaseMetadata struct {
	Version       string
	Released      bool
	Supported     bool
	Archived      bool
	HasDocs       bool
	ReleaseDate   time.Time
	EndOfLifeDate *time.Time
	LatestVersion string
}

func (o *Timeline) ReleaseMetadata(release string) ReleaseMetadata {
	for _, r := range o.Releases {
		if r.Version == release {
			return r
		}
	}

	return ReleaseMetadata{}
}

func (o *Timeline) HasRelease(release string) bool {
	return o.ReleaseMetadata(release).Version != ""
}

type APIGroup struct {
	Name               string
	Archived           bool
	PreferredVersions  map[string]string // lists the prefered version per release
	ReleasesOfInterest []string          // releases which have notable changes for this API group
	APIVersions        []APIVersion
}

// helper functions for templating :grin:

func (o *APIGroup) PreferredVersion(release string) string {
	return o.PreferredVersions[release]
}

type APIVersion struct {
	Version            string // e.g. "v1beta1"
	Archived           bool
	Releases           []string // releases which have this API version
	ReleasesOfInterest []string // releases which have notable changes for this API version
	Resources          []APIResource
}

func (o *APIVersion) HasRelease(release string) bool {
	for _, r := range o.Releases {
		if r == release {
			return true
		}
	}

	return false
}

type APIResource struct {
	Kind               string
	Singular           string
	Plural             string
	Archived           bool
	Scopes             map[string]string
	Releases           []string // releases which have this resource
	ReleasesOfInterest []string // releases which have notable changes for this resource
	Description        string
	DocRelease         string // release which should be linked to for this resource (can be empty!)
}

func (o *APIResource) HasRelease(release string) bool {
	for _, r := range o.Releases {
		if r == release {
			return true
		}
	}

	return false
}
