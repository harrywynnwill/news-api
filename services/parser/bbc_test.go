package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func Test_bbcParse(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	xmlTestNews, err := os.Open("bbc-test-news.xml")
	require.Nil(err)

	byteTestNews, err := ioutil.ReadAll(xmlTestNews)
	require.Nil(err)

	articles, err := BBCParseService.Parse(byteTestNews)
	require.Nil(err)

	defer xmlTestNews.Close()

	assert.Equal("What3words: 'Life-saving app' divides opinion", articles[0].Title, "Title failed to parse")
	assert.Equal("https://www.bbc.co.uk/news/technology-49754820", articles[0].Url, "Category failed to parse")
	assert.Equal(time.Date(2019, time.September, 21, 00, 53, 41, 0, time.Local), articles[0].Date, "Date has failed to parser")
	assert.Equal(22, len(articles))
}