package logic

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"plugin"
	"strings"
	"text/template"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
)

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == "pic" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func NewPluginManager(informer chan OperatingMessage, engine *qml.QQmlApplicationEngine) (*PluginManager, error) {
	p := &PluginManager{}
	p.path = PLUGINFOLDER
	p.informer = informer
	p.engine = engine
	if files, err := filePathWalkDir(p.path); err != nil {
		return p, err
	} else {
		//what to do with the list of files?
		p.plugins = files //note these are just the plugins found at the plugin directory. They have not yet been tested as to whether they are suitable plugins
	}
	return p, nil
}

// const viewTemplate = `
// import QtQuick 2.0
// Item {
// 	id: {{.Identity}}
// 	anchors.fill: parent
// 	Component.onCompleted: {
// 			var subComponent = Qt.createQmlObject(' \
// 			import {{.Name}} 1.0; \
// 			{{.Name}} { \
// 					width: parent.width; \
// 					height: parent.height;}', {{.Identity}});
// 			}
// 	}
// 	`

const viewTemplate = `
import {{.Name}} 1.0;
{{.Name}} {
	anchors.fill: parent
}
`
const tabTemplate = `
import QtQuick.Controls 2.4;
TabButton {
	text: qsTr("{{.Title}}")
}
`

func (p *PluginManager) InitialisePlugins() []error {
	var errs []error
	for _, pluginPath := range p.plugins {
		p.informer <- Op_EnablingPlugin.Retrieve()
		if trusted, err := p.initialisePlugin(pluginPath); err != nil {
			errs = append(errs, err)
			p.informer <- Op_PluginError.Retrieve()
		} else {
			if err := p.enablePlugin(trusted); err != nil {
				errs = append(errs, err)
				p.informer <- Op_PluginError.Retrieve()
			} else {
				p.informer <- Op_PluginEnabled.Retrieve()
			}
		}
	}
	return errs
}

func (p *PluginManager) initialisePlugin(path string) (CustomPlugin, error) {

	tmp, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	untrusted, err := tmp.Lookup("Plugin")
	if err != nil {
		return nil, err
	}
	// var customPlugin CustomPlugin
	trusted, ok := untrusted.(CustomPlugin)
	if !ok {
		return nil, errors.New("unexpected type from module symbol - plugin not a plugin")
	}
	trusted.Init()

	return trusted, nil
}

func (p *PluginManager) enablePlugin(c CustomPlugin) error {

	//here we should pass the api to the plugin
	// ...
	//if the engine is not nil we can configure the gui
	if p.engine == nil {
		//we can't configure the gui without an engine
		return errors.New("There is no gui enabled. Configuration completing early.")
	}
	if err := p.configureTitle(c); err != nil {
		return err
	}
	if err := p.configureView(c); err != nil {
		return err
	}
	return nil
}
func (p *PluginManager) configureTitle(c CustomPlugin) error {
	type templateObject struct {
		Title string
	}
	var obj templateObject
	obj.Title = strings.Title(strings.ToLower(c.Name()))
	tmpl, err := template.New("qmlTitle").Parse(tabTemplate)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, obj)
	if err != nil {
		return err
	}
	var qmlString = buf.String()
	p.addToGui("tabBar", qmlString)
	return nil
}
func (p *PluginManager) configureView(c CustomPlugin) error {
	type templateObject struct {
		Name     string
		Identity string
	}
	var obj templateObject
	if len(strings.Split(c.Name(), " ")) > 1 {
		return errors.New("name must be a single word")
	}
	obj.Name = strings.Title(strings.ToLower(c.Name()))
	obj.Identity = strings.ToLower(c.Name())
	tmpl, err := template.New("qmlStackView").Parse(viewTemplate)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, obj)
	if err != nil {
		return err
	}
	var qmlString = buf.String()
	p.addToGui("stackLayout", qmlString)
	return nil
}
func (p *PluginManager) addToGui(childElement, qmlString string) error {
	//https://stackoverflow.com/questions/31890372/add-objects-to-qml-layout-from-c%5D
	if p.engine == nil {
		return errors.New("Engine not intialised")
	}
	//get a reference to the object you want to put this as a child of
	stackLayout := p.engine.RootObjects()[0].FindChild(childElement, core.Qt__FindChildrenRecursively)
	stackLayoutPointer := quick.NewQQuickItemFromPointer(stackLayout.Pointer())
	//create a component to hold the child element
	mainComponent := qml.NewQQmlComponent2(p.engine, nil)
	mainComponent.SetData(core.NewQByteArray2(qmlString, -1), core.NewQUrl())
	//create the actual component as a qml item
	item := quick.NewQQuickItemFromPointer(mainComponent.Create(p.engine.RootContext()).Pointer())
	p.engine.SetObjectOwnership(item, qml.QQmlEngine__JavaScriptOwnership)
	// //specify the parent
	item.SetParent(stackLayout)
	item.SetParentItem(stackLayoutPointer)
	return nil
}
