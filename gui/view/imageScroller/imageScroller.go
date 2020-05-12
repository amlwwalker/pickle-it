package imagescroller

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() {
	ImageController_QmlRegisterType2("ImageScroller", 1, 0, "ImageController")
	// fmt.Println("registered")
}

type ImageController struct {
	quick.QQuickItem
	_    func() `constructor:"init"`
	qApp *widgets.QApplication

	listOfPaths []*anImage
	_           *core.QAbstractListModel `property:"imageModel"`

	_ func(imagePath string) `signal:"requestExpandImage"`
}

type anImage struct {
	core.QObject

	_ string `property:"name"`
	_ string `property:"path"`
}

func (i *ImageController) init() {
	// fmt.Println("init the image controller")
	i.SetImageModel(core.NewQAbstractListModel(nil))
	i.ConnectRequestExpandImage(controller.Instance().RequestExpandImage)

	i.ImageModel().ConnectRowCount(func(*core.QModelIndex) int {
		return len(i.listOfPaths)
	})
	i.ImageModel().ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		return core.NewQVariant1(i.getImagePaths()[index.Row()])
	})

	err := filepath.Walk("/Users/alex/go/src/github.com/amlwwalker/pickleit/images", i.walkingImages)
	if err != nil {
		fmt.Println("error walking images")
	}
	fmt.Println(i.listOfPaths)
}

func (i *ImageController) walkingImages(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}
	if info.IsDir() {
	} else {
		ev := NewAnImage(nil)
		ev.SetName(path)
		ev.SetPath(path)
		i.listOfPaths = append(i.listOfPaths, ev)
	}
	return nil
}

func (i *ImageController) getImagePaths() []*anImage {
	return i.listOfPaths
}
