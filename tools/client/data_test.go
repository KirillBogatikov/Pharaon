package client

const (
	apiPostUrl = "https://petstore.swagger.io/v2/pet"
	apiUrl     = "https://petstore.swagger.io/v2/pet/%d"
)

var (
	natasha = &Pet{
		Id:   0,
		Name: "Natasha",
		Category: Category{
			Id:   154,
			Name: "Chicken",
		},
		PhotoUrls: []string{
			"test.png",
		},
		Tags: []Tag{
			{Id: 179, Name: "tag-179"},
		},
		Status: "available",
	}
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Pet struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Category  Category `json:"category"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags"`
	Status    string   `json:"available"`
}

func equals(a, b *Pet) bool {
	if a.Id == b.Id &&
		a.Category.Id == b.Category.Id &&
		a.Category.Name == b.Category.Name &&
		a.Name == b.Name {

		for i := range a.Tags {
			aTag := a.Tags[i]
			bTag := b.Tags[i]

			if aTag.Id != bTag.Id || aTag.Name != bTag.Name {
				return false
			}
		}

		for i := range a.PhotoUrls {
			if a.PhotoUrls[i] != b.PhotoUrls[i] {
				return false
			}
		}

		return true
	}

	return false
}
