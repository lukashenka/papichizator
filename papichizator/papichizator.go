package papichizator


type Papichizable interface {
	Papichize(text string)(string)
}

type Papichizator struct {
	toPapichize string
}

func (p *Papichizator) Papichize(text string) string  {

	return  p.toPapichize
}