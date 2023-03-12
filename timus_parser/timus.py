import requests
from bs4 import BeautifulSoup


def timus_submission_sender(judgeid='342187EL', code='', language=65, task_id=1000):
    # This is a function, that sends solve to timus
    # and returns the id of the submission
    # then we can use that id to check the verdict

    # creating a session
    r = requests.session()

    # this is a data of a post request to timus server
    data = {
        'action': 'submit',
        'SpaceID': 1,
        'JudgeID': judgeid,
        'Language': language,
        'ProblemNum': task_id,
        'Source': code,
    }

    # sending submission to timus server with our data
    r.post('https://acm.timus.ru/submit.aspx', data=data)

    # getting a html of testing results page
    k = r.get('https://acm.timus.ru/status.aspx?space=1&count=100')

    # use beautiful soup to parse the html page with results and get the id of submission
    soup = BeautifulSoup(k.text, 'html.parser')

    # now we get all elements with html tag <tr>
    # and we have to find the latest submission from our account and get the id of it
    for link in soup.find_all('tr'):
        if link.get('class') != None and link.get('class')[0] in ['even', 'odd']:
            current_nickname = link.contents[2].text
            if current_nickname == '$tup1d2281337':
                return link.td.text


print(timus_submission_sender(code='some my solution'))
