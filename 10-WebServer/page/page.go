// package to represent a page to render
package page

import "io/ioutil"

// declaration of Page structure
type Structure struct {
    Title string
    Body []byte
}

// save a Page in /data/p.Title+".txt" with content of p.Body
func (p* Structure) Save() error {
    filename := "data/"+ p.Title +".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

// load a current file into a Page structure
func Load(title string) (*Structure, error) {
    filename := "data/"+ title +".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Structure{Title:title, Body:body}, nil
}

