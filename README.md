## Esquimo Tech Test- News App (API)

Our client wants to be able to read news articles from a provided feed that can be shown in a mobile app. They already 
have another team working on the app itself, but need help with the backend API.

For context, the mobile app has the following functionality:
- Load news articles from a public news feed

     &#x2611; *Endpoint added to load news articles from both RSS feeds with separate parsers to handle the difference in XML*
     
- Display a scrollable list of news articles

     &#x2611; *Get articles endpoint with pagination*

- Provide the option to filter news articles by category (such as UK News and Technology news), where this information 
is available
    
    &#x2611; *Get articles endpoint has query params for __category__ and __provider__.

- Show a single news article on screen, using an HTML display

    &#x2611; *Get article endpoint with richer article object that includes the content of the article

- Provide the option to share news articles via email and/or social networks
    
    &#x2612; *There is a url on the article. But i'm not sure if that is the requirement here...*

- Display a thumbnail of each article in the list of articles
    
    &#x2611; *There were no images on the items. I added the news providers logo as the image.*
    
- Present news articles in the order in which they are published

    &#x2611; *Sorting by published date done at the DB level*

- Allow the selection of different sources of news by category and provider

    &#x2611; *Get articles endpoint has query params for __category__ and __provider__.


# The Task

In terms of an API, the client wants it to be able to support the mobile app with all of the above functionality. 
They have not specified how the API should be constructed, nor have they defined any contracts. We are expected to do
this and document it accordingly. Because we don't know where the client is actually going to source their news from, 
we need to be flexible about what feeds they want to use. They have told us to use at least one of the following news feeds 
to read this data from, but they want to be able to change this at any time:

- http://feeds.bbci.co.uk/news/uk/rss.xml (BBC News UK)
- http://feeds.bbci.co.uk/news/technology/rss.xml (BBC News Technology)
- http://feeds.reuters.com/reuters/UKdomesticNews?format=xml (Reuters UK)
- http://feeds.reuters.com/reuters/technologyNews?format=xml (Reuters Technology)

### The API should have the following:
- Clear separation of data and endpoints
- Provide endpoints that are able to serve the usage patterns for the mobile app, as described above
- All endpoints must be RESTful

### You should be able to demonstrate:
- SOLID principles
- Ability to create simple, meaningful contracts/interfaces for each appfunction
- Secure practices, good sanity checking and stability
- Ability to store data, such as the news feed locations
- Bonus: Adopt a microservices architecture to provide resilience and scalability
- Bonus: Use a third-party API provider to leverage any functionality
- Bonus: Provide caching in the API to allow for faster response times

# Pro Tips for Success
1.Build the API in an environment you are comfortable with, such as
Visual Studio, PyCharm, IntelliJ, vim, emacs or what else suits you.
2.Language doesn't matter, just as long as it is readable, clear and
concise.
3.Feel free to use 3rd-party frameworks or libraries as you see fit. Of
course we will ask why you made a choice to use/not use it.
4.Make sure it builds and runs in a clean environment. If you can
provide a working example, all the better.
5.Create a clean public repo at github, push your code there and give us
the link, we’ll checkout from there.
6.As this task is purely for APIs, you don’t need to build a client.
Documentation about how to use it helps, though.
7.Do not spend more than a few hours on the project. If you are unable
to complete a feature, don't worry, but at least show your working.
8.Be ready to explain your code! We are looking for robust, clean code.
If you need to hack anything in, be sure you're ready to explain why.
9.Don't forget to check for null references!
Good luck!

# Prerequisites 
Linux - tested on Ubuntu 18 (might work on mac)
Go version 1.12 (might run on 1.11!) 
Docker

# Installation

1. Install and start the DB. It is a mysql docker container.

    * `chmod +x start.db`
    * `./start.db`

    ... Wait until its up

2. Seed the DB with news sources.

    * `chmod +x seed-db.sh`
    * `./seed-db.sh`

3. Compile and run the app.

    * `go build main.go` - compile
    * `./main` - run the app!


# Endpoints
  * GET `http://localhost:8000/news?offset=0&pageSize=10&provider=reuters&category=technologyNews`
    
    News summaries in a list with pagination
  
  ### Response Object
  ```
{
    "articles": [
        {
            "id": 1,
            "title": "SoftBank mulls bringing 40 companies to Brazil",
            "category": "technologyNews",
            "urlToImage": "https://www.reuters.com/resources_v2/images/reuters125.png",
            "date": "2019-09-21T21:00:34+01:00",
            "provider": "reuters"
        },
        {
            "id": 2,
            "title": "Huawei to join forces with China Mobile to bid for Brazil's Oi: report",
            "category": "technologyNews",
            "urlToImage": "https://www.reuters.com/resources_v2/images/reuters125.png",
            "date": "2019-09-21T15:25:20+01:00",
            "provider": "reuters"
        },
        {
            "id": 3,
            "title": "U.S. trade regulators approve some Apple tariff exemptions amid broader reprieve",
            "category": "technologyNews",
            "urlToImage": "https://www.reuters.com/resources_v2/images/reuters125.png",
            "date": "2019-09-21T02:42:59+01:00",
            "provider": "reuters"
        },
        {
            "id": 4,
            "title": "Facebook suspends tens of thousands of apps in response to Cambridge Analytica row",
            "category": "technologyNews",
            "urlToImage": "https://www.reuters.com/resources_v2/images/reuters125.png",
            "date": "2019-09-21T02:08:17+01:00",
            "provider": "reuters"
        },
        {
            "id": 5,
            "title": "Few U.S. lawmakers hit 'like' button after Facebook CEO visits Capitol Hill",
            "category": "technologyNews",
            "urlToImage": "https://www.reuters.com/resources_v2/images/reuters125.png",
            "date": "2019-09-20T22:21:53+01:00",
            "provider": "reuters"
        }
    ],
    "pageSize": 5,
    "offSet": 0,
    "totalRecords": 91
}
```
  ### Request Parameters
  
  __pageSize__ number of articles in the response
  __offSet__ number of articles to offset
  __category__ query by category
  __provider__ query by news provider
  
  
  * GET http://localhost:8000/news/1`
   
   Get the article by ID. Returns a richer article model for viewing per article in HTML.
   
```
{
    "id": 1,
    "title": "SoftBank mulls bringing 40 companies to Brazil",
    "category": "technologyNews",
    "urlToImage": "https://www.reuters.com/resources_v2/images/reuters125.png",
    "date": "2019-09-21T21:00:34+01:00",
    "provider": "reuters",
    "url": "http://feeds.reuters.com/~r/reuters/technologyNews/~3/Vlo31M2hAvo/softbank-mulls-bringing-40-companies-to-brazil-idUSKBN1W52G2",
    "description": "Japan's SoftBank Group Corp is considering bringing around 40 companies with high growth potential to Brazil and expects to announce a large investment in the country in around two weeks' time, the group's head in Brazil, André Maciel, said on Friday.<div class=\"feedflare\">\n<a href=\"http://feeds.reuters.com/~ff/reuters/technologyNews?a=Vlo31M2hAvo:CfWsZZzm3N4:yIl2AUoC8zA\"><img src=\"http://feeds.feedburner.com/~ff/reuters/technologyNews?d=yIl2AUoC8zA\" border=\"0\"></img></a> <a href=\"http://feeds.reuters.com/~ff/reuters/technologyNews?a=Vlo31M2hAvo:CfWsZZzm3N4:V_sGLiPBpWU\"><img src=\"http://feeds.feedburner.com/~ff/reuters/technologyNews?i=Vlo31M2hAvo:CfWsZZzm3N4:V_sGLiPBpWU\" border=\"0\"></img></a> <a href=\"http://feeds.reuters.com/~ff/reuters/technologyNews?a=Vlo31M2hAvo:CfWsZZzm3N4:-BTjWOF_DHI\"><img src=\"http://feeds.feedburner.com/~ff/reuters/technologyNews?i=Vlo31M2hAvo:CfWsZZzm3N4:-BTjWOF_DHI\" border=\"0\"></img></a>\n</div><img src=\"http://feeds.feedburner.com/~r/reuters/technologyNews/~4/Vlo31M2hAvo\" height=\"1\" width=\"1\" alt=\"\"/>"
}
```
    
    
  * POST `http://localhost:8000/news`
  
  This Would not be part of the public API...
 
  At the moment it requests the news from the stored sources and loads the articles into the DB.
    


