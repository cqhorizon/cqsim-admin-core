package pkg

type (
	Mode string
)

const (
	ModeProd Mode = "prod" //生产模式
)

func (e Mode) String() string {
	return string(e)
}
