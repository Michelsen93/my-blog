package data

type PostCategory struct {
	Name  string
	CSSClass string
	Posts []BlogMetaData
}

type BlogMetaData struct {
	Title string
	Path  string
}

var techPosts = []BlogMetaData{
	{
		Title: "Ny på cloud",
		Path:  "/my-blog/generelt/cloud_IaC.md",
	},
	{
		Title: "Hei bloggen",
		Path:  "/my-blog/generelt/hei_bloggen.md",
	},
}

var generalPosts = []BlogMetaData{
	{
		Title: "Test",
		Path:  "Test",
	},
}

var sportPosts = []BlogMetaData{
	{
		Title: "Sport",
		Path:  "Test",
	},
}

func InitialPost() BlogMetaData {
    return techPosts[len(techPosts)-1]
}

func PostLists() []PostCategory {
    return []PostCategory{
        {
            Name: "Sport",
            CSSClass: "#f4f4f4",
            Posts: sportPosts,
        },
        {
            Name: "Teknologi",
            CSSClass: "#f4f4f4",
            Posts: techPosts,
        },
        {
            Name: "Generelt",
            CSSClass: "#f4f4f4",
            Posts: generalPosts,
        },
    }
}
