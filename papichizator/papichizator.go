package papichizator


type Papichizable interface {
	Papichize(text string)(string)
}

type Papichizator struct {
}

func (p *Papichizator) Papichize(text string) string  {

	return  text + "ич"
}