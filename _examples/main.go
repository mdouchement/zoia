package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/emicklei/dot"
	"github.com/mdouchement/zoia"
	"github.com/mdouchement/zoia/binary"
	"github.com/mdouchement/zoia/datatype"
	"github.com/mdouchement/zoia/module"
	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	data, err := ioutil.ReadFile(os.Args[2])
	check(err)
	patch := zoia.Parse(data)

	switch os.Args[1] {
	case "yaml":
		payload, err := yaml.Marshal(patch)
		check(err)
		fmt.Println(string(payload))
	case "graph":
		g := dot.NewGraph(dot.Directed)

		modules := make([]*label, len(patch.Modules))
		for i, m := range patch.Modules {
			modules[i] = &label{
				graph:  g,
				idx:    i,
				module: m,
			}
		}
		for _, cable := range patch.PatchCables {
			modules[cable.SourceModule].add("source", cable)
			modules[cable.DestinationModule].add("destination", cable)
		}

		//

		for i, cable := range patch.PatchCables {
			modules[cable.SourceModule].nodeField(cable.SourceSignal).
				Edge(modules[cable.DestinationModule].nodeField(cable.DestinationSignal), fmt.Sprintf("%d\n(%s)", i, cable.Gain)).
				Attr("shape", "vee")
		}

		fmt.Println(g.String())
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: go run . yaml|graph patch.bin")
	fmt.Println("     - go run . yaml /Volumes/ZOIA/from_zoia/010_zoia_Play_Thing.bin")
	fmt.Println("     - go run . graph /Volumes/ZOIA/from_zoia/010_zoia_Play_Thing.bin | dot -Tpng > 010_zoia_Play_Thing.png")
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

////////////////////////
//                    //
// Module             //
//                    //
////////////////////////

type label struct {
	graph *dot.Graph
	nod   *dot.Node

	idx        int
	module     module.Module
	cablesidx  []datatype.PatchCableSignal
	cablesmidx map[datatype.PatchCableSignal]int
	cables     map[datatype.PatchCableSignal]*lcable
}

func (l *label) add(side string, c *zoia.PatchCable) {
	if l.cables == nil {
		l.cablesmidx = make(map[datatype.PatchCableSignal]int, 0)
		l.cables = make(map[datatype.PatchCableSignal]*lcable, 0)
	}

	var signal datatype.PatchCableSignal

	switch side {
	case "source":
		signal = c.SourceSignal
	case "destination":
		signal = c.DestinationSignal
	default:
		panic("bad value")
	}

	if _, ok := l.cables[signal]; !ok {
		l.cablesmidx[signal] = len(l.cables)
		l.cablesidx = append(l.cablesidx, signal)
	}

	l.cables[signal] = &lcable{
		side:   side,
		signal: signal,
	}
}

func (l *label) node() dot.Node {
	if l.nod != nil {
		return *l.nod
	}

	nod := l.graph.Node(l.string()).Attr("color", l.color()).Attr("shape", "Mrecord")
	l.nod = &nod
	return *l.nod
}

func (l *label) nodeField(id datatype.PatchCableSignal) dot.Node {
	return l.node().Field(fmt.Sprintf("c%04d", id))
}

func (l *label) string() string {
	base := l.module.GetBase()
	name := base.Type.String()
	if base.Name != "" {
		name = fmt.Sprintf("%s\n\"%s\"", base.Type, base.Name)
	}
	label := fmt.Sprintf("%d | { %d | %d } | %s", l.idx, base.Page, base.Position, name)

	//

	var params []string

	// var cables int
	for _, p := range l.module.Params() {
		// TODO: Find a way to identify patch cable's ports connection.
		// Cables are stored in bin in order of creation.
		// Currently we do not have information about the ports on which a cable is plugged.

		if p.Pluggable {
			continue
		}
		// if p.Pluggable {
		// 	if cables == len(l.cablesidx) {
		// 		params = params[:0]
		// 		break
		// 	}

		// 	id := l.cablesidx[cables]
		// 	params = append(params, fmt.Sprintf("<c%04d> %s", id, l.cables[id].name(&p)))
		// 	cables++
		// 	continue
		// }

		if p.Attribute != nil {
			params = append(params, fmt.Sprintf("%s\n%s", p.Name, p.Attribute))
			continue
		}
		params = append(params, p.Name)
	}

	for id, cable := range l.cables {
		params = append(params, fmt.Sprintf("<c%04d> %s", id, cable.name(nil)))
	}
	sort.Strings(params)

	return fmt.Sprintf("%s | %s", label, strings.Join(params, " | "))
}

func (l *label) color() string {
	switch c := l.module.GetBase().Extra.Color; c {
	case binary.ColorExtraYello:
		return "#E5E500"
	case binary.ColorExtraWhite:
		return "#C8C8CE"
	default:
		return c.HexColor()
	}
}

////////////////////////
//                    //
// Cable              //
//                    //
////////////////////////

type lcable struct {
	side   string
	signal datatype.PatchCableSignal
}

func (c *lcable) name(params *module.Parameter) string {
	if params != nil {
		if params.Attribute != nil {
			return fmt.Sprintf("%s\n%s", params.Name, params.Attribute)
		}
		return params.Name
	}
	return c.signal.String()
}

func countCables(params []module.Parameter) (n int) {
	for _, p := range params {
		if p.Pluggable {
			n++
		}
	}
	return n
}
