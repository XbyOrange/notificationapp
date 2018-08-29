# Notification App

## Architecture

![GitHub Logo](./NotificationApp.png)

## Store a new Template

This part of the API offer the possibility of create, delete and get the templates available to send notifications

### API:

http://notification-app:8080/templates

    - Method: POST
    - Json example:
    
        ```
        {
        "TemplateId": "template1",
        "Channel": "Email",
        "Lang": "ES",
        "Content": "<html><head></head><body> Hola {{.CustomData.Name}}, Gracias por contratar {{.CustomData.ItemName}}.<br><br>Pulse <a href={{.CustomData.ActivationLink}}>aquí</a> para activar de su cuenta </body></html>"
        }
        ```
http://notification-app:8080/templates/{templateID}

    - Method: GET | DELETE
    - Json response:
    
        ```
        [
            {
                "Template": "template1",
                "Content": "<html><head></head><body> Hola {{.CustomData.Name}}, Gracias por contratar {{.CustomData.ItemName}}.<br><br>Pulse <a href={{.CustomData.ActivationLink}}>aquí</a> para activar de su cuenta </body></html>"
            }
        ]
        ``` 
As you can see, the template key store in mongo will be the concatenation of:
  - TemplateType
  - Channel
  - Language
 

## Send Notification (Flow Description)
The main points will be:
  - First any microservice generate the Notification Message json:
  
  ```
  {
  "AccountId":"evento1",     
  "Subject":"Sending mail",    
  "Recipient":["albertoxxxxx@mail.es","22222222"],
  "TemplateType":"template1",
  "Channel":{
  	"Email": true
  },
  "Language": "EN",
  "CustomData":{
  	"Name":     "Peter",
  	"ItemName": "Product1",
  	"ActivationLink": "http://www.google.es"
  }
  }
  ```
In that case the translation of this message to human language should be:

### Account Id:
Represent a site id in salesforce

### Subject:
Represent the subject for the email

### Recipient:
The list of recipient to send the notification. Could be different type of recipient (email, sms, etc..) but 
in that case, the notification sent will be just the available if match with the channel.
In the example, the notification will not be sent to the telephone number as sms, because the channel is one of the list to send the notification

### TemplateType:
This is the name of the template loaded previously

### Channel:
This is the way to specify the channel or channels to use in the notification
(By default if a available channel not appear here, the value will be False. In the example, the sms channel will not be available)

### Language:
This is the language for the notification. If there are several templates available for a notification, with the language will be select the correct one.

### Custom data:
This is the value of the variables in order to be matched into the {{.CustomData.XXXX}} templates placesholder



## Build 

```
docker-compose -d up .
```

## Testing:

Need the special environment var in order to mock the mongo and email channel:
```
MOCK=MOCK
```

