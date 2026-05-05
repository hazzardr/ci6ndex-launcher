package modscan

import "encoding/xml"

type ModInfo struct {
	XMLName     xml.Name              `xml:"Mod"`
	ID          string                `xml:"id,attr"`
	Version     string                `xml:"version,attr"`
	Properties  ModInfoProperties     `xml:"Properties"`
}

type ModInfoProperties struct {
	Name               string `xml:"Name"`
	Description        string `xml:"Description"`
	Created            string `xml:"Created"`
	Teaser             string `xml:"Teaser"`
	Authors            string `xml:"Authors"`
	SpecialThanks      string `xml:"SpecialThanks"`
	AffectsSavedGames  string `xml:"AffectsSavedGames"`
	CompatibleVersions string `xml:"CompatibleVersions"`
}