import codecs
from urllib import response
import csv
import json
import requests
import os


res = requests.get('https://docs.google.com/spreadsheets/d/1hc_ckooHG4sEygqtoefxwQV2F6TwtfEdGGSbr5y0CJs/export?format=csv') # get the data from google sheets
res.encoding = "utf-8"

res.close()

f = open('src/locale/translation.csv', 'w', encoding='utf-8') # open the file to write
f.write(res.text)
f.close()

csvFilePath = 'src/locale/translation.csv'
jsonFilePath = 'src/locale/translation.json'

csvData = ''

data = {}

with open(csvFilePath, 'r') as csvFile:
    #print out contents of csv file
    csvReader = csv.DictReader(csvFile)
    for row in csvReader:
      # id = row['id']
      string = row['string']
      data[string] = {
        'en' : row['en'],
        'sv-sv' : row['sv-sv'],
        'ar' : row['ar'],
      }

with codecs.open(jsonFilePath, 'w', 'utf-8') as jsonFile:
    jsonFile.write(json.dumps(data, ensure_ascii=False, indent=4))
    jsonFile.close()


# remove the csv file
os.remove('src/locale/translation.csv')