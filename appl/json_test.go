package appl

import (
	"fmt"
	"testing"
)

func TestBeautify(t *testing.T) {
	input := []string{
		"&amp;&lt;",
		"S&amp;L's",
		"&amp;amp;",
		"&amp;gt;",
		"&lt;&#228;&amp;gt;",
		"S&amp;#xe3;n Paulo",
		" &lt;a&amp;gt; \n \r\n &amp;amp;",
		" &#160; &lt;&#xe3;&amp;gt; \n &amp;lt;a href=\"#\"&gt;test</a>",
		" &#160; &lt;&#228;&amp;gt; \n &amp;&lt;a href=\"#\"&gt;test</a>",
		"&amp;#x3C;&amp;#x3E;&amp;#x201C;",
		"&#1 &#a; &#xs;",
	}
	expected := []string{
		"&<",
		"S&L's",
		"&",
		">",
		"<ä>",
		"Sãn Paulo",
		"<a> &",
		"<ã> <a href=\"#\">test</a>",
		"<ä> &<a href=\"#\">test</a>",
		"<>“",
		"&#1 &#a; &#xs;",
	}
	for i, s := range input {
		test := beautify(s)
		fmt.Printf("%d %s\n", i, test)
		if expected[i] != test {
			t.Error(fmt.Printf("Expected: %s, got %s\n", expected[i], test))
		}
	}
}
