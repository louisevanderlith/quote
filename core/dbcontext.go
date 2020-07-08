package core

import "github.com/louisevanderlith/husk"

type context struct {
	Submissions husk.Tabler
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
