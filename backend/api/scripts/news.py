from bs4 import BeautifulSoup
import requests
import os


def main():
    get_news()


def get_news():
    url = 'https://www.bth.se/category/nyheter/'
    r = requests.get(url, allow_redirects=True)
    open('news.html', 'wb').write(r.content)

    with open('news.html', 'r') as html_file:
        html_content = html_file.read()

        soup = BeautifulSoup(html_content, 'lxml')
        news = soup.find_all('article', class_='ArticleItem')
        for content in news:
            title = content.find('h2').text
            date = content.find('p').text
            text = content.find('div', class_='article-category-page').text
            # I WILL ADD LINKS AND IMAGES LATER
    os.remove('news.html')



if __name__ == '__main__':
    main()