package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	button := widgets.NewQPushButton2("Click Me", nil)

	// Create Tab Widget
	tabs := widgets.NewQTabWidget(nil)

	// Create First Tab
	tab1 := widgets.NewQWidget(nil, 0)
	layout1 := widgets.NewQVBoxLayout()
	label1 := widgets.NewQLabel2("This is Tab 1", nil, 0)
	layout1.AddWidget(label1, 0, 0)
	tab1.SetLayout(layout1)

	tabs.AddTab(tab1, "Subscriptions")

	button.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "Message", "Button Clicked!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	window.SetCentralWidget(tabs)
	window.Show()
	app.Exec()
}
