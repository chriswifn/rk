package rk

import (
	"bufio"
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
	Version:     `v0.1.2`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/rk.git`,
	Issues:      `github.com/chriswifn/rk/issues`,
	Commands:    []*Z.Cmd{compareCmd, help.Cmd, vars.Cmd, conf.Cmd, initCmd, filterCmd},
	Summary:     help.S(_rk),
	Description: help.D(_rk),
}

var initCmd = &Z.Cmd{
	Name:        `init`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_init),
	Description: help.D(_init),
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
				fmt.Printf("%s,%s,%f\n", files[i], files[j], checker.GetRate())
			}
		}
		return nil
	},
}

var filterCmd = &Z.Cmd{
	Name:        `filter`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_filter),
	Description: help.D(_filter),
	Call: func(x *Z.Cmd, _ ...string) error {
		var files []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			path := scanner.Text()
			fileInfo, err := os.Stat(path)
			if err != nil {
				return err
			}
			if fileInfo.IsDir() {
				continue
			} else {
				files = append(files, path)
			}
		}
		for i := 0; i < len(files)-1; i++ {
			for j := i + 1; j < len(files); j++ {
				checker := NewPlagarismChecker(
					files[i],
					files[j],
				)
				fmt.Printf("%s,%s,%f\n", files[i], files[j], checker.GetRate())
			}
		}
		return nil
	},
}
