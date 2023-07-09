package sructurs

type MasPatern struct {
	MasPatternPath []struct {
		NumberInAggoritm string `json:"NumberInAggoritm"`
		PathInEX         string `json:"PathInEX"`
		PaternInEX       string `json:"PaternInEX"`
	} `json:"MasPatternPath"`
}

type PathAndPatern struct {
	Path1 string //это тестовые пути
	Path2 string
	Path3 string
	Path4 string
	Path5 string

	Patern1 string
	Patern2 string
	Patern3 string
	Patern4 string
	Patern5 string
}

type ExploitJson struct {
	MasTemplates []struct {
		NameTeamplates string `json:"Name_Teamplates"`
		Exploits       struct {
			Ex1 string `json:"ex1"`
			Ex2 string `json:"ex2"`
			Ex3 string `json:"ex3"`
		} `json:"Exploits"`
	} `json:"Mas_Templates"`
}
