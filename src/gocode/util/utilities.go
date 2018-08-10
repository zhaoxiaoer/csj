package util

const DOCTYPE = "<!DOCTYPE html>\n"

func HeadWithTitle(title string) string {
	return DOCTYPE +
		"<html>\n" +
		"<head>\n" +
		"<title>" + title + "</title>\n" +
		"</head>\n"
}
