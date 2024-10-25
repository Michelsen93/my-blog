package data

type PostCategory struct {
	Name     string
	CSSClass string
	Posts    []BlogMetaData
}

type BlogMetaData struct {
	Title string
	Path  string
}

var techPosts = []BlogMetaData{
	{
		Title: "Ny på cloud",
		Path:  "/my-blog/teknologi/cloud_IaC.md",
	},
	{
		Title: "Hei bloggen",
		Path:  "/my-blog/teknologi/hei_bloggen.md",
	},
	{
		Title: "Farvel Go",
		Path:  "/my-blog/teknologi/farvel_go.md",
	},
}

func InitialPost() BlogMetaData {
	return techPosts[len(techPosts)-1]
}

func PostLists() []PostCategory {
	return []PostCategory{
		{
			Name:     "Teknologi",
			CSSClass: "technology",
			Posts:    techPosts,
		},
	}
}
