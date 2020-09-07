package core

import "github.com/louisevanderlith/husk"

type context struct {
	Submissions husk.Table
}

var ctx context

func CreateContext() {
	ctx = context{
		Submissions: husk.NewTable(Submission{}),
	}
}

func Shutdown() {
	ctx.Submissions.Save()
}
