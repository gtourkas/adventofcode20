package adventofcode20

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Rule struct {
	Color      string
	CanContain map[string]int
}

type Node struct {
	Color    string
	Children []*NodeLink
	IsRoot   bool
}

type NodeLink struct {
	Number int
	Child  *Node
}

func buildNodesSet(rules []Rule) map[string]*Node {
	nodes := make(map[string]*Node)
	for _, r := range rules {
		var ok bool

		var n *Node
		if n, ok = nodes[r.Color]; !ok {
			n = &Node{
				Color:    r.Color,
				Children: make([]*NodeLink, 0),
				IsRoot:   true,
			}
			nodes[n.Color] = n
		}

		for c, num := range r.CanContain {
			var childNode *Node
			if childNode, ok = nodes[c]; !ok {
				childNode = &Node{
					Color:    c,
					Children: make([]*NodeLink, 0),
					IsRoot:   false,
				}
				nodes[c] = childNode
			}
			n.Children = append(n.Children, &NodeLink{
				Number: num,
				Child:  childNode,
			})
		}
	}
	return nodes
}
func walkToFind(n *Node, findColor string, path []string, found map[string]bool) {
	if n.Color == findColor {
		for _, p := range path {
			if p != findColor {
				found[p] = true
			}
		}
		return
	}
	for _, c := range n.Children {
		subPath := append(path, c.Child.Color)
		walkToFind(c.Child, findColor, subPath, found)
	}
}

func HandyHaversacksPartA(rules []Rule) int {

	nodes := buildNodesSet(rules)

	const FindColor = "shiny gold"
	found := map[string]bool{}
	for _, n := range nodes {
		if n.IsRoot && n.Color != FindColor {
			walkToFind(n, FindColor, []string{n.Color}, found)
		}
	}

	return len(found)
}

func walkToCount(n *Node, multiplier int) int {
	r := 0
	for _, c := range n.Children {
		total := multiplier * c.Number
		r += total
		r += walkToCount(c.Child, total)
	}
	return r
}

func HandyHaversacksPartB(rules []Rule) int {

	nodes := buildNodesSet(rules)

	n := nodes["shiny gold"]
	r := walkToCount(n, 1)

	return r
}

func RunHandyHaversacks(fileName string, f func([]Rule) int) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	outerRe := regexp.MustCompile("^(?P<c>[a-z]+\\s[a-z]+)\\sbags\\scontain\\s")
	innerRe := regexp.MustCompile("(((?P<n>\\d)\\s(?P<c>[a-z]+\\s[a-z]+)\\sbags*)+)")

	rules := make([]Rule, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		r := Rule{
			Color:      "",
			CanContain: map[string]int{},
		}

		matchOuter := outerRe.FindStringSubmatch(line)
		r.Color = matchOuter[1]
		for _, matchInner := range innerRe.FindAllStringSubmatch(line, -1) {
			num, _ := strconv.Atoi(matchInner[3])
			color := matchInner[4]
			r.CanContain[color] = num
		}

		rules = append(rules, r)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return f(rules)
}
