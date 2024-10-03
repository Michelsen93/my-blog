package posts

import "github.com/Michelsen93/my-blog/model"

func BlogPosts() []model.BlogMetaData {
    return []model.BlogMetaData{
        {   
            Title: "Ny på cloud", 
            Path: "/my-blog/generelt/cloud_IaC.md",
        },
        {   
            Title: "Hei bloggen", 
            Path: "/my-blog/generelt/hei_bloggen.md",
        },
    }
}
