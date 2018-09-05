package c9.shoppingcart;

public class Catalog {
	private static CatalogItem[] items =
		{
				new CatalogItem(
						"hall001",
						"<i>Core Servlets and JavaServer Pages " +
						  "2nd Edition</i> (Volume 1)" +
						  " by Marty Hall and Larry Brown",
						"The definitive reference on servlets " +
						  "and JSP from Prentice Hall and \n" +
						  "Sun Microsystems Press.<p>Nominated for " +
						  "the Nobel Prize in Literature.",
						39.95
						),
				new CatalogItem(
						"hall002",
						"<i>Core Web Programming, 2nd Edition</i> " +
						  "by Marty Hall and Larry Brown",
						"One stop shopping for the Web programmer. " +
						  "Topics include \n" +
						  "<ul><li>Thorough coverage of Java 2; " +
						  "including Threads, Networking, Swing, \n" +
						  "Java 2D, RMI, JDBC, and Collections</li>\n" +
						  "<li>A fast introduction to HTML 4.01, " +
						  "including frames, style sheets, and layers.</li>\n" +
						  "<li>A fast introduction to HTTP 1.1, " +
						  "servlets, and JavaServer Pages.</li>\n" +
						  "<li>A quick overview of JavaScript 1.2</li>\n" +
						  "</ul>",
						49.99
						),
				new CatalogItem(
						"lewis001",
						"<i>The Chronicles of Narnia</i> by C.S. Lewis",
						"The classic children's adventure pitting " +
						  "Aslan the Great Lion and his followers\n" +
						  "against the White Witch and the forces " +
						  "of evil. Dragons, magicians, quests, \n" +
						  "and talking animals wound around a deep " +
						  "spiritual allegory. Series includes\n" +
						  "<i>The Magician's Nephew</i>,\n" +
						  "<i>The Lion, the Witch and the Wardrobe</i>,\n" +
						  "<i>The horse and His Boy</i>,\n" +
						  "<i>Prince Caspian</i>,\n",
						19.95
						),
				new CatalogItem(
						"alexander001",
						"<i>The Prydain Series</i> by Lloyd Alexander",
						"Humble pig-keeper Taran joins mighty " +
						  "Lord Gwydion in his battle against",
						19.95
						),
				new CatalogItem(
						"rowling001",
						"<i>The Harry Potter Series</i> by J.K. Rowling",
						"The first five of the popular stories " +
						  "about wizard-in-training Harry Potter\n",
						59.95
						),
		};
	
	public static CatalogItem getItem(String itemID) {
		CatalogItem item;
		if (itemID == null) {
			return null;
		}
		for (int i = 0; i < items.length; i++) {
			if (itemID.equals(items[i].getItemID())) {
				return items[i];
			}
		}
		return null;
	}
}
