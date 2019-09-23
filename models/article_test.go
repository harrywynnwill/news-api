package models

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func Test_SortArticleList(t *testing.T) {
	assert := assert.New(t)
	testArticleChronilogicallyFirst := &ArticleSummary{
		ID:   1,
		Date: time.Date(2017, time.Month(9), 25, 7, 22, 51, 0, time.Local),
	}
	testArticleChronilogicallySecond := &ArticleSummary{
		ID:   2,
		Date: time.Date(2018, time.Month(9), 25, 7, 22, 51, 0, time.Local),
	}
	testArticleChronilogicallyThird := &ArticleSummary{
		ID:   3,
		Date: time.Date(2019, time.Month(9), 25, 7, 22, 51, 0, time.Local),
	}

	testArticles := Articles{
		testArticleChronilogicallyFirst,
		testArticleChronilogicallyThird,
		testArticleChronilogicallySecond,
	}

	sort.Sort(testArticles)

	assert.Equal(uint(3), testArticles[0].ID, "Should be first article with ID 1")
	assert.Equal(uint(2), testArticles[1].ID, "Should not be second article with ID 2")
	assert.Equal(uint(1), testArticles[2].ID, "Should not be second article with ID 3")
}
