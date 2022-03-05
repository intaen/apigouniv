package domain

type MasterDataService interface {
	FindUniversity(string, string) ([]University, error)
	FindAPI([]string, string) (*PublicAPIs, error)
	FindAPIByAuth([]string) (*PublicAPIs, error)
	FindAPIByCategory(string) (*PublicAPIs, error)
}

type MasterDataRepo interface {
	GetAllUniversity(string, string, string, string) ([]Universities, error)
	GetAllAPI(string, string, string, string) (*PublicAPIs, error)
}

type (
	University struct {
		Name    string   `json:"name"`
		Country string   `json:"country_name"`
		URL     []string `json:"url"`
	}

	// API
	Universities struct {
		Country       string   `json:"country"`
		CountryID     string   `json:"alpha_two_code"`
		StateProvince string   `json:"state-province"`
		Name          string   `json:"name"`
		Domains       []string `json:"domains"`
		WebPages      []string `json:"web_pages"`
	}

	PublicAPIs struct {
		Auths     []string      `json:"-"`
		TotalAuth int           `json:"-"`
		Count     int           `json:"count"`
		Entries   []DetailEntry `json:"entries"`
	}
	DetailEntry struct {
		API         string `json:"API"`
		Description string `json:"Description"`
		Auth        string `json:"Auth"`
		HTTPS       bool   `json:"HTTPS"`
		Cors        string `json:"Cors"`
		Link        string `json:"Link"`
		Category    string `json:"Category"`
	}
)
