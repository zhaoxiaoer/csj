package util;

public class SearchUtilities {
	private static SearchSpec[] commonSpecs = {
			new SearchSpec("Google", "http://www.google.com/search?q="),
			new SearchSpec("Yahoo", "http://search.yahoo.com/bin/search?p="),
			new SearchSpec("baidu", "http://www.baidu.com/baidu?word="),
	};
	
	public static SearchSpec[] getCommonSpecs() {
		return commonSpecs;
	}
	
	public static String makeURL(String searchEngineName, String searchString) {
		SearchSpec[] searchSpecs = getCommonSpecs();
		String searchURL = null;
		for (int i = 0; i < searchSpecs.length; i++) {
			SearchSpec spec = searchSpecs[i];
			if (spec.getName().equalsIgnoreCase(searchEngineName)) {
				searchURL = spec.makeURL(searchString);
				break;
			}
		}
		return searchURL;
	}
}
