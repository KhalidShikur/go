package airportrobot

type Greeter interface {
    LanguageName() string
    Greet(vname string) string
}

type Italian struct {
    vname string
}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(vname string) string {
    return "Ciao " + vname + "!"
}

type Portuguese struct {
    vname string
}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(vname string) string {
    return "Olá " + vname + "!"
}

func SayHello(vname string, xgreeter Greeter) string {
    s1 := xgreeter.LanguageName()
    s2 := xgreeter.Greet(vname)
    return "I can speak " + s1 + ": " + s2
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
