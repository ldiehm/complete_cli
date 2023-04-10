# graphQLMetric('https://github.com/cloudinary/cloudinary_npm.git')

import requests
import sys

import os

headers = {"Authorization":  "token " + os.getenv("GITHUB_TOKEN")}

def graphQLMetric(url):

    variables = {
        "url": url
      }
        
    query ="""
    query MyQuery($url: URI!) {
    resource(url: $url) {
        ... on Repository {
            pullRequests(last:100, states: [CLOSED, MERGED]) {
            nodes {
            number
            url
            state
            reviews(last:10) {
                nodes{
                author {
                    login
                }
                state
                }
            }
            }
        } 
        }
    }
    }
    """

    result = performQuery(query, variables)
    #print('result: ')
    #print(result)
    file1 = open('QueryOutput.txt', 'w')
    file1.write(str(result))
    file1.close()
    if result == 0:
        return 0

    return 1

def performQuery(query, variables):
    request = requests.post('https://api.github.com/graphql', json={'query': query, 'variables': variables}, headers=headers)
    print(request.reason)
    if request.status_code == 200: # checks if request was successful
        return request.json() # returns results
    else: # in case of expection
      return 0 # error case
    
graphQLMetric(sys.argv[1])