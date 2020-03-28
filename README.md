# MurgeMachine-RestAPI
Rest API for the Murge Machine project


JSON sent via MQTT to activate the pumps format : 
```json
{
  "pumps" : [
    {
      "id" : "1",
      "part" : "67"
    },
    {
      "id" : "3",
      "part" : "33"
    }
  ]
  "led" : {
    "ledeffect" : "blink",
    "colors" : [
      {
        "color" : "#ff0000",
        "time" : "0.3"
      },
      {
        "color" : "#0000ff",
        "time" : "0.7"
      }
    ]
  }
}
```
