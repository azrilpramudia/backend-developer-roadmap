package labels

type Label interface {
	SetLabel(label string)
	GetLabel() (string, string)
	SetBarCode(barCode string)
	GetBarCode() string
}

type LabelImpl struct {
	Label string
	BarCode string
}

func NewLabel(label string, barCode string) Label {
	return &LabelImpl{
		Label: label,
		BarCode: barCode,
	}
}

func (l *LabelImpl) SetLabel(label string) {
	l.Label = label
}

func (l *LabelImpl) GetLabel() (string, string) {
	return l.Label, l.BarCode
}

func (l *LabelImpl) SetBarCode(barCode string) {
	l.BarCode = barCode
} 

func (l *LabelImpl) GetBarCode() string {
	return l.BarCode
}