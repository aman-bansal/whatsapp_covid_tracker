## WhatsApp Covid Tracker

---
This project has no external dependencies. Only one route is exposed which is added for webhook over twilio console.

#### Requirements:

---
1. Twilio Account is needed. Enable whatsapp notification and add webhook url of your deployment in twilio console.
2. For top 10 news related to covid for a country, you need to login at https://api.smartable.ai . 
Enable free api subscription and subscription as env variable with key as ```SMARTABLE_AI_SUBS_KEY```
3. Country codes are standard ISO2 codes which are two character long . You can find your country code in the response of API https://api.covid19api.com/countries .


#### Valid whatsapp codes

---
```CASES TOTAL``` : Total number of cases globally

```CASES NEW``` : New cases globally

```CASES <COUNTRY CODE IN CAPS>``` : Total cases in a given country

```DEATHS TOTAL``` : Total number of deaths globally

```DEATHS NEW```: New deaths reported globally

```DEATHS <COUNTRY CODE IN CAPS>```: Total deaths reported in a given country

```NEWS GLOBAL```: Top 10 global news

```NEWS <COUNTRY CODE IN CAPS>```: Top 10 news for a given country

#### Code Structure

---

`pkg` folder is the main folder with multiple directories and main.go as the primary entry point.

`pkg/constant` contains all standard message strings format.

`pkg/controller` handles the logic to control the behaviour of webhook url.

`pkg/data_service` handles retrieval of data from external sources.

`pkg/job` handles continous sync of covid data from source. simple job which sleeps for 60 sec to refresh data.

`pkg/model` contains basic objects struct that is needed to data serialization/deserialization.

`pkg/repository` basic in memory based repository to handle temporary data storage.

`pkg/use_case` handles all the valid use cases the product has.