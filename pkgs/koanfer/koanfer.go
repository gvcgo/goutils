package koanfer

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	utils "github.com/moqsien/goutils/pkgs/gutils"
)

type KoanfJSON struct{}

func NewJsonParser() *KoanfJSON {
	return &KoanfJSON{}
}

// Unmarshal parses the given JSON bytes.
func (p *KoanfJSON) Unmarshal(b []byte) (map[string]interface{}, error) {
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Marshal marshals the given config map to JSON bytes.
func (p *KoanfJSON) Marshal(o map[string]interface{}) ([]byte, error) {
	return json.MarshalIndent(o, "", "    ")
}

type JsonKoanfer struct {
	k      *koanf.Koanf
	parser *KoanfJSON
	fpath  string // file path
}

func NewKoanfer(path string) (r *JsonKoanfer, err error) {
	r = &JsonKoanfer{
		k:      koanf.New("."),
		parser: &KoanfJSON{},
		fpath:  path,
	}
	err = r.initDirs()
	return
}

func (that *JsonKoanfer) initDirs() (err error) {
	pDir := filepath.Dir(that.fpath)
	if ok, _ := utils.PathIsExist(pDir); !ok {
		if err = os.MkdirAll(pDir, os.ModePerm); err != nil {
			return
		}
	}
	return
}

func (that *JsonKoanfer) Save(obj interface{}) (err error) {
	that.k.Load(structs.Provider(obj, "koanf"), nil)
	if b, err := that.k.Marshal(that.parser); err == nil && len(b) > 0 {
		return os.WriteFile(that.fpath, b, 0666)
	} else {
		return err
	}
}

func (that *JsonKoanfer) Load(obj interface{}) (err error) {
	err = that.k.Load(file.Provider(that.fpath), that.parser)
	if err != nil {
		return
	}
	err = that.k.UnmarshalWithConf("", obj, koanf.UnmarshalConf{Tag: "koanf"})
	return
}
