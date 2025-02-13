package feed

import (
	"fmt"
	"proj1/lock"
	"strings"
)

// Feed represents a user's twitter feed
// You will add to this interface the implementations as you complete them.
type Feed interface {
	Add(body string, timestamp float64)
	Remove(timestamp float64) bool
	Contains(timestamp float64) bool
	Display() string
	ContentsAsMap() *map[float64]string
}

// feed is the internal representation of a user's twitter feed (hidden from outside packages)
// You CAN add to this structure but you cannot remove any of the original fields. You must use
// the original fields in your implementation. You can assume the feed will not have duplicate posts
type feed struct {
	rw    *lock.RWLock // an interface defined for RWMutex
	start *post        // a pointer to the beginning post
}

// post is the internal representation of a post on a user's twitter feed (hidden from outside packages)
// You CAN add to this structure but you cannot remove any of the original fields. You must use
// the original fields in your implementation.
type post struct {
	body      string  // the text of the post
	timestamp float64 // Unix timestamp of the post
	next      *post   // the next post in the feed
}

// NewPost creates and returns a new post value given its body and timestamp
func newPost(body string, timestamp float64, next *post) *post {
	return &post{body, timestamp, next}
}

// NewFeed creates a empty user feed
func NewFeed() Feed {
	rw := lock.NewRWLock()
	return &feed{rw, nil}
}

// Add inserts a new post to the feed. The feed is always ordered by the timestamp where
// the most recent timestamp is at the beginning of the feed followed by the second most
// recent timestamp, etc. You may need to insert a new post somewhere in the feed because
// the given timestamp may not be the most recent.
func (f *feed) Add(body string, timestamp float64) {

	f.rw.Lock()

	// base case: empty LL
	if f.start == nil {
		f.start = newPost(body, timestamp, nil)
		f.rw.Unlock()
		return
	}

	// else: traverse LL until hit end or find item with 'larger' timestamp
	// while curr is not nil and curr.timestamp < timestamp, advance
	var pred *post = nil
	var curr *post = f.start
	for curr != nil && curr.timestamp < timestamp {
		pred = curr
		curr = curr.next
	}

	// case: curr is nil; reached end of LL. insert at end of LL
	if curr == nil {
		newPost := newPost(body, timestamp, nil)
		pred.next = newPost
		f.rw.Unlock()
		return
	}

	// case: curr.timestamp > timestamp. insert before curr.
	if curr.timestamp >= timestamp {
		newPost := newPost(body, timestamp, curr)
		// case: curr is head of LL (update head of LL)
		if pred == nil {
			f.start = newPost
		} else { // case: curr is not head of LL (need to update pred)
			pred.next = newPost
		}
		f.rw.Unlock()
		return
	}

	f.rw.Unlock()
	fmt.Println("ERROR: did not correctly perform Add() operation")
}

// Remove deletes the post with the given timestamp. If the timestamp
// is not included in a post of the feed then the feed remains
// unchanged. Return true if the deletion was a success, otherwise return false
func (f *feed) Remove(timestamp float64) bool {

	f.rw.Lock()

	// base case: empty LL; trivial false
	if f.start == nil {
		f.rw.Unlock()
		return false
	}

	// else: traverse LL until hit end or find item
	// while curr is not nil and curr.timestamp != timestamp, advance
	var pred *post = nil
	var curr *post = f.start
	for curr != nil && curr.timestamp != timestamp {
		pred = curr
		curr = curr.next
	}

	// case: curr is nil; reached end of LL. item not present
	if curr == nil {
		f.rw.Unlock()
		return false
	}

	// case: curr.timestamp == timestamp. remove item.
	if curr.timestamp == timestamp {
		// case: curr is last item in LL
		if curr.next == nil {
			// case: curr is also first item in LL
			if pred == nil {
				f.start = nil // curr is only item in LL. head of LL is now nil
			} else { // case: curr is not first item in LL
				pred.next = nil
			}
		} else { // case: curr is not last item in LL
			// case: curr is first item in LL
			if pred == nil {
				f.start = curr.next
			} else { // case: curr is not first item in LL
				next := curr.next
				pred.next = next
			}
		}
		f.rw.Unlock()
		return true
	}

	f.rw.Unlock()
	fmt.Println("ERROR: did not correctly perform Remove() operation!")
	return false
}

// Contains determines whether a post with the given timestamp is
// inside a feed. The function returns true if there is a post
// with the timestamp, otherwise, false.
func (f *feed) Contains(timestamp float64) bool {

	f.rw.RLock()
	// base case: empty LL; trivial false
	if f.start == nil {
		f.rw.RUnlock()
		return false
	}

	// else: traverse LL until hit end or find item
	// while curr is not nil and curr.timestamp != timestamp, advance
	var curr *post = f.start
	for curr != nil && curr.timestamp != timestamp {
		curr = curr.next
	}

	// case: curr is nil; reached end of LL. item not present
	if curr == nil {
		f.rw.RUnlock()
		return false
	}

	// case: curr.timestamp == timestamp. return true.
	if curr.timestamp == timestamp {
		f.rw.RUnlock()
		return true
	}

	f.rw.RUnlock()
	fmt.Println("ERROR: did not correctly perform Contains() operation!")
	return false
}

// Display will return a pretty string representation of this feed
func (f *feed) Display() string {

	f.rw.RLock()
	var items []post = make([]post, 0)

	// base case: empty LL; trivial false
	if f.start == nil {
		f.rw.RUnlock()
		return "[ ]"
	}

	// else: traverse LL until hit end or find item
	// while curr is not nil and curr.timestamp != timestamp, advance
	var pred *post = f.start
	var curr *post = pred.next

	items = append(items, *pred) // add first item

	for curr != nil {
		items = append(items, *curr) // add curr
		pred = curr
		curr = curr.next
	}

	var strPosts []string = make([]string, len(items))
	for idx, item := range items {
		strPosts[idx] = fmt.Sprintf("%v", item)
	}

	f.rw.RUnlock()
	return "[ " + strings.Join(strPosts, " -> ") + " ]"
}

// ContentsToMap will return the contents of this feed as a map
func (f *feed) ContentsAsMap() *map[float64]string {

	f.rw.RLock()
	contents := map[float64]string{}

	// base case: empty LL
	if f.start == nil {
		return &contents
	}

	// else: traverse LL until hit end or find item
	// while curr is not nil and curr.timestamp != timestamp, advance
	var pred *post = f.start
	var curr *post = pred.next

	contents[pred.timestamp] = pred.body // add first item

	for curr != nil {
		contents[curr.timestamp] = curr.body // add first item // add curr
		pred = curr
		curr = curr.next
	}

	return &contents
}

//// INITIAL IMPLEMENTATION FOR PART 1 BELOW ////
// // feed is the internal representation of a user's twitter feed (hidden from outside packages)
// // You CAN add to this structure but you cannot remove any of the original fields. You must use
// // the original fields in your implementation. You can assume the feed will not have duplicate posts
// type feed struct {
// 	start *post // a pointer to the beginning post
// }

// // post is the internal representation of a post on a user's twitter feed (hidden from outside packages)
// // You CAN add to this structure but you cannot remove any of the original fields. You must use
// // the original fields in your implementation.
// type post struct {
// 	body      string  // the text of the post
// 	timestamp float64 // Unix timestamp of the post
// 	next      *post   // the next post in the feed
// }

// // NewPost creates and returns a new post value given its body and timestamp
// func newPost(body string, timestamp float64, next *post) *post {
// 	return &post{body, timestamp, next}
// }

// // NewFeed creates a empty user feed
// func NewFeed() Feed {
// 	return &feed{start: nil}
// }

// // Add inserts a new post to the feed. The feed is always ordered by the timestamp where
// // the most recent timestamp is at the beginning of the feed followed by the second most
// // recent timestamp, etc. You may need to insert a new post somewhere in the feed because
// // the given timestamp may not be the most recent.
// func (f *feed) Add(body string, timestamp float64) {

// 	// base case: empty LL
// 	if f.start == nil {
// 		f.start = newPost(body, timestamp, nil)
// 		return
// 	}

// 	// else: traverse LL until hit end or find item with 'larger' timestamp
// 	// while curr is not nil and curr.timestamp < timestamp, advance
// 	var pred *post = nil
// 	var curr *post = f.start
// 	for curr != nil && curr.timestamp < timestamp {
// 		pred = curr
// 		curr = curr.next
// 	}

// 	// case: curr is nil; reached end of LL. insert at end of LL
// 	if curr == nil {
// 		newPost := newPost(body, timestamp, nil)
// 		pred.next = newPost
// 		return
// 	}

// 	// case: curr.timestamp > timestamp. insert before curr.
// 	if curr.timestamp >= timestamp {
// 		newPost := newPost(body, timestamp, curr)
// 		// case: curr is head of LL (update head of LL)
// 		if pred == nil {
// 			f.start = newPost
// 		} else { // case: curr is not head of LL (need to update pred)
// 			pred.next = newPost
// 		}
// 		return
// 	}

// 	fmt.Println("ERROR: did not correctly perform Add() operation")
// }

// // Remove deletes the post with the given timestamp. If the timestamp
// // is not included in a post of the feed then the feed remains
// // unchanged. Return true if the deletion was a success, otherwise return false
// func (f *feed) Remove(timestamp float64) bool {

// 	// base case: empty LL; trivial false
// 	if f.start == nil {
// 		return false
// 	}

// 	// else: traverse LL until hit end or find item
// 	// while curr is not nil and curr.timestamp != timestamp, advance
// 	var pred *post = nil
// 	var curr *post = f.start
// 	for curr != nil && curr.timestamp != timestamp {
// 		pred = curr
// 		curr = curr.next
// 	}

// 	// case: curr is nil; reached end of LL. item not present
// 	if curr == nil {
// 		return false
// 	}

// 	// case: curr.timestamp == timestamp. remove item.
// 	if curr.timestamp == timestamp {
// 		// case: curr is last item in LL
// 		if curr.next == nil {
// 			// case: curr is also first item in LL
// 			if pred == nil {
// 				f.start = nil // curr is only item in LL. head of LL is now nil
// 			} else { // case: curr is not first item in LL
// 				pred.next = nil
// 			}
// 		} else { // case: curr is not last item in LL
// 			// case: curr is first item in LL
// 			if pred == nil {
// 				f.start = curr.next
// 			} else { // case: curr is not first item in LL
// 				next := curr.next
// 				pred.next = next
// 			}
// 		}
// 		return true
// 	}

// 	fmt.Println("ERROR: did not correctly perform Remove() operation!")
// 	return false
// }

// // Contains determines whether a post with the given timestamp is
// // inside a feed. The function returns true if there is a post
// // with the timestamp, otherwise, false.
// func (f *feed) Contains(timestamp float64) bool {

// 	// base case: empty LL; trivial false
// 	if f.start == nil {
// 		return false
// 	}

// 	// else: traverse LL until hit end or find item
// 	// while curr is not nil and curr.timestamp != timestamp, advance
// 	var curr *post = f.start
// 	for curr != nil && curr.timestamp != timestamp {
// 		curr = curr.next
// 	}

// 	// case: curr is nil; reached end of LL. item not present
// 	if curr == nil {
// 		return false
// 	}

// 	// case: curr.timestamp == timestamp. return true.
// 	if curr.timestamp == timestamp {
// 		return true
// 	}

// 	fmt.Println("ERROR: did not correctly perform Contains() operation!")
// 	return false
// }

// // Display will return a pretty string representation of this feed
// func (f *feed) Display() string {

// 	var items []post = make([]post, 0)

// 	// base case: empty LL; trivial false
// 	if f.start == nil {
// 		return "[ ]"
// 	}

// 	// else: traverse LL until hit end or find item
// 	// while curr is not nil and curr.timestamp != timestamp, advance
// 	var pred *post = f.start
// 	var curr *post = pred.next

// 	items = append(items, *pred) // add first item

// 	for curr != nil {
// 		items = append(items, *curr) // add curr
// 		pred = curr
// 		curr = curr.next
// 	}

// 	var strPosts []string = make([]string, len(items))
// 	for idx, item := range items {
// 		strPosts[idx] = fmt.Sprintf("%v", item)
// 	}

// 	return "[ " + strings.Join(strPosts, " -> ") + " ]"
// }
