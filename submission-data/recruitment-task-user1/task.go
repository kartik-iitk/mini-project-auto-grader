package task

import (
	"bytes"
	"io/ioutil"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

// ReadHTMLFromFile should read the file from the current directory, if it exists.
// The file data should be returned as a string.
func ReadHTMLFromFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	return string(data), err
}

// CreateBuffer should transfer the contents of a string to a buffer.
func CreateBuffer(data string) bytes.Buffer {
	return *bytes.NewBufferString(data)
}

// CreateTree should create the tree representation of HTML represented by the buffer.
func CreateTree(buf bytes.Buffer) (*html.Node, error) {
	doc, err := html.Parse(&buf)
	//bytes.Buffer already implements io.Reader but only when it is a pointer, so need of html.Parse(bytes.NewReader(buf.Bytes())).
	return doc, err
}

// CountDivTags should return the count of <div> tags in the document tree.
func CountDivTags(node *html.Node) int {
	count := 0
	if node.Type == html.ElementNode && node.Data == "div" {
		count++
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		count += CountDivTags(c)
	}
	return count
}

// dfs is a utility function which will help you count the number of unique tags.
func dfs(node *html.Node, tagsCount map[string]int) {
	if node.Type == html.ElementNode {
		tagsCount[node.Data] += 1
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, tagsCount)
	}
}

// ExtractAllUniqueTagsInSortedOrder should return the unique tags in the document.
// These tags should also be sorted alphabetically.
func ExtractAllUniqueTagsInSortedOrder(node *html.Node) []string {
	tagsCount := make(map[string]int)             //Initialise map
	dfs(node, tagsCount)                          //Store in the map tag versus count
	uniqTags := make([]string, 0, len(tagsCount)) //create an empty slice to unique tags
	for key, _ := range tagsCount {               //Iterate over the map to fill the slice
		uniqTags = append(uniqTags, key)
	}
	sort.Strings(uniqTags)
	return uniqTags
}

// ExtractAllComments returns the list of all comments as they appear in the document.
// You also need to remove all the leading and trailing spaces in the comments.
// HINT: You might need to read about variadic functions.
func ExtractAllComments(node *html.Node) []string {
	comments := make([]string, 0)
	if node.Type == html.CommentNode {
		comments = append(comments, strings.TrimSpace(node.Data))
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		//In the following line, we will append and not simply call the function, as we aren't passing the string as a parameter to the function!
		comments = append(comments, ExtractAllComments(c)...) //Variadic Function. Necessary, as number of arguments keep on growing (unknown).
	}
	return comments
}

// ExtractAllLinks returns all the links in the document, in order of appearance.
func ExtractAllLinks(node *html.Node) []string {
	links := make([]string, 0)
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attribute := range node.Attr { //Iterate over all the attributes of the tag. Attributes are like maps for range.
			if attribute.Key == "href" {
				links = append(links, attribute.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		//Following line is not common. Here also we will append as we are not passing the string as a parameter to the function!
		links = append(links, ExtractAllLinks(c)...) //Variadic Function. Necessary, as number of arguments keep on growing (unknown).
	}
	return links
}
