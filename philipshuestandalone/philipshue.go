package philipshue

import (
	"fmt"
	"log"

	"gbbr.io/hue"
	"github.com/lucasb-eyer/go-colorful"
)

type PhilipsHue struct {
	settings string
	bridge   *hue.Bridge
}

func NewPhilipsHue(settings string) *PhilipsHue {
	p := PhilipsHue{}
	p.settings = settings
	p.setup()
	return &p
}

func (this *PhilipsHue) setup() {
	bridge, err := hue.Discover()
	if err != nil {
		log.Fatal(err)
	}
	if !bridge.IsPaired() {
		// link button must be pressed for non-error response
		if err := bridge.Pair(); err != nil {
			log.Fatal(err)
		}
	}
	this.bridge = bridge
}

func (this *PhilipsHue) turnLightOn(name string) {
	this.switchLight(name, true)
}

func (this *PhilipsHue) turnLightOff(name string) {
	this.switchLight(name, false)
}

func (this *PhilipsHue) switchLight(name string, on bool) {
	light, err := this.bridge.Lights().Get(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	if on {
		err = light.On()
	} else {
		err = light.Off()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	if on {
		fmt.Println("Turned", name, "on")
	} else {
		fmt.Println("Turned", name, "off")

	}
}

func (this *PhilipsHue) setLightColor(name string, hex string) {
	c, err := colorful.Hex(hex)
	if err != nil {
		fmt.Println(err)
	}
	x, y, _ := c.Xyy()
	l, err := this.bridge.Lights().Get(name)
	if err != nil {
		fmt.Println(err)
	}
	if !l.State.On {
		this.turnLightOn(name)
	}

	err = l.Set(&hue.State{
		TransitionTime: 0,
		Brightness:     255,
		XY:             &[2]float64{x, y},
	})
	if err != nil {
		fmt.Println(err)
	}
}
