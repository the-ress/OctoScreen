package ui

import (
	"github.com/Z-Bolt/OctoScreen/interfaces"
	"github.com/Z-Bolt/OctoScreen/utils"
)

var printMenuPanelInstance *printMenuPanel

type printMenuPanel struct {
	CommonPanel
}

func PrintMenuPanel(
	ui				*UI,
	parentPanel		interfaces.IPanel,
) *printMenuPanel {
	if printMenuPanelInstance == nil {
		instance := &printMenuPanel {
			CommonPanel: NewCommonPanel(ui, parentPanel),
		}
		instance.initialize()
		printMenuPanelInstance = instance
	}

	return printMenuPanelInstance
}

func (this *printMenuPanel) initialize() {
	defer this.Initialize()

	moveButton := utils.MustButtonImageStyle("Move",               "move.svg",           "color1", this.showMove)
	this.Grid().Attach(moveButton,        0, 0, 1, 1)

	filamentButton := utils.MustButtonImageStyle("Filament",       "filament-spool.svg", "color2", this.showFilament)
	this.Grid().Attach(filamentButton,    1, 0, 1, 1)

	temperatureButton := utils.MustButtonImageStyle("Temperature", "heat-up.svg",        "color3", this.showTemperature)
	this.Grid().Attach(temperatureButton, 2, 0, 1, 1)

	fanButton := utils.MustButtonImageStyle("Fan",                 "fan.svg",            "color4", this.showFan)
	this.Grid().Attach(fanButton,         0, 1, 1, 1)

	networkButton := utils.MustButtonImageStyle("Network",         "network.svg",        "color1", this.showNetwork)
	this.Grid().Attach(networkButton,     1, 1, 1, 1)

	systemButton := utils.MustButtonImageStyle("System",           "info.svg",           "color3", this.showSystem)
	this.Grid().Attach(systemButton,      2, 1, 1, 1)
}

func (this *printMenuPanel) showMove() {
	this.UI.Add(MovePanel(this.UI, this))
}

func (this *printMenuPanel) showFilament() {
	this.UI.Add(FilamentPanel(this.UI, this))
}

func (this *printMenuPanel) showTemperature() {
	this.UI.Add(TemperaturePanel(this.UI, this))
}

func (this *printMenuPanel) showFan() {
	this.UI.Add(FanPanel(this.UI, this))
}

func (this *printMenuPanel) showNetwork() {
	this.UI.Add(NetworkPanel(this.UI, this))
}

func (this *printMenuPanel) showSystem() {
	this.UI.Add(SystemPanel(this.UI, this))
}