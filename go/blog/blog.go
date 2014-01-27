package blog

import (
    "os"
    "log"
    "net/http"
    "html/template"
    "path/filepath"
    "io/ioutil"

    "github.com/russross/blackfriday"
)

// post is a single blog entry.
type post struct {
    Body template.HTML
}

// Blog is a blog site.
type Blog struct {
    // TplPath is the relative path to the HTML template for the Blog.
    TplPath string
    // PostDir is the relative path to the directory containing blog posts.
    PostDir string
}

// NewBlog returns a new Blog instance.
func NewBlog(tplPath, postDir string) *Blog {
    return &Blog{TplPath: tplPath, PostDir: postDir}
}

func (self Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tpl, err := template.ParseFiles(self.TplPath)
    if err != nil {
        log.Fatal(err)
    }

    // data is the template data for the Blog.
    data := struct {
        // Testimonials is the slice of testimonials.
        Posts []post
    } {}

    buildPosts := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Printf("Error walking: %v\n", err)
            return nil
        }

        if !info.IsDir() {
            postFile, err := os.Open(path)
            if err != nil {
                log.Printf("Error opening blog post file: %v\n", err)
                return nil
            }
            postMarkdown, err := ioutil.ReadAll(postFile)
            if err != nil {
                log.Printf("Error reading blog post file markdown %v\n", err)
            }
            postHTML := template.HTML(blackfriday.MarkdownCommon(postMarkdown))
            data.Posts = append(data.Posts, post{Body: postHTML})
        }
        return nil
    }

    filepath.Walk(self.PostDir, buildPosts)
    // Reverse order.
    for i, j := 0, len(data.Posts)-1; i < j; i, j = i+1, j-1 {
        data.Posts[i], data.Posts[j] = data.Posts[j], data.Posts[i]
    }

    tpl.Execute(w, data)
}

