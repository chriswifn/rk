package rk

import (
	_ "embed"
	"fmt"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/to"
	"github.com/rwxrob/vars"
	"os"
)

var (
	comment = `%`
)

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
	Z.Dynamic[`dcomment`] = func() string { return comment }
}

var Cmd = &Z.Cmd{
	Name:        `rk`,
	Aliases:     []string{},
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/rk.git`,
	Issues:      `github.com/chriswifn/rk/issues`,
	Commands:    []*Z.Cmd{compareCmd, help.Cmd, vars.Cmd, conf.Cmd, initCmd},
	Summary:     help.S(_rk),
	Description: help.D(_rk),
}

var initCmd = &Z.Cmd{
	Name:     `init`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, _ ...string) error {
		val, _ := x.Caller.C(`comment`)
		if val == "null" {
			val = comment
		}
		x.Caller.Set(`comment`, val)
		return nil
	},
}

var compareCmd = &Z.Cmd{
	Name:        `compare`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_compare),
	Description: help.D(_compare),
	MinArgs:     1,
	MaxArgs:     1,
	Call: func(x *Z.Cmd, args ...string) error {
		filename := to.String(args[0])
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		files, err := WalkDir(path, filename)
		if err != nil {
			return err
		}

		for i := 0; i < len(files)-1; i++ {
			for j := i + 1; j < len(files); j++ {
				checker := NewPlagarismChecker(
					files[i],
					files[j],
				)
				// fmt.Printf("[%s-%s]\nPropability: %f\n", files[i], files[j], checker.GetRate())
				fmt.Printf("%s,%s,%f\n", files[i], files[j], checker.GetRate())
			}
		}
		return nil
	},
}
