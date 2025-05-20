package funcs

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"embed"

	"github.com/qz-io/tcode-modules/pkg/common/callback"
	"github.com/qz-io/tcode-modules/pkg/model"
)

//go:embed all:so/f
var pluginFSFunc embed.FS

//go:embed all:so/g
var pluginGSFunc embed.FS
var (
	A           func(*model.CodeRequest, string, *callback.ProgressWriter, *sync.Map) *exec.Cmd
	V           func(*model.CodeRequest, *callback.ProgressWriter, string, *sync.Map) (*exec.Cmd, []string)
	M           func(*model.CodeRequest, string, *callback.ProgressWriter, *sync.Map) *exec.Cmd
	AG          func() string
	RG          func(string)
	CG          func() (string, int, bool)
	Atn         func(*strings.Reader) (map[string]float64, bool, error)
	GetBody     func(callback.BaseProgressStatusService, string, bool) string
	BlackDetect func(url string, start int) (float64, float64, float64)
	IsBlack     func(imgPath string) (bool, error)
	ExImage     func(time int, basePath, url string) (string, int)
	R0          func()
)

func fc() []byte {
	bs, err := pluginFSFunc.ReadFile("so/f")
	if err != nil {
		fmt.Printf("f %v\n", err)
		return nil
	}
	return bs
}

func gc() []byte {
	bs, err := pluginFSFunc.ReadFile("so/g")
	if err != nil {
		fmt.Printf("g %v\n", err)
		return nil
	}
	return bs
}

func Load() error {
	p, l := callback.LoadPlugin(fc())

	symbol2, err2 := p.Lookup("A")
	if err2 != nil {
		fmt.Printf("Lookup audioFunc function failed: %v\n", err2)
		panic("Lookup audioFunc function failed")
	}
	A = *symbol2.(*func(*model.CodeRequest, string, *callback.ProgressWriter, *sync.Map) *exec.Cmd)
	fmt.Println("A loaded")

	symbol3, err2 := p.Lookup("V")
	if err2 != nil {
		panic("Lookup V function failed")
	}
	V = *symbol3.(*func(*model.CodeRequest, *callback.ProgressWriter, string, *sync.Map) (*exec.Cmd, []string))
	fmt.Println("v loaded")

	symbol4, err2 := p.Lookup("M")
	if err2 != nil {
		panic("Lookup M function failed")
	}
	M = *symbol4.(*func(*model.CodeRequest, string, *callback.ProgressWriter, *sync.Map) *exec.Cmd)
	fmt.Println("M loaded")

	//
	symbol5, err2 := p.Lookup("AG")
	if err2 != nil {
		return err2
	}
	AG = *symbol5.(*func() string)

	//
	symbol6, err2 := p.Lookup("RG")
	if err2 != nil {
		return err2
	}
	RG = *symbol6.(*func(string))

	//
	symbol7, err2 := p.Lookup("CG")
	if err2 != nil {
		return err2
	}
	CG = *symbol7.(*func() (string, int, bool))

	symbol8, err2 := p.Lookup("BD")
	if err2 != nil {
		return err2
	}
	BlackDetect = *symbol8.(*func(string, int) (float64, float64, float64))

	symbol9, err2 := p.Lookup("IB")
	if err2 != nil {
		return err2
	}
	IsBlack = *symbol9.(*func(string) (bool, error))

	symbol10, err2 := p.Lookup("EI")
	if err2 != nil {
		return err2
	}
	ExImage = *symbol10.(*func(int, string, string) (string, int))

	symbol11, err2 := p.Lookup("ATN")
	if err2 != nil {
		return err2
	}
	Atn = *symbol11.(*func(*strings.Reader) (map[string]float64, bool, error))

	symbol12, err2 := p.Lookup("GMB")
	if err2 != nil {
		return err2
	}
	GetBody = *symbol12.(*func(callback.BaseProgressStatusService, string, bool) string)

	go func() {
		defer func() {
			recover()
		}()
		l(gc())
	}()

	return nil

}
