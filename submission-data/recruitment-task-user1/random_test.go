package task

import (
	"reflect"
	"testing"
)

func TestTask(t *testing.T) {
	html, err := ReadHTMLFromFile("data.html")
	if err != nil {
		t.Errorf("got error while reading html file: %v, want no error",
			err)
	}

	buf := CreateBuffer(html)
	root, err := CreateTree(buf)
	if err != nil {
		t.Errorf("got error while creating tree: %v, want no error", err)
	}

	gotDivTagsCount := CountDivTags(root)
	wantDivTagsCount := 8
	if gotDivTagsCount != wantDivTagsCount {
		t.Errorf("got div tags count as %d, want %d",
			gotDivTagsCount, wantDivTagsCount)
	}

	gotUniqueTagsSortedList := ExtractAllUniqueTagsInSortedOrder(root)
	wantUniqueTagsSortedList := []string{"a", "body", "div", "head", "html",
		"meta", "title"}

	if !reflect.DeepEqual(gotUniqueTagsSortedList, wantUniqueTagsSortedList) {
		t.Errorf("got tags list as %v, want %v",
			gotUniqueTagsSortedList, wantUniqueTagsSortedList)
	}

	gotCommentList := ExtractAllComments(root)
	wantCommentList := []string{"Comment 1", "Comment 2", "Comment 3",
		"Comment 4", "Comment 5", "Comment 6"}

	if !reflect.DeepEqual(gotCommentList, wantCommentList) {
		t.Errorf("got comments list as %v, want %v",
			gotCommentList, wantCommentList)
	}

	gotLinksList := ExtractAllLinks(root)
	wantLinksList := []string{"https://cfstress.com", "https://facebook.com",
		"https://google.com"}
	if !reflect.DeepEqual(gotLinksList, wantLinksList) {
		t.Errorf("got links list as %v, want %v",
			gotLinksList, wantLinksList)
	}
}
