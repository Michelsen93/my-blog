package model

type Category struct {
    Name string
    Color string
}

type BlogMetaData struct {
    Title string
    Category
    Path string
}
