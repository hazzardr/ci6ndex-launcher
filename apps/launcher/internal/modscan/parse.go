package modscan

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/hazzardr/ci6ndex-launcher/shared"
)

func ParseModInfo(path string) (shared.Mod, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return shared.Mod{}, fmt.Errorf("read modinfo file: %w", err)
	}

	var mi ModInfo
	if err := xml.Unmarshal(data, &mi); err != nil {
		return shared.Mod{}, fmt.Errorf("parse modinfo xml: %w", err)
	}

	id, err := uuid.Parse(mi.ID)
	if err != nil {
		return shared.Mod{}, fmt.Errorf("parse mod uuid %q: %w", mi.ID, err)
	}

	dir := filepath.Dir(path)
	wsid, _ := strconv.ParseUint(filepath.Base(dir), 10, 0)

	return shared.Mod{
		UUID:       id,
		Name:       mi.Properties.Name,
		Version:    mi.Version,
		Author:     mi.Properties.Authors,
		Source:     "workshop",
		Path:       dir,
		WorkshopID: uint(wsid),
	}, nil
}
