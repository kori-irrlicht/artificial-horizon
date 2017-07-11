package main

import (
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"os"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/kori-irrlicht/artificial-horizon/core"
)

const (
	textAsset core.Type = iota
	textureAsset
)

type assetStatus uint

const (
	assetUnknown assetStatus = iota
	assetReady
	assetError
)

type AssetManager struct {
	loader map[core.Type]core.Loader
	status map[string]assetStatus
	asset  map[string]interface{}
}

func (am *AssetManager) Register(loader core.Loader, t core.Type) {
	am.loader[t] = loader
}
func (am *AssetManager) Load(name string, t core.Type) error {
	if am.status[name] != assetUnknown {
		return fmt.Errorf("Asset (%s) already requested.", name)
	}
	loader, ok := am.loader[t]
	if !ok {
		return fmt.Errorf("Unknown loader type: %d", t)
	}
	res, err := loader.Load(name)
	if err != nil {
		am.status[name] = assetError
	} else {
		am.status[name] = assetReady
		am.asset[name] = res
	}

	return err
}

func (am *AssetManager) check(name string) error {
	status, ok := am.status[name]
	if !ok {
		return fmt.Errorf("No asset with name '%s'", name)
	}
	if status == assetError {
		return fmt.Errorf("Asset not correctly loaded")
	}

	return nil
}

func (am *AssetManager) Get(name string, t core.Type) (interface{}, error) {
	if err := am.check(name); err != nil {
		return nil, err
	}
	return am.asset[name], nil
}
func (am *AssetManager) Wait(name string, t core.Type) chan interface{} {
	ch := make(chan interface{}, 1)
	ticker := time.NewTicker(5 * time.Millisecond)
	quit := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := am.check(name); err == nil {
					quit <- true
				} else {
					fmt.Println(err)
				}
			case <-quit:
				ch <- am.asset[name]
				return

			}
		}
	}()

	return ch
}

func newAssetManager() core.AssetManager {
	am := &AssetManager{}
	am.asset = make(map[string]interface{})
	am.status = make(map[string]assetStatus)
	am.loader = make(map[core.Type]core.Loader)
	am.Register(&textLoader{}, textAsset)
	am.Register(&textureLoader{}, textureAsset)

	return am
}

type textLoader struct{}

func (tl textLoader) Load(name string) (interface{}, error) {
	return ioutil.ReadFile(name)
}

type textureLoader struct{}

func (tl textureLoader) Load(name string) (interface{}, error) {
	imgFile, err := os.Open(name)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", name, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}
