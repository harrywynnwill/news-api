# Esquimo Tech Test- News App (API)

## Prerequisites 
Linux - tested on Ubuntu 18 (might work on mac)  
Go version 1.12 (might run on 1.11 - requires go mod)   
Docker (I am running v. 19.03.2)  

## Installation

1. Install and start the DB. It is a mysql docker container.

    * `chmod +x start.db`
    * `./start.db`

    ... Wait until its up

2. Compile and run the app.

    * `go build main.go` - compile
    * `./main` - run the app! (will run the migrations on start up)
    
3. Seed the DB with news sources.

    * `chmod +x seed-db.sh`
    * `./seed-db.sh`

## Design

Designed in a micro-services (within monolith) pattern, As we scale, we can break the app up and scale it horizontally.
The services and repos are programmed to interface and implemented as singletons. Down the line we can generate clients for our 
extracted services and easily break the app up.

I focused on delivering value. I have taken a pragmatic approach to testing in order to deliver the API. We could do 
with some integration tests and probably more unit tests. Any logic that is not simple object mapping I have unit tested.

I would like to have added a cache layer but ran out-of-time. I would have implemented the cache on the GET news/:id route.

There is also no security on the app.

## Spec

_inline comments in italics_

Our client wants to be able to read news articles from a provided feed that can be shown in a mobile app. They already 
have another team working on the app itself, but need help with the backend API.

For context, the mobile to achieveapp has the following functionality:
- Load news articles from a public news feed

     &#x2611; *Endpoint added to load news articles from both RSS feeds, separate parsers to handle the difference in XML*
     
- Display a scrollable list of news articles

     &#x2611; *Get articles endpoint with pagination*

- Provide the option to filter news articles by category (such as UK News and Technology news), where this information 
is available
    
    &#x2611; *Get _/articles_ endpoint has query params for __category__ and __provider__.

- Show a single news article on screen, using an HTML display

    &#x2611; *Get _/article_ by ID endpoint with richer article object that includes the content of the article*

- Provide the option to share news articles via email and/or social networks
    
     &#x2610; *There is a url on the article object and _GET article_ endpoint. 
     But I'm not sure if that is the requirement here...*

- Display a thumbnail of each article in the list of articles
    
    &#x2611; *There were no images on the items. I added the news providers logo as the image.*
    
- Present news articles in the order in which they are published

    &#x2611; *Sorting by published date done at the DB level. I have added an index to the published date* 

- Allow the selection of different sources of news by category and provider
    
    &#x2611; *Get articles endpoint has query params for __category__ and __provider__. In the interests of delivering the 
    API I have opted for the DB solution. The ideal solution might be to store the data and use Elastic Search which is great
    for searching data.*

## The Task

In terms of an API, the client wants it to be able to support the mobile app with all of the above functionality. 
They have not specified how the API should be constructed, nor have they defined any contracts. We are expected to do
this and document it accordingly. 
Because we don't know where the client is actually going to source their news from, we need to be flexible about what 
feeds they want to use. They have told us to use at least one of the following news feeds to read this data from, but
 they want to be able to change this at any time:

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

## Endpoints
   ### GET /news
   `http://localhost:8000/news?offset=0&pageSize=10&provider=reuters&category=technologyNews`  
   
   News summaries in a list with pagination.
  
  #### Response Object
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
  #### Request Parameters
  
  __pageSize__ number of articles in the response (maximum 300 articles)
  __offSet__ number of articles to offset (maximum 300 articles)
  __category__ query by category 
  __provider__ query by news provider
  
  
  ### GET news/:id
  
  Get the article by ID. Returns a richer article model for viewing per article in HTML.
  
  #### Response Object
   
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
    
    
  ## POST /news
  
  Request the news from 3rd party sources and load the articles into the DB.
  __This Would not be part of the public API...__
    




## Spec

_inline comments in italics_

Our client wants to be able to read news articles from a provided feed that can be shown in a mobile app. They already 
have another team working on the app itself, but need help with the backend API.

For context, the mobile to achieveapp has the following functionality:
- Load news articles from a public news feed

     &#x2611; *Endpoint added to load news articles from both RSS feeds, separate parsers to handle the difference in XML*
     
- Display a scrollable list of news articles

     &#x2611; *Get articles endpoint with pagination*

- Provide the option to filter news articles by category (such as UK News and Technology news), where this information 
is available
    
    &#x2611; *Get _/articles_ endpoint has query params for __category__ and __provider__.

- Show a single news article on screen, using an HTML display

    &#x2611; *Get _/article_ by ID endpoint with richer article object that includes the content of the article*

- Provide the option to share news articles via email and/or social networks
    
     &#x2611; *There is a url on the article object and _GET article_ endpoint. 
     But I'm not sure if that is the requirement here...*

- Display a thumbnail of each article in the list of articles
    
    &#x2611; *There were no images on the items. I added the news providers logo as the image.*
    
- Present news articles in the order in which they are published

    &#x2611; *Sorting by published date done at the DB level. I have added an index to the published date* 

- Allow the selection of different sources of news by category and provider
    
    &#x2611; *Get articles endpoint has query params for __category__ and __provider__. In the interests of delivering the 
    API I have opted for the DB solution. The ideal solution might be to store the data and use Elastic Search which is great
    for searching data.*

## The Task

In terms of an API, the client wants it to be able to support the mobile app with all of the above functionality. 
They have not specified how the API should be constructed, nor have they defined any contracts. We are expected to do
this and document it accordingly. 
Because we don't know where the client is actually going to source their news from, we need to be flexible about what 
feeds they want to use. They have told us to use at least one of the following news feeds to read this data from, but
 they want to be able to change this at any time:

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
