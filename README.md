# MurgeMachine-API
API for the Murge Machine project

https://docs.google.com/drawings/d/1gY9Sr3afUbbaoN3JzfTktn9-hpBiD_OB6CGR5jjYZoM/edit?usp=sharing

## Endpoints

/drinks    -> get drinks  
/cocktails -> get cocktails  
/pumps -> get pumps

## JSON TEMPLATES
### JSON sent via MQTT to activate the pumps format : 

size -> 1 for 25 cl, 2 for 50cl

```json
{
  "preparation" : {
    "size" : 1,
    "pumpsActivation" : [
      {
        "number" : 1,
        "part" : 67
      },
      {
        "number" : 3,
        "part" : 33
      }
    ]
  },  
  "light" : {
    "color" : "#ff0000",
    "effect" : "fixed"
  }
}
```

### JSON sent to API to request a mqtt publish

```json
{
  "cocktailId":0,
  "size":1,
  "light" : {
    "color" : "#ff0000",
    "effect" : "fixed"
  }
}
```

### config file example

config.json : 
```json
{
  "drinks" : [
    {
      "id" : "vodka",
      "name" : "Vodka",
      "type" : "alcohol"
    },
    {
      "id" : "jager",
      "name" : "Jägermeister",
      "type" : "alcohol"
    },
    {
      "id" : "gin",
      "name" : "Gin",
      "type" : "alcohol"
    },
    {
      "id" : "ricard",
      "name" : "Ricard",
      "type" : "alcohol"
    },
    {
      "id" : "agrum",
      "name" : "Schweppes Agrum",
      "type" : "soft"
    },
    {
      "id" : "tonic",
      "name" : "Schweppes Tonic",
      "type" : "soft"
    },
    {
      "id" : "coca",
      "name" : "Coca Cola",
      "type" : "soft"
    },
    {
      "id" : "redbull",
      "name" : "RedBull",
      "type" : "soft" 
    }
  ],
  "cocktails" : [
    {
      "id":0,
      "name" : "Vodka Schweppes",
      "color" : "#ffc042",
      "ingredients" : [
        {
          "id" : "vodka",
          "part" : 33
        },
        {
          "id" : "agrum",
          "part" : 67
        }
      ]
    },
    {
      "id":1,
      "name" : "Vodka RedBull",
      "color" : "#2730d9",
      "ingredients" : [
        {
          "id" : "vodka",
          "part" : 33
        },
        {
          "id" : "redbull",
          "part" : 67
        }
      ]
    },
    {
      "id":2,
      "name" : "Jaeger Bomb",
      "color" : "#2730d9",
      "ingredients" : [
        {
          "id" : "jager",
          "part" : 33
        },
        {
          "id" : "redbull",
          "part" : 67
        }
      ]
    }
  ],
  "pumps" : [
    {
      "number" : 1,
      "content" : "jager"
    },
    {
      "number" : 2,
      "content" : "vodka"
    },
    {
      "number" : 3,
      "content" : "redbull" 
    }
  ]
}
```
