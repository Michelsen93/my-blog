package posts

import "github.com/Michelsen93/my-blog/model"

func BlogPosts() []model.BlogMetaData {
    return []model.BlogMetaData{
        {   
            Title: "Consistency", 
            Category: model.Category{
                Name: "Sport",  
                Color: "#444",
            },
            Path: "/sport/consistency.md",
        },
    }
}
