import json
import datetime
from decimal import Decimal

abcd = {
    "president": {
        "name": "Zaphod Beeblebrox",
        "species": "Betelgeusian"
    }
}

json_string = json.loads(abcd)
print(json_string)
