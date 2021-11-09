package wevtapi

import "syscall"

/* Create a new, empty bookmark. Bookmark handles must be closed with CloseEventHandle. */
func CreateBookmark() (EVT_HANDLE, error) {
	bookmark, err := EvtCreateBookmark(nil)
	if err != nil {
		return 0, err
	}
	return bookmark, nil
}

/* Create a bookmark from a XML-serialized bookmark. Bookmark handles must be closed with CloseEventHandle. */
func CreateBookmarkFromXmlString(xmlString string) (EVT_HANDLE, error) {
	wideXmlString, err := syscall.UTF16PtrFromString(xmlString)
	if err != nil {
		return 0, err
	}
	bookmark, err := EvtCreateBookmark(wideXmlString)
	if bookmark == 0 {
		return 0, err
	}
	return bookmark, nil
}

/* Update a bookmark to store the channel and ID of the given event */
func UpdateBookmark(bookmarkHandle EVT_HANDLE, eventHandle EVT_HANDLE) error {
	return EvtUpdateBookmark(bookmarkHandle, eventHandle)
}

/* Serialize the bookmark as XML */
func RenderBookmark(bookmarkHandle EVT_HANDLE) (string, error) {
	ret, err := EvtRenderBook(bookmarkHandle)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}
