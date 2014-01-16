package main

import (
    "os"
    "log"
    "net/http"
    "html/template"
    "path/filepath"
    "io/ioutil"

    "github.com/russross/blackfriday"
)

type testimonial struct {
    Body template.HTML
}

type Testimonials struct {}

func (self Testimonials) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tpl, err := template.ParseFiles("../template/testimonials.tpl")
    if err != nil {
        log.Fatal(err.Error())
    }

    data := struct {
        Testimonials []testimonial
    } {}

    buildTestimonials := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Printf("Error walking: %v\n", err)
            return nil
        }

        if !info.IsDir() {
            testimonialFile, err := os.Open(path)
            if err != nil {
                log.Printf("Error opening testimonial file: %v\n", err)
                return nil
            }
            testimonialMarkdown, err := ioutil.ReadAll(testimonialFile)
            if err != nil {
                log.Printf("Error reading testimonial file markdown %v\n", err)
            }
            testimonialHTML := template.HTML(blackfriday.MarkdownCommon(testimonialMarkdown))
            data.Testimonials = append(data.Testimonials,
                        testimonial{Body: testimonialHTML})
        }
        return nil
    }

    filepath.Walk("../testimonials/", buildTestimonials)
    // Reverse order
    for i, j := 0, len(data.Testimonials)-1; i < j; i, j = i+1, j-1 {
        data.Testimonials[i], data.Testimonials[j] = data.Testimonials[j], data.Testimonials[i]
    }

    tpl.Execute(w, data)
}

