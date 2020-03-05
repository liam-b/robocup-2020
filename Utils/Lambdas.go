package Utils

// maybe helpful who knows

// all are called in the form WaitUntil(func() bool { <code> })

func WaitUntil(when func() bool) {
	for !when() {
		continue
	}
}

func WaitWhile(when func() bool) {
	for when() {
		continue
	}
}

func DoWhile(what func(), when func() bool) {
	for when() {
		what()
	}
}

func DoUntil(what func(), when func() bool) {
	for !when() {
		what()
	}
}
