package main

import (
	"cni_plugin_demo/pkg/types"
	"cni_plugin_demo/pkg/version"
	"fmt"
	"github.com/containernetworking/cni/pkg/skel"
	cniVersion "github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ns"
)

func main() {
	skel.PluginMain(
		cmdAdd,
		nil,
		cmdDel,
		cniVersion.PluginSupports("0.1.0", "0.2.0", "0.3.0", "0.4.0"),
		"KubeCon CNI "+version.Version)
}

func cmdAdd(args *skel.CmdArgs) (err error) {
	n, err = types.LoadNetConf(args.StdinData)
	if err != nil {
		err = fmt.Errorf("Error parsing CNI configuration \"%s\": %s", args.StdinData, err)
		return
	}
	return nil
}

func cmdDel(args *skel.CmdArgs) (err error) {
	netNS, err := ns.GetNS(args.Netns)
	if err != nil {
		return err
	}
	defer netNS.Close()
	return nil
}
