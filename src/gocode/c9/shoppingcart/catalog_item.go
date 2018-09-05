package shoppingcart

import (
	"encoding/gob"
)

type CatalogItem struct {
	ItemID           string  `schema:"itemID"`
	ShortDescription string  `schema:"shortDescription"`
	LongDescription  string  `schema:"longDescription"`
	Cost             float64 `schema:"cost"`
}

func init() {
	gob.Register(&CatalogItem{})
}

var items []CatalogItem = []CatalogItem{
	CatalogItem{
		"hall001",
		"<i>Core Servlets and JavaServer Pages " +
			"2nd Edition</i> (Volume 1)" +
			" by Marty Hall and Larry Brown",
		"The definitive reference on servlets " +
			"and JSP from Prentice Hall and \n" +
			"Sun Microsystems Press.<p>Nominated for " +
			"the Nobel Prize in Literature.",
		39.95,
	},
	CatalogItem{
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
		49.99,
	},
	CatalogItem{
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
		19.95,
	},
	CatalogItem{
		"alexander001",
		"<i>The Prydain Series</i> by Lloyd Alexander",
		"Humble pig-keeper Taran joins mighty " +
			"Lord Gwydion in his battle against",
		19.95,
	},
	CatalogItem{
		"rowling001",
		"<i>The Harry Potter Series</i> by J.K. Rowling",
		"The first five of the popular stories " +
			"about wizard-in-training Harry Potter\n",
		59.95,
	}}

func GetItem(itemID string) *CatalogItem {
	if itemID == "" {
		return nil
	}
	for i := 0; i < len(items); i++ {
		if items[i].ItemID == itemID {
			return &items[i]
		}
	}
	return nil
}
