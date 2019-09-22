package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func Test_reutersParse(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	xmlTestNews, err := os.Open("reuters-test-news.xml")
	require.Nil(err)

	byteTestNews, err := ioutil.ReadAll(xmlTestNews)
	require.Nil(err)

	articles, err := ReutersParseService.Parse(byteTestNews)
	require.Nil(err)

	defer xmlTestNews.Close()

	assert.Equal("Facebook CEO Zuckerberg dines with senators in DC", articles[0].Title)
	assert.Equal("technologyNews", articles[0].Category)
	assert.Equal("http://feeds.reuters.com/~r/reuters/technologyNews/~3/2jzt4dJ98YM/facebook-ceo-zuckerberg-dines-with-senators-in-dc-idUSKBN1W32SW", articles[0].Url)
	minus4 := time.FixedZone("", -60*60*4)
	assert.Equal(time.Date(2019, time.September, 19, 00, 06, 03, 0, minus4), articles[0].Date, "Date has failed to parser")
	assert.Equal(20, len(articles))
}
