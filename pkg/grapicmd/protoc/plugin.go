package protoc

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
)

// Plugin contains args and plugin name for using in protoc command.
type Plugin struct {
	Path string
	Name string
	Args map[string]interface{}
}

// BinName returns a executable binary name.
func (p *Plugin) BinName() string {
	return filepath.Base(p.Path)
}

func (p *Plugin) toProtocArg(outputPath string) string {
	buf := new(bytes.Buffer)
	buf.WriteString("--" + p.Name + "_out")
	args := make([]string, 0, len(p.Args))
	for k, v := range p.Args {
		args = append(args, k+"="+fmt.Sprint(v))
	}
	if len(args) > 0 {
		buf.WriteString("=" + strings.Join(args, ","))
	}
	buf.WriteString(":" + outputPath)
	return buf.String()
}
