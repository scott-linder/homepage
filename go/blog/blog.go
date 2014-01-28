package blog

import (
    "os"
    "fmt"
    "log"
    "time"
    "strings"
    "net/http"
    "html/template"
    "path/filepath"
    "io/ioutil"

    "github.com/gorilla/mux"
    "github.com/russross/blackfriday"
)

const (
    // dateFormat is the expected format of filename dates.
    dateFormat = "2006-01-02"
)

// post is a single blog entry.
type post struct {
    // Name is the name of the post for perma-link purposes.
    Name string
    // Body is the content of the post.
    Body template.HTML
    // Date is the date the post was published.
    Date time.Time
    // Permalink is a permanant URL pointing to this post.
    Permalink string
}

// Blog is a blog site.
type Blog struct {
    // Router is the mux (sub)router which Blog works under.
    Router *mux.Router
    // TplPath is the relative path to the HTML template for the Blog.
    TplPath string
    // PostDir is the relative path to the directory containing blog posts.
    PostDir string
    // PageSize is the number of posts to a page.
    PageSize int
}

// NewBlog returns a new Blog instance.
func NewBlog(router *mux.Router, tplPath, postDir string, pageSize int) (blog *Blog) {
    blog = &Blog{Router: router, TplPath: tplPath, PostDir: postDir, PageSize: pageSize}
    router.Handle("/", blog)
    router.Handle("/post/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}/{name}/",
                    blog).Name("post")
    return
}

func (self Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    tpl, err := template.ParseFiles(self.TplPath)
    if err != nil {
        log.Fatal(err)
    }

    // data is the template data for the Blog.
    data := struct {
        // Posts is the slice of posts for this blog page.
        Posts []post
    } {}

    vars := mux.Vars(r)

    switch mux.CurrentRoute(r).GetName() {
    case "post":
        data.Posts = make([]post, 1)
        data.Posts[0], err = self.getPost(vars["year"], vars["month"],
                                            vars["day"], vars["name"])
    default:
        data.Posts, err = self.getPage()
    }


    tpl.Execute(w, data)
}

func (self Blog) newPost(path string, info os.FileInfo) post {
    nameFields := strings.Split(info.Name(), ".")
    var dateField, nameField string
    if len(nameFields) >= 2 {
        dateField = nameFields[0]
        nameField = nameFields[1]
    }
    postDate, err := time.ParseInLocation(dateFormat, dateField, time.UTC)
    if err != nil {
        log.Fatalf("Error parsing date from %v; skipping: %v\n",
                    info.Name(), err)

    }
    postName := nameField

    postPermalinkURL, err := self.Router.Get("post").
        URL("year", fmt.Sprintf("%d", postDate.Year()),
            "month", fmt.Sprintf("%d", postDate.Month()),
            "day", fmt.Sprintf("%d", postDate.Day()),
            "name", postName)
    if err != nil {
        log.Printf("Error creating permalink: %v\n", err)
    }
    postPermalink := postPermalinkURL.String()

    postFile, err := os.Open(path)
    if err != nil {
        log.Fatalf("Error opening blog post file: %v\n", err)
    }
    postMarkdown, err := ioutil.ReadAll(postFile)
    if err != nil {
        log.Printf("Error reading blog post file markdown %v\n", err)
    }
    postHTML := template.HTML(blackfriday.MarkdownCommon(postMarkdown))
    return post{Name: postName, Body: postHTML, Date: postDate, Permalink: postPermalink}
}

func (self Blog) getPost(year, month, day, name string) (post, error) {
    path := self.PostDir + fmt.Sprintf("%04s-%02s-%02s.%s.md", year, month, day, name);
    info, err := os.Stat(path)
    if err != nil {
        log.Fatalf("Unable to stat blog post: %v\n", err)
    }
    return self.newPost(path, info), nil
}

func (self Blog) getPage() (posts []post, err error) {
    buildPosts := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Printf("Error walking: %v\n", err)
            return nil
        }

        if !info.IsDir() {
            posts = append(posts, self.newPost(path, info))
        }
        return nil
    }

    filepath.Walk(self.PostDir, buildPosts)
    // Reverse order.
    for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
        posts[i], posts[j] = posts[j], posts[i]
    }

    return posts, nil
}
