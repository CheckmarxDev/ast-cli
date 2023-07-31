package wrappers

import (
	"net/url"
	"time"
)

const (
	Analyzer_Name = "CxOne"
	Analyzer_Id   = Analyzer_Name + "-SAST"
	Analyzer_url  = "https://checkmarx.company.com/"
	Vendor_name   = "Checkmarx"
)

var sast_verion string

type GlSastResultsCollection struct {
	Scan            ScanGlReport        `json:"scan"`
	Schema          string              `json:"schema"`
	Version         string              `json:"version"`
	Vulnerabilities []GlVulnerabilities `json:"vulnerabilities"`
}
type GlVulnerabilities struct {
	ID          string       `json:"id"`
	Category    string       `json:"category"`
	Name        string       `json:"name"`
	Message     string       `json:"message"`
	Description string       `json:"description"`
	CVE         string       `json:"cve"`
	Severity    string       `json:"severity"`
	Confidence  string       `json:"confidence"`
	Solution    string       `json:"solution"`
	Scanner     GlScanner    `json:"scanner"`
	Identifiers []Identifier `json:"identifiers"`
	Links       []string     `json:"links"`
	Tracking    Tracking     `json:"tracking"`
	Flags       Flag         `json:"flags"`
	Location    Location     `json:"location"`
}
type Identifier struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Value string `json:"value"`
}
type Flag struct {
	Type        string `json:"type"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
}
type Location struct {
	File      string `json:"file"`
	StartLine uint   `json:"start_line"`
	EndLine   uint   `json:"end_line"`
	Class     string `json:"class"`
}

type Tracking struct {
	Items Item `json:"items"`
}
type Item struct {
	Signatures Signature `json:"signatures"`
	File       string    `json:"file"`
	EndLine    uint      `json:"end_line"`
	StartLine  uint      `json:"start_line"`
}
type Signature struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}
type ScanGlReport struct {
	EndTime   time.Time `json:"end_time"`
	Analyzer  Analyzer  `json:"analyzer"`
	Scanner   GlScanner `json:"scanner"`
	StartTime time.Time `json:"start_time"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
}

type Analyzer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Vendor  Vendor `json:"vendor"`
	Version string `json:"version"`
}
type GlScanner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Vendor struct {
	Name string `json:"name"`
}
type GLSastIdentifiers struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Url   url.URL `json:"url"`
	Value string  `json:"value"`
}
type GlSastTracking struct {
	items []GlSastTrackingItems `json:"items"`
}

type GlSastTrackingItems struct {
	Signatures GlSastTrackingItemsSignatures `json:"signatures"`
	File       string                        `json:"file"`
	End_line   string                        `json:"end_line"`
	Start_line string                        `json:"start_line"`
}
type GlSastTrackingItemsSignatures struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}
type GlSastFlags struct {
	Type        string `json:"type"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
}
