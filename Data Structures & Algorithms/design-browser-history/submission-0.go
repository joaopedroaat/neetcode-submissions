type HistoryNode struct {
	url  string
	prev *HistoryNode
	next *HistoryNode
}

type BrowserHistory struct {
	homepage    *HistoryNode
	currentPage *HistoryNode
	lastPage    *HistoryNode
}

func Constructor(homepage string) BrowserHistory {
	homeNode := HistoryNode{
		url: homepage,
	}

	return BrowserHistory{
		homepage:    &homeNode,
		currentPage: &homeNode,
		lastPage:    &homeNode,
	}
}

func (bh *BrowserHistory) Visit(url string) {
	bh.currentPage.next = &HistoryNode{
		url:  url,
		prev: bh.currentPage,
	}
	bh.currentPage = bh.currentPage.next
	bh.lastPage = bh.currentPage
}

func (bh *BrowserHistory) Back(steps int) string {
	for steps > 0 && bh.currentPage.prev != nil {
		bh.currentPage = bh.currentPage.prev
		steps--
	}
	return bh.currentPage.url
}

func (bh *BrowserHistory) Forward(steps int) string {
	for steps > 0 && bh.currentPage.next != nil {
		bh.currentPage = bh.currentPage.next
		steps--
	}
	return bh.currentPage.url
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */