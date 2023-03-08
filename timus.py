import requests
from bs4 import BeautifulSoup


def timus_solve_sender(judgeid='342187EL', code='', language=65, task_id=1000):
    r = requests.session()

    data = {
        'action': 'submit',
        'SpaceID': 1,
        'JudgeID': judgeid,
        'Language': language,
        'ProblemNum': task_id,
        'Source': code,
    }

    # sending a solve
    r.post('https://acm.timus.ru/submit.aspx', data=data)

    # getting a html of results page
    k = r.get('https://acm.timus.ru/status.aspx?space=1&count=100')

    # use beautiful soup to parse the html page with results and get the id of the solve
    soup = BeautifulSoup(k.text, 'html.parser')
    for link in soup.find_all('tr'):
        if link.get('class') != None and link.get('class')[0] in ['even', 'odd']:
            current_nickname = link.contents[2].text
            if current_nickname == '$tup1d2281337':
                return link.td.text


print(timus_solve_sender(code='some my solution'))
