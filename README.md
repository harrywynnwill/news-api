## Esquimo Tech Test

News App (API)
Our client wants to be able to read news articles from a provided feed that
can be shown in a mobile app. They already have another team working on the
app itself, but need help with the backend API.
For context, the mobile app has the following functionality:
- Load news articles from a public news feed
- Display a scrollable list of news articles
- Provide the option to filter news articles by category (such as UK
News and Technology news), where this information is available
- Show a single news article on screen, using an HTML display
- Provide the option to share news articles via email and/or social
networks
- Display a thumbnail of each article in the list of articles
- Present news articles in the order in which they are published
- Allow the selection of different sources of news by category and
provider
The Task
In terms of an API, the client wants it to be able to support the mobile
app with all of the above functionality. They have not specified how the
API should be constructed, nor have they defined any contracts. We are
expected to do this and document it accordingly.
Because we don't know where the client is actually going to source their
news from, we need to be flexible about what feeds they want to use. They
have told us to use at least one of the following news feeds to read this
data from, but they want to be able to change this at any time:
- http://feeds.bbci.co.uk/news/uk/rss.xml (BBC News UK)
- http://feeds.bbci.co.uk/news/technology/rss.xml (BBC News Technology)
- http://feeds.reuters.com/reuters/UKdomesticNews?format=xml (Reuters
UK)
- http://feeds.reuters.com/reuters/technologyNews?format=xml (Reuters
Technology)
The API should have the following:
- Clear separation of data and endpoints
- Provide endpoints that are able to serve the usage patterns for the
mobile app, as described above
1- All endpoints must be RESTful
You should be able to demonstrate:
- SOLID principles
- Ability to create simple, meaningful contracts/interfaces for each app
function
- Secure practices, good sanity checking and stability
- Ability to store data, such as the news feed locations
- Bonus: Adopt a microservices architecture to provide resilience and
scalability
- Bonus: Use a third-party API provider to leverage any functionality
- Bonus: Provide caching in the API to allow for faster response times
Pro Tips for Success
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


model : {
    articles :[
      "category: "technology3"
     "author": "Zack Beauchamp",
     "title": "Israel election: how it works, who might win, and why it matters",
     "description": "Israeli Prime Minister Benjamin Netanyahu’s desperate reelection bid, explained.",
     "url": "https://www.vox.com/world/2019/9/17/20869050/israel-election-results-netanyahu-jordan-valley-democracy",
     "urlToImage": "https://cdn.vox-cdn.com/thumbor/Mx_E-moZUFFua6DLSaoWoHg0skQ=/0x161:5290x2931/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19207260/1167234615.jpg.jpg",
     "publishedAt": "2019-09-17T13:40:00+00:00",
     "content": "Israeli Prime Minister Benjamin Netanyahu is standing for reelection today in one of the strangest and potentially most significant elections in the countrys modern history.\r\nThis is actually the second national election Israels had in just six months. The la… [+15433 chars]"
     },
    ]
}